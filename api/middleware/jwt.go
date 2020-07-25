package middleware

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/rithikjain/quickscan-backend/api/view"
	"github.com/rithikjain/quickscan-backend/pkg"
	"log"
	"net/http"
	"os"
)

func Validate(h http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwt_secret")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return jwtMiddleware.Handler(h)
}

func ValidateAndGetClaims(ctx context.Context, role string) (map[string]interface{}, error) {
	token, ok := ctx.Value("user").(*jwt.Token)
	if !ok {
		log.Println(token)
		return nil, view.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Println(claims)
		return nil, view.ErrInvalidToken
	}

	if claims.Valid() != nil {
		return nil, view.ErrInvalidToken
	}

	if claims["role"].(string) != role {
		log.Println(claims["role"])
		return nil, pkg.ErrUnauthorized
	}
	return claims, nil
}
