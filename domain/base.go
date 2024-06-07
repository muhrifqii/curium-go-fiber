package domain

import "time"

type BaseModel struct {
	ID        int64      `json:"-" db:"id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type AuditableModel struct {
	BaseModel
	CreatedBy *int64 `json:"-" db:"created_by"`
	UpdatedBy *int64 `json:"-" db:"updated_by"`
}
