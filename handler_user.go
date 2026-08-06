package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ha165/rssagg/internal/auth"
	"github.com/ha165/rssagg/internal/database"
)

func (apiCFG *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := apiCFG.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 201, user)
}

func (apiCFG *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Error getting API key: %v", err))
		return
	}
	user, err := apiCFG.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Error getting user: %v", err))
		return
	}
	respondWithJSON(w, 200, user)
}
