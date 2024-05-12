package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// const (
// 	ERR_PARSE_TOKEN            = err.NewError("jwt: parse JWT token")
// 	ERR_INVALID_TOKEN          = err.NewError("jwt: invalid JWT token")
// 	ERR_INVALID_CLAIMS         = err.NewError("jwt: invalid JWT claims")
// 	ERR_INVALID_CLAIMS_UID     = err.NewError("jwt: invalid JWT claims uid")
// 	ERR_INVALID_CLAIMS_EXP     = err.NewError("jwt: invalid JWT claims, token expired")
// 	ERR_UNEXPECTED_SIGN_METHOD = err.NewError("jwt: unexpected jwt signing method")
// )

type Claims struct {
	UID uuid.UUID `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateToken(uuid uuid.UUID) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := &Claims{
		uuid,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (uuid.UUID, error) {
	secret := os.Getenv("JWT_SECRET")

	var uid uuid.NullUUID

	token, err := jwt.ParseWithClaims(tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected jwt signing method")
			}
			return []byte(secret), nil
		},
	)

	if err != nil {
		fmt.Println(err)
		return uid.UUID, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return uid.UUID, errors.New("invalid jwt claims")
	}

	if !token.Valid {
		return uid.UUID, errors.New("invalid jwt token")
	}

	if claims.UID == uuid.Nil {
		return uid.UUID, errors.New("invalid uid in jwt token")
	}

	uid.UUID = claims.UID

	return uid.UUID, nil

}
