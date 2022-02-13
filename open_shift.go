package fiscalizer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (self *Kkm) OpenShift() (err error) {
	var (
		res      *http.Response
		respbody Response
		body     []byte
	)
	jsonData := map[string]string{"IdKkm": self.kkm_id}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/shifts", self.host), bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.Status == http.StatusOK {
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
	err = errors.New(respbody.Msg)
	return
}
