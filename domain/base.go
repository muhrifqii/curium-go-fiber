package domain

import "time"

type BaseModel struct {
	ID        int64     `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuditableModel struct {
	BaseModel
	CreatedBy int64 `json:"-"`
	UpdatedBy int64 `json:"-"`
}
