package domain

import "time"

type (
	BaseModel struct {
		ID        int64      `json:"-" db:"id"`
		CreatedAt *time.Time `json:"created_at" db:"created_at"`
		UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	}

	BaseOrganizationModel struct {
		ID        int64      `json:"-" db:"id"`
		OrgID     string     `json:"org_id" db:"organization_id"`
		CreatedAt *time.Time `json:"created_at" db:"created_at"`
		UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	}
)

type AuditableModel struct {
	BaseModel
	CreatedBy *int64 `json:"-" db:"created_by"`
	UpdatedBy *int64 `json:"-" db:"updated_by"`
}

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}
