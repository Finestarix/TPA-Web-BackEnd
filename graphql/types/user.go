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
				"firstname": &graphql.Field{
					Type: graphql.String,
				},
				"lastname": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"isemailconfirm": &graphql.Field{
					Type: graphql.Int,
				},
				"phonecode": &graphql.Field{
					Type: graphql.String,
				},
				"phone": &graphql.Field{
					Type: graphql.String,
				},
				"isphoneconfirm": &graphql.Field{
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
				"zipcode": &graphql.Field{
					Type: graphql.Int,
				},
			},
			Description: "Type for User",
		})
	}
	return userType
}
