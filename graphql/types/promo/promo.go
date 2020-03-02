package promo

import (
	"github.com/graphql-go/graphql"
)

var promoType *graphql.Object

func GetPromoType() *graphql.Object {
	if promoType == nil {
		promoType = graphql.NewObject(graphql.ObjectConfig{
			Name: "PromoType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"detail": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Promo",
		})
	}
	return promoType
}

