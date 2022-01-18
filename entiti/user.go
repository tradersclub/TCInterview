package entiti

import "time"

type User struct {
	id        int
	user      string
	Password  string
	createdAt *time.Time `db:"created_at" json:"created_at"`
}
