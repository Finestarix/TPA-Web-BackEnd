package resolvers

import (
	"github.com/graphql-go/graphql"
	"../models"
)

func AllUsers(p graphql.ResolveParams) (i interface{}, err error) {
	users := models.GetAllUser()
	return users, nil
}

func GetUserByID(p graphql.ResolveParams) (i interface{}, err error) {
	searchID := p.Args["id"].(int)
	user := models.GetUserByID(searchID)
	return user, nil
}

func GetUserByPhoneEmail(p graphql.ResolveParams) (i interface{}, err error) {
	searchEmailPhone := p.Args["emailphone"].(string)
	user := models.GetUserByEmail(searchEmailPhone)

	if len(user) == 0 {
		user = models.GetUserByPhone(searchEmailPhone)
	}

	return user, nil
}

func InsertUser(p graphql.ResolveParams) (i interface{}, err error) {
	firstname := p.Args["firstname"].(string)
	lastname := p.Args["lastname"].(string)
	email := p.Args["email"].(string)
	phonecode := p.Args["phonecode"].(string)
	phone := p.Args["phone"].(string)
	password := p.Args["password"].(string)

	newUser := models.InsertUser(firstname, lastname, email, phonecode, phone, password)

	return newUser, nil
}

func UpdateUserProfile(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	title := p.Args["title"].(string)
	firstname := p.Args["firstname"].(string)
	lastname := p.Args["lastname"].(string)
	city := p.Args["city"].(string)
	address := p.Args["address"].(string)
	zipcode := p.Args["zipcode"].(int)

	updatedUser := models.UpdateUserProfile(id, title, firstname, lastname, city, address, zipcode)

	return updatedUser, nil
}

func DeleteUser(p graphql.ResolveParams) (i interface{}, err error) {
	searchID := p.Args["id"].(int)

	deletedUser := models.DeleteUser(searchID)

	return deletedUser, nil
}