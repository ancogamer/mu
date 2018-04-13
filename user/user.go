package user

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CommonModelFields ia a base for gorm.Model with json type
type CommonModelFields struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// User is a Human
type User struct {
	CommonModelFields
	FacebookID string `json:"facebookID,omitempty"`
	Token      string `json:"token,omitempty"`
}

// SecretJWT is var on OS ENV
var SecretJWT string

// MyCustomClaims is used for build JWT
type MyCustomClaims struct {
	User
	jwt.StandardClaims
}

// GenerateToken generate JWT for auth
func (u User) GenerateToken(secret string) (string, error) {
	mySigningKey := []byte(secret)

	u.Token = ""

	claims := MyCustomClaims{
		u,
		jwt.StandardClaims{
			ExpiresAt: 2500000000,
			Issuer:    "mu",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

// ValidateToken validate a JWT
func (u User) ValidateToken(secret string) (bool, error) {
	// mySigningKey := []byte(secret)
	// Token from another example.  This token is expired
	var tokenString = u.Token

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("fiscaluno"), nil
	// })

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("fiscaluno"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	return false, nil
}
