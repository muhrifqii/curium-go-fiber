package domain

import "time"

type BaseModel struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuditableModel struct {
	BaseModel
	CreatedBy int64
	UpdatedBy int64
}
