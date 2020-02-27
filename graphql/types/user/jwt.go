package user

import "github.com/graphql-go/graphql"

var jwtType *graphql.Object

func GetJWTType() *graphql.Object {
	if jwtType == nil {
		jwtType = graphql.NewObject(graphql.ObjectConfig{
			Name: "JWTType",
			Fields: graphql.Fields{
				"jwtToken": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for JWT",
		})
	}
	return jwtType
}
