package resolvers

import (
	"../models"
	"github.com/graphql-go/graphql"
)

func AllPhoneCode(p graphql.ResolveParams) (i interface{}, err error) {
	phoneCode := models.GetAllPhoneCode()
	return phoneCode, nil
}

func GetPhoneCode(p graphql.ResolveParams) (i interface{}, err error) {
	searchPhoneCode := p.Args["phonecode"].(string)
	phoneCode := models.GetPhoneCode(searchPhoneCode)
	return phoneCode, nil
}