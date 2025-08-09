package database

import (
	"context"
	"main/internal/models"
	"time"

	"github.com/google/uuid"
)

func (db *Database) CreateEvent(event models.Event) (uuid.UUID, error) {
	var u uuid.UUID
	err := db.Pg.QueryRow(context.Background(), `INSERT INTO events (
	Title,
	Description,
	MaxMembers,
	CountMembers,
	Expaired,
	Creator) VALUES ($1,$2,$3,$4,$5,$6) RETURNING uuid;`, event.Title, event.Description, event.MaxMembers, 0, time.Now(), event.Creator).Scan(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (db *Database) GetEvent(u uuid.UUID) (models.Event, error) {
	var event models.Event
	tx, err := db.Pg.Begin(context.Background())
	if err != nil {
		return event, err
	}
	defer tx.Rollback(context.Background())
	err = db.Pg.QueryRow(context.Background(), `SELECT * FROM events WHERE uuid = $1;`, u).Scan(
		&event.Uuid,
		&event.Title,
		&event.Description,
		&event.MaxMembers,
		&event.CountMembers,
		&event.Creator,
		&event.Created,
		&event.Expaired)
	if err != nil {
		return event, err
	}

	rows, err := db.Pg.Query(context.Background(), `SELECT uuid_user FROM events_members WHERE uuid_event = $1;`, u)
	if err != nil {
		return event, err
	}
	defer rows.Close()
	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return event, err
		}
		event.Members = append(event.Members, userID)
	}
	tx.Commit(context.Background())
	return event, nil
}

func (db *Database) SubscribeEvent(user, event uuid.UUID) error {
	_, err := db.Pg.Exec(context.Background(), `INSERT INTO events_members (uuid_event, uuid_user) VALUES ($1,$2);`, event, user)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UnSubscribeEvent(user, event uuid.UUID) error {
	_, err := db.Pg.Exec(context.Background(), `DELETE FROM events_members WHERE uuid_event = $1 AND uuid_user = $2;`, event, user)
	if err != nil {
		return err
	}
	return nil
}
