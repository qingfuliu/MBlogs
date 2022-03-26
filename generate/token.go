package generate

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"goProject/models"
	"time"
)

var Msecret = []byte("blogs")

const (
	refreshTokenLife = time.Hour * 2
	assessTokenLife  = time.Minute * 10
)

func getSecret(*jwt.Token) (interface{}, error) {
	return Msecret, nil
}

func getToken(username string) (string, error) {
	mClaims := &models.MClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Subject:   "ifisLogins",
			ExpiresAt: time.Now().Add(assessTokenLife).Unix(),
		},
	}
	//创建一个token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mClaims)
	return token.SignedString(Msecret)
}

func GetAssAndRefToken(username string) (aToken string, rToken string, err error) {
	mClaims := &models.MClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(assessTokenLife).Unix(),
			Issuer:    "lqf",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mClaims)
	aToken, err = token.SignedString(Msecret)
	if err != nil {
		return
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(refreshTokenLife).Unix(),
		Issuer:    "lqf",
	})
	rToken, err = token.SignedString(Msecret)
	return
}

func ParseToken(jwtToken string) (*models.MClaims, error) {
	token, err := jwt.Parse(jwtToken, getSecret)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.MClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func RefreshToken(aToken, rToken string) (newToken string, err error) {
	if _, err = jwt.Parse(rToken, getSecret); err != nil {
		return
	}
	var mClaims models.MClaims
	_, err = jwt.ParseWithClaims(aToken, &mClaims, getSecret)
	v, ok := err.(*jwt.ValidationError)
	if ok && v.Errors == jwt.ValidationErrorExpired {
		return getToken(mClaims.Username)
	}
	return
}
