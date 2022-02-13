package fiscalizer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Kkm struct {
	kkm_id    int
	kkm_login string
	kkm_pass  string
	uid       string
	token     string
	host      string
	queue     chan document
	result    chan FiscalizationResult
}

func NewKkm(kkm_id int, kkm_login, kkm_pass string) *Kkm {
	kkm := Kkm{kkm_id: kkm_id, kkm_login: kkm_login, kkm_pass: kkm_pass}
	kkm.queue = make(chan document)
	kkm.result = make(chan FiscalizationResult)
	kkm.run()
	return &kkm
}

func (self *Kkm) run() {

	var (
		doc document
		err error
		res FiscalizationResult
	)

	for {
		doc = <-self.queue
		res, err = send_document(doc)
		if err != nil {
			res.Err = err
		}
		self.result <- res
	}
}

func (self *Kkm) auth() (err error) {
	var (
		res  *http.Response
		body []byte
	)
	var respbody struct {
		response
		Data auth_response `json:"Data,omitempty"`
	}
	jsonData := map[string]string{"Login": self.kkm_login, "Password": self.kkm_pass}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth", self.host), bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = json.Unmarshal(body, &respbody)
	if err != nil {
		return
	}
	if respbody.Status != http.StatusOK {
		err = errors.New(respbody.Msg)
		return
	}
	self.token = respbody.Data.Token
	return
}

func (self *Kkm) recover_uid() (err error) {
	var (
		body   []byte
		status int
	)
	var uid_resp struct {
		response
		Data struct {
			Uid string `json:"Uid" `
		} `json:"Data" `
	}

	status, body, err = self.make_fiscalrequest(nil, fiscal_operation{endpoint: "uid/%d", method: "GET"})
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Error geting UID %d", res.StatusCode))
		return
	}

	err = json.Unmarshal(body, &uid_resp)
	if err != nil {
		return
	}
	self.uid = uid_resp.Data.Uid
	return
}

func (self *Kkm) make_fiscal_request(payload []byte, operation fiscal_operation) (status int, body []byte, err error) {
	var (
		req     *http.Request
		res     *http.Response
		new_uid string
	)

	url := fmt.Sprintf("%s/%s", self.host, fmt.Sprintf(operation.endpoint, self.kkm_id))
	req, err = http.NewRequest(operation.method, url, bytes.NewBuffer(payload))
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Uid", uid)
	req.Header.Add("authorization", "Bearer "+self.token)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	status = res.StatusCode
	new_uid = res.Header.Get("Uid")
	if new_uid != "" {
		self.uid = new_uid
	}

	body, err = ioutil.ReadAll(res.Body)
	return
}

func (self *Kkm) send_document(doc document) (result FiscalizationResult, err error) {

	var (
		payload []byte
	)

	var resp struct {
		response
		Data FiscalizationResult `json:"Data"`
	}

	payload, err = json.Marshal(document)
	if err != nil {
		return
	}

reqloop:
	for {
		status, body, err = self.make_fiscalrequest(payload, doc.doc_type)
		if err != nil {
			return
		}
		switch status {
		case 200:
			break reqloop
		case 208:
			break reqloop
		case 449:
			continue reqloop
		case 452:
			err = self.recover_uid()
			if err != nil {
				return
			}
			continue reqloop
		case 401:
			err = self.auth()
			if err != nil {
				return
			}
			continue reqloop
		default:
			data, _ := ioutil.ReadAll(res.Body)
			err = errors.New("Fiscalization error: " + strconv.Itoa(res.StatusCode) + string(data))
			return

		}
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	result = resp.Data
	return
}
