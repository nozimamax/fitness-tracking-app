package token

import (
	"errors"
	"log"
	"time"

	"github.com/nozimamax/fitness-tracking-app/internal/pkg/load"
	"github.com/form3tech-oss/jwt-go"
)

type Tokens struct {
	RefreshToken string
}

var cfg load.Config
var tokenKey = cfg.TokenKey

func GenereteJWTToken(email string) *Tokens {
	
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["user_email"] = email
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting refresh token : ", err)
	}

	return &Tokens{
		RefreshToken: refresh,
	}
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
} 
