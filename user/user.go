package user

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fiscaluno/mu/db"
	"github.com/fiscaluno/pandorabox"
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
	User User `json:"user,omitempty"`
	jwt.StandardClaims
}

// TokenDetail response for token details
type TokenDetail struct {
	Valid  bool   `json:"valid"`
	UserID uint   `json:"userID,omitempty"`
	Active bool   `json:"active"`
	Note   string `json:"note,omitempty"`
}

// GetAll Users
func GetAll() []User {
	db := db.Conn()
	defer db.Close()
	var users []User
	db.Find(&users)
	return users
}

func (user User) AddWithVerification() (User, error) {

	// validation Facebook ID don't repeat
	users := GetByQuery("facebook_id = ?", user.FacebookID)
	if len(users) > 0 {
		user, err := users[0].NewToken()
		if err != nil {
			return user, err
		}
		return user, nil
	}

	db := db.Conn()
	defer db.Close()

	user, err := user.NewToken()
	if err != nil {
		return user, err
	}

	db.Create(&user)

	return user, nil
}

func GetByID(id int) User {
	db := db.Conn()
	defer db.Close()

	var user User

	db.Find(&user, id)

	return user
}

func GetByQuery(query string, value interface{}) []User {
	db := db.Conn()
	defer db.Close()

	var users []User

	db.Find(&users, query, value)
	return users
}

func (user User) NewToken() (User, error) {
	secret := pandorabox.GetOSEnvironment("SECRET_JWT", "fiscaluno")

	timeExp := pandorabox.DateAddDays(1)

	token, err := user.GenerateToken(secret, timeExp)
	if err != nil {
		return user, err
	}

	user.Token = token

	return user, nil
}

// GenerateToken generate JWT for auth
func (u User) GenerateToken(secret string, expDate int64) (string, error) {
	mySigningKey := []byte(secret)

	u.Token = ""

	claims := MyCustomClaims{
		u,
		jwt.StandardClaims{
			ExpiresAt: expDate,
			Issuer:    "mu",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

// ValidateToken validate a JWT
func (u User) ValidateToken(secret string) (bool, error) {

	valid, err := TokenInfos(u.Token)
	if err != nil {
		return false, err
	}

	ret := valid.Active

	return ret, nil
}

// TokenInfos return ifos of JWT
func TokenInfos(tokenString string) (TokenDetail, error) {
	// mySigningKey := []byte(secret)
	// Token from another example.  This token is expired

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("fiscaluno"), nil
	// })

	detail := TokenDetail{
		Valid:  true,
		Active: true,
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("fiscaluno"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		detail.Note = "You look nice today"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		detail.Valid = false
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			detail.Valid = false
			detail.Active = false
			detail.Note = "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			detail.Active = false
			detail.Note = "Timing is everything"
			return detail, nil

		} else {
			fmt.Println("Couldn't handle this token:", err)
			detail.Active = false
			detail.Note = err.Error()
			return detail, err

		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		detail.Valid = false
		detail.Active = false
		detail.Note = err.Error()
		return detail, err
	}

	return detail, nil

}
