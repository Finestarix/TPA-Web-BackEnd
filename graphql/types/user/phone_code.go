package user

import "github.com/graphql-go/graphql"

var phoneCodeType *graphql.Object

func GetPhoneCodeType() *graphql.Object {
	if phoneCodeType == nil {
		phoneCodeType = graphql.NewObject(graphql.ObjectConfig{
			Name: "UserPhoneCode",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Phone Code",
		})
	}
	return phoneCodeType
}
