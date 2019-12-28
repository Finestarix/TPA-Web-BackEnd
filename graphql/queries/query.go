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
			"User": {
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:     resolvers.GetUser,
				Description: "Get User",
			},
		},
	})

}
