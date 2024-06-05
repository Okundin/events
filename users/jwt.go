package users

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "6bee3ff8034bd862eee945e8fee35a0dda51795a23174e7f98b03ff71d37dadf" // Load key from somewhere, for example a file

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // token is valid for 2 hours: what then?
	})

	// converting token to a string
	signedJWT, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("error signing the token")
	}

	return signedJWT, nil
}

func VerifyToken(token string) (userID int64, err error) {
	// parse the token provided by the user
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// verifying the signing method type
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// comparing the user token with the secret key
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, jwt.ErrTokenNotValidYet
	}

	// checking if Claims is of the type MapClaims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, jwt.ErrInvalidType
	}
	// accessing data in claims variable to get userID
	userID = int64(claims["userID"].(float64))

	return userID, nil
}
