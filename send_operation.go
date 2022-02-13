package fiscalizer

type FiscalizationResult struct {
	FiscalNumber     uint64 `json:"FiscalNumber"`
	AutonomousNumber uint32 `json:"AutonomousNumber"`
	DocumentId       uint64 `json:"IdDocument"`
	Location         string `json:"Location"`
	Receipt          string `json:"Reciept"`
	err              error  `json:"-"`
}

func (self *FiscalizationResult) Error() (*FiscalizationResult, error) {
	return self, self.Err
}

func (self *Kkm) SendDocument(doc document) (result FiscalizationResult, err error) {
	self.queue <- doc
	result = <-self.result
	err = result.Error()
	return
}
