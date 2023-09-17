package main

import (	
	"time"

	"github.com/MCanhisares/chessbeli/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email			string `json:"email"`
	Password	string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username      string `json:"name"`
}

func dbUserToUser(dbUser database.User) User {	
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username: dbUser.Username,
		Email: dbUser.Email,
		Password: dbUser.Password,
	}
}

type ClientInfo struct {
	Id       string
	Secret   string
	Domain   string
	Public bool
	UserId   string
}

func (c ClientInfo) GetDomain() string {
	return c.Domain
}
func (c ClientInfo) GetID() string {
	return c.Id
}
func (c ClientInfo) GetSecret() string {
	return c.Secret
}
func (c ClientInfo) IsPublic() bool {
	return c.Public
}
func (c ClientInfo) GetUserID() string {
	return c.UserId
}
