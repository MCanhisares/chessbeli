package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MCanhisares/chessbic/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json: "name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	dbUser, err := apiConfig.DB.Createuser(r.Context(), database.CreateuserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user %v", err))
		return
	}
	user := dbUserToUser(dbUser)
	respondWithJson(w, 200, user)
}