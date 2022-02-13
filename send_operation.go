package fiscalizer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

type FiscalizationResult struct {
	FiscalNumber     uint64 `json:"FiscalNumber"`
	AutonomousNumber uint32 `json:"AutonomousNumber"`
	DocumentId       uint64 `json:"IdDocument"`
	Location         string `json:"Location"`
	Receipt          string `json:"Reciept"`
}

func (self *Kkm) send_operation(operation fiscal_operation, document document) (result FiscalizationResult, err error) {

	var (
		payload []byte
	)

	var resp struct {
		Response
		Data FiscalizationResult `json:"Data"`
	}

	payload, err = json.Marshal(document)
	if err != nil {
		return
	}

reqloop:
	for {
		status, body, err = self.make_fiscalrequest(payload, fiscal_operation)
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
