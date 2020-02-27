package core

import (
	models "../../models/core"
	"github.com/graphql-go/graphql"
)

func InsertSubscription(p graphql.ResolveParams) (i interface{}, err error) {
	email := p.Args["email"].(string)

	newSubscription := models.InsertNewSubscription(email)

	return newSubscription, nil
}

func SendSubscriptionToAll(p graphql.ResolveParams) (i interface{}, err error) {
	subscription := models.SendSubscriptionToAll()

	return subscription, nil
}