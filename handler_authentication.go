package main

import (
	"context"
	"github.com/MCanhisares/chessbeli/internal/database"
)

func (apiConfig *apiConfig)handlerAuthentication(ctx context.Context, user string, password string) (userID string, err error){
	dbUser, err := apiConfig.DB.Authenticate(ctx, database.AuthenticateParams{
		Email: user,
		Password: password,
	})
	return dbUser.ID.String(), err	
}