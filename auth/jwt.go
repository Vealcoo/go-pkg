package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Meta interface{}
	jwt.StandardClaims
}

func SetJWTToken(secrct string, expireAt int64, meta interface{}) (string, error) {
	claims := Claims{
		Meta: meta,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secrct))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ReadJWT(secrct, token string) (interface{}, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secrct), nil
	})
	if err != nil || !tokenClaims.Valid {
		return nil, ErrInvaildToken
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok {
		return nil, ErrInvaildToken
	}

	return claims.Meta, nil
}
