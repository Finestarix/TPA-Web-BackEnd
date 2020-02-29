package blog

import "github.com/graphql-go/graphql"

var blogType *graphql.Object

func GetBlogType() *graphql.Object {
	if blogType == nil {
		blogType = graphql.NewObject(graphql.ObjectConfig{
			Name: "BlogType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userID": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"viewCount": &graphql.Field{
					Type: graphql.Int,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
				"category": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Blog Model",
		})
	}
	return blogType
}

