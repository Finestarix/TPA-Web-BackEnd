package user

import (
	models "../../models/user"
	"github.com/graphql-go/graphql"
)

func AllPhoneCode(p graphql.ResolveParams) (i interface{}, err error) {
	phoneCode := models.GetAllPhoneCode()
	return phoneCode, nil
}

func GetPhoneCode(p graphql.ResolveParams) (i interface{}, err error) {
	searchPhoneCode := p.Args["code"].(string)
	phoneCode := models.GetPhoneCode(searchPhoneCode)
	return phoneCode, nil
}