package types

import "github.com/graphql-go/graphql"

var phoneCodeType *graphql.Object

func GetPhoneCodeType() *graphql.Object {
	if phoneCodeType == nil {
		phoneCodeType = graphql.NewObject(graphql.ObjectConfig{
			Name: "UserPhoneCode",
			Fields: graphql.Fields{
				"phonecodeid": &graphql.Field{
					Type: graphql.Int,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
				"phonecode": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Phone Code",
		})
	}
	return phoneCodeType
}
