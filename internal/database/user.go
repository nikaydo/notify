package database

import (
	"context"
	"main/internal/models"

	"github.com/google/uuid"
)

func (db *Database) CreateUser(user models.User) (uuid.UUID, error) {
	var uuid uuid.UUID
	err := db.Pg.QueryRow(context.Background(), "INSERT INTO users (login,password) VALUES ($1,$2) RETURNING uuid;", user.Login, user.Password).Scan(&uuid)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}

func (db *Database) UpdateUser() {

}

func (db *Database) DeleteUser() {

}
