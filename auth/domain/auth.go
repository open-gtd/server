package domain

type (
	Name string
	SecurityCode string
	Token string
)

type Auth struct {
	UserName     Name
	SecurityCode SecurityCode
}

type Cert struct {
	Token Token
}
