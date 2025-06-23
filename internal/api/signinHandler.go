package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
	"github.com/golang-jwt/jwt/v5"
)

var pass struct {
	Password string `json:"password"`
}

func signinHandler(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&pass)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": "error to decode json"})
		return
	}

	envPassword := os.Getenv("TODO_PASSWORD")
	if envPassword == "" {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": "error authentication is not implemented"})
		return
	}

	if pass.Password != envPassword {
		service.SendResponse(w, http.StatusUnauthorized, map[string]string{"error": "error invalid password"})
		return
	}

	hash := sha256.Sum256([]byte(envPassword))
	hashStr := hex.EncodeToString(hash[:])

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"hash": hashStr,
		"exp":  time.Now().Add(8 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(configs.JwtKey)
	if err != nil {
		service.SendResponse(w, http.StatusUnauthorized, map[string]string{"error": "error create token"})
		return
	}

	service.SendResponse(w, http.StatusOK, map[string]string{"token": tokenString})
}
