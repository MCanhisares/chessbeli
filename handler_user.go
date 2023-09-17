package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MCanhisares/chessbeli/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	dbUser, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username: params.Username,
		Email: params.Email,
		Crypt: params.Password,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user %v", err))
		return
	}
	user := dbUserToUser(dbUser)
	respondWithJson(w, 200, user)
}