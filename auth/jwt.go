package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "ji3g45/4u;6"
)

type Claims struct {
	ObjectId string
	jwt.StandardClaims
}

func SetToken(objectId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	claims := Claims{
		ObjectId: objectId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func AuthJWT(token string) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims.ObjectId, nil
	}

	return "", ErrVaild
}
