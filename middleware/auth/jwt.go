package auth

import (
	"SensorProject/models"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

// Exception struct declaration
type Exception struct {
	Message string `json:"message"`
}

// Checks if token is passed in the request headers
// Validates and decodes the token

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") // Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			// Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
