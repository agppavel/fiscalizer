package fiscalizer

type auth_response struct {
	Token string `json:"Token"`
}

type response struct {
	Status int    `json:"Status"`
	Msg    string `json:"Message,omitempty"`
}

