package service

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if pass == "" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			SendResponse(w, http.StatusUnauthorized, "Authentication required")
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return configs.JwtKey, nil
		})

		if err != nil || !token.Valid {
			SendResponse(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			SendResponse(w, http.StatusUnauthorized, map[string]string{"error": "invalid token claims"})
			return
		}

		tokenHash, ok := claims["hash"].(string)
		if !ok {
			log.Println("1234")
			SendResponse(w, http.StatusUnauthorized, map[string]string{"error": "invalid token claims"})
			return
		}

		hash := sha256.Sum256([]byte(pass))
		currentHash := hex.EncodeToString(hash[:])

		if tokenHash != currentHash {
			SendResponse(w, http.StatusUnauthorized, map[string]string{"error": "token is no longer valid"})
			return
		}

		next(w, r)
	})
}
