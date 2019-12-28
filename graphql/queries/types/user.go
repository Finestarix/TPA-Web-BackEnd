package types

import "github.com/graphql-go/graphql"

var userType *graphql.Object

func GetUserType() *graphql.Object {
	if userType == nil {
		userType = graphql.NewObject(graphql.ObjectConfig{
			Name: "UserType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"is_email_confirm": &graphql.Field{
					Type: graphql.Int,
				},
				"phone_code": &graphql.Field{
					Type: graphql.String,
				},
				"phone": &graphql.Field{
					Type: graphql.String,
				},
				"is_phone_confirm": &graphql.Field{
					Type: graphql.Int,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"zip_code": &graphql.Field{
					Type: graphql.Int,
				},
			},
			Description: "Type for User",
		})
	}
	return userType
}
