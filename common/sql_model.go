package common

import "time"

type SQLModel struct {
	ID        int64     `json:"id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
