package converters

import (
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/auth/presentation"
)

func ConvertLoginDataToBusiness(loginData presentation.LoginData) (business.LoginData, error) {

	return business.LoginData{
		Name:         domain.Name(loginData.UserName),
		SecurityCode: domain.SecurityCode(loginData.SecurityCode),
	}, nil
}
