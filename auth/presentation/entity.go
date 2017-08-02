package presentation

type Cert struct {
	Token string `json:"token"`
}

type LoginData struct {
	UserName     string `json:"user_name"`
	SecurityCode string `json:"security_code"`
}
