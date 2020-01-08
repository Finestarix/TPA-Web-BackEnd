package queries

import (
	"../../resolvers"
	"../types"
	"github.com/graphql-go/graphql"
)

func Root() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"AllUser": {
				Type:        graphql.NewList(types.GetUserType()),
				Resolve:     resolvers.AllUsers,
				Description: "Get All Users",
			},
			"UserByID": {
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:     resolvers.GetUserByID,
				Description: "Get User By ID",
			},
			"UserByEmailAndPhone": {
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"emailphone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetUserByPhoneEmail,
				Description: "Get User By Email or Phone",
			},
			"UserLogin": {
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"emailphone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.UserLogin,
				Description: "User Login (Email, Phone, Password)",
			},

			"AllPhoneCode": {
				Type:        graphql.NewList(types.GetPhoneCodeType()),
				Resolve:     resolvers.AllPhoneCode,
				Description: "Get All Phone Code",
			},
			"GetPhoneCode": {
				Type: types.GetPhoneCodeType(),
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetPhoneCode,
				Description: "Get Phone Code",
			},
		},
	})

}
