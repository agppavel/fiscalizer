package fiscalizer

type AuthResponse struct {
	Token string `json:"Token"`
}

type Response struct {
	Status int    `json:"Status"`
	Msg    string `json:"Message,omitempty"`
}
