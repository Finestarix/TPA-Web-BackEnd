package resolvers

import (
	"github.com/graphql-go/graphql"
	"../models"
)

func AllUsers(p graphql.ResolveParams) (i interface{}, err error) {
	users := models.GetAllUser()
	return users, nil
}

func GetUser(p graphql.ResolveParams) (i interface{}, err error) {
	searchID := p.Args["id"].(int)
	user := models.GetUser(searchID)
	return user, nil
}