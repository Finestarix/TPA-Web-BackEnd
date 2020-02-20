package resolvers

import (
	"../models"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllUsers(p graphql.ResolveParams) (i interface{}, err error) {
	users := models.GetAllUser()

	if len(users) == 0 {
		return nil, errors.New("error: no data found")
	}

	return users, nil
}

func GetUserByID(p graphql.ResolveParams) (i interface{}, err error) {
	searchID := p.Args["id"].(string)
	user := models.GetUserByID(searchID)

	return user, nil
}

func GetUserByPhoneEmail(p graphql.ResolveParams) (i interface{}, err error) {
	searchEmailPhone := p.Args["emailphone"].(string)
	user := models.GetUserByEmail(searchEmailPhone)

	if user.ID == 0 {
		user = models.GetUserByPhone(searchEmailPhone)
	}

	return user, nil
}

func UserLogin(p graphql.ResolveParams) (i interface{}, err error) {
	searchEmailPhone := p.Args["emailphone"].(string)
	searchPassword := p.Args["password"].(string)

	user := models.GetUserByEmailAndPassword(searchEmailPhone, searchPassword)

	if user.ID == 0 {
		user = models.GetUserByPhoneAndPassword(searchEmailPhone, searchPassword)
	}
	jwtToken := models.CreateNewJWTTokenUser(user)

	if user.ID == 0 {
		jwtToken.JwtToken = "";
	}

	return jwtToken, nil
}

func InsertUser(p graphql.ResolveParams) (i interface{}, err error) {
	firstname := p.Args["firstname"].(string)
	lastname := p.Args["lastname"].(string)
	email := p.Args["email"].(string)
	phonecode := p.Args["phonecode"].(string)
	phone := p.Args["phone"].(string)
	password := p.Args["password"].(string)
	image := p.Args["image"].(string)

	newUser := models.InsertUser(firstname, lastname, email, phonecode, phone, password, image)

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