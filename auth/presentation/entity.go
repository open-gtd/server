package presentation

type (
	Token        string
	UserName     string
	SecurityCode string
)

type Cert struct {
	Token Token `json:"token"`
}

type LoginData struct {
	UserName     UserName     `json:"user_name"`
	SecurityCode SecurityCode `json:"security_code"`
}
