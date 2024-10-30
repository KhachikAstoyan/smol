package models

import "time"

type URLModel struct {
	Id        int       `db:"id"`
	Url       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
}
