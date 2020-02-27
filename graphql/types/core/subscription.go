package core

import "github.com/graphql-go/graphql"

var subscriptionType *graphql.Object

func GetSubscriptionType() *graphql.Object {
	if subscriptionType == nil {
		subscriptionType = graphql.NewObject(graphql.ObjectConfig{
			Name: "SubscriptionType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Subscription",
		})
	}
	return subscriptionType
}
