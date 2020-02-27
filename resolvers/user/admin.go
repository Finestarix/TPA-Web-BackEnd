package user

import (
	models "../../models/user"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllAdmin(p graphql.ResolveParams) (i interface{}, err error) {
	admins := models.GetAllAdmin()

	if len(admins) == 0 {
		return nil, errors.New("error: no data found")
	}

	return admins, nil
}

func InsertAdmin(p graphql.ResolveParams) (i interface{}, err error) {
	username := p.Args["username"].(string)
	password := p.Args["password"].(string)

	newAdmin := models.InsertAdmin(username, password)

	return newAdmin, nil
}

func LoginAdmin(p graphql.ResolveParams) (i interface{}, err error) {
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	newAdmin := models.ValidateAdminLogin(email, password)

	jwtToken := models.CreateNewJWTTokenAdmin(newAdmin)

	if newAdmin.ID == 0 {
		jwtToken.JwtToken = ""
	}

	return jwtToken, nil
}
