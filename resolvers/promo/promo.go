package promo

import (
	models "../../models/promo"
	"github.com/graphql-go/graphql"
)

func GetPromo(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	promo := models.GetPromoBYID(id);
	return promo, nil
}

func OtherPromo(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	otherPromo := models.GetAnotherPromo(id)

	return otherPromo, nil
}