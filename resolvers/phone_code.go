package resolvers

import (
	"../models"
	"github.com/graphql-go/graphql"
)

func AllPhoneCode(p graphql.ResolveParams) (i interface{}, err error) {
	phoneCode := models.GetAllPhoneCode()
	return phoneCode, nil
}