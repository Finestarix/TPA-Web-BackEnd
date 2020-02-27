package train

import (
	"../core"
	"github.com/graphql-go/graphql"
)

var trainStationType *graphql.Object

func GetTrainStationType() *graphql.Object {
	if trainStationType == nil {
		trainStationType = graphql.NewObject(graphql.ObjectConfig{
			Name: "TrainStationType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"location": &graphql.Field{
					Type: core.GetLocationType(),
				},
				"locationId": &graphql.Field{
					Type: graphql.Int,
				},
			},
			Description: "Type for Train Station",
		})
	}
	return trainStationType
}
