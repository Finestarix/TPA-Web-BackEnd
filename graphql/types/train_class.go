package types

import "github.com/graphql-go/graphql"

var trainClassType *graphql.Object

func GetTrainClassType() *graphql.Object {
	if trainClassType == nil {
		trainClassType = graphql.NewObject(graphql.ObjectConfig{
			Name: "TrainClassType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"trainId": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Train Class",
		})
	}
	return trainClassType
}
