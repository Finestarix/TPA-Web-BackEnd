package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	ID        int    `json:"id"`
	Email     string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Image     string `gorm:"Image"`
}

type JWTClaimsAdmin struct {
	jwt.StandardClaims
	ID    int    `json:"id"`
	Email string `json:"Email"`
}

type JWTClass struct {
	JwtToken string
}

var ApplicationName = "Karcis.co"
var LoginExpirationDuration = time.Duration(1) * time.Hour
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("Authentication")

func CreateNewJWTTokenUser(user User) JWTClass {

	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: time.Now().Add(LoginExpirationDuration).Unix(),
		},
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Image:     user.Image,
	}

	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		fmt.Println(err)
	}

	returnJWT := JWTClass{
		JwtToken: signedToken,
	}

	return returnJWT
}

func CreateNewJWTTokenAdmin(admin Admin) JWTClass {

	claims := JWTClaimsAdmin{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: time.Now().Add(LoginExpirationDuration).Unix(),
		},
		ID:        admin.ID,
		Email:     admin.Email,
	}

	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		fmt.Println(err)
	}

	returnJWT := JWTClass{
		JwtToken: signedToken,
	}

	return returnJWT
}