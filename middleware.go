package authPract

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
	"strings"
	"time"
)

const (
	salt       = "asdjkalhsd123123laksj"
	signingKey = "kaijdhOAS;KD'JJAKsjd"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func CheckJWT(h http.Handler, ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")

		userId, err := ParseToken(bearer)

		if err != nil {
			logrus.Fatalf("Error parsing token: %v", err)
			return
		}
		
		ctx = context.WithValue(ctx, "UserID", userId)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ParseToken(accessToken string) (int, error) {
	headerParts := strings.Split(accessToken, " ")
	token, err := jwt.ParseWithClaims(headerParts[1], &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, err
	}
	return claims.UserID, nil
}
