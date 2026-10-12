package models

import "time"

// User represents the database entity for a user.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
