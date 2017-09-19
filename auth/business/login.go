package business

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/open-gtd/server/auth/domain"
)

type LoginPresenter interface {
	ShowAuthFailed() error
	Show(cert domain.Cert) error
}

type LoginDao interface {
	Authorize(auth domain.Auth) error
}

type LoginController Controller

type LoginLogger interface {
	Logged(auth domain.Auth)
}

type Login interface {
	Run(cd LoginData) error
}

type LoginData struct {
	Name         domain.Name
	SecurityCode domain.SecurityCode
}

type login struct {
	presenter LoginPresenter
	dao       LoginDao
	logger    LoginLogger
}

func NewLogin(p LoginPresenter, d LoginDao, l LoginLogger) Login {
	return login{presenter: p, dao: d, logger: l}
}

func (c login) Run(ld LoginData) error {

	auth := domain.Auth{
		UserName:     ld.Name,
		SecurityCode: ld.SecurityCode,
	}
	err := c.dao.Authorize(auth)

	if err != nil {
		err = c.presenter.ShowAuthFailed()
		if err != nil {
			return err
		}
		return nil
	}

	t, err := generateToken(string(auth.UserName))
	if err != nil {
		return err
	}

	cert := domain.Cert{
		Token: domain.Token(t),
	}
	c.logger.Logged(auth)

	return c.presenter.Show(cert)
}

func generateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
