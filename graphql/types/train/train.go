package train

import "github.com/graphql-go/graphql"

var trainType *graphql.Object

func GetTrainType() *graphql.Object {
	if trainType == nil {
		trainType = graphql.NewObject(graphql.ObjectConfig{
			Name: "TrainType",
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
				"class": &graphql.Field{
					Type: graphql.String,
				},
				"arrival": &graphql.Field{
					Type: GetTrainStationType(),
				},
				"arrivalId": &graphql.Field{
					Type: graphql.Int,
				},
				"arrivalTime": &graphql.Field{
					Type: graphql.DateTime,
				},
				"departure": &graphql.Field{
					Type: GetTrainStationType(),
				},
				"departureId": &graphql.Field{
					Type: graphql.Int,
				},
				"departureTime": &graphql.Field{
					Type: graphql.DateTime,
				},
				"transit": &graphql.Field{
					Type: GetTrainStationType(),
				},
				"transitId": &graphql.Field{
					Type: graphql.Int,
				},
				"seat": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
			},
			Description: "Type for Train",
		})
	}
	return trainType
}
