package domain

import "time"

type Organization struct {
	Identifier string     `json:"identifier" db:"identifier"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
	Name       string     `json:"name" db:"a_name"`
	Status     string     `json:"status" db:"a_status"`
	Address    string     `json:"address" db:"an_address"`
	Email      string     `json:"email" db:"contact_email"`
	Phone      string     `json:"phone" db:"contact_phone"`
}

type (
	CreateOrganizationRequest struct {
		Identifier string            `json:"identifier" validate:"required,ne_ignore_case=public,min=3,max=18"`
		Name       string            `json:"name" validate:"required"`
		Address    string            `json:"address"`
		Email      string            `json:"email" validate:"email"`
		Phone      string            `json:"phone" validate:"e164"`
		Owner      CreateUserRequest `json:"owner"`
	}

	UpdateOrganizationRequest struct {
		Name    string `json:"name" validate:"required"`
		Address string `json:"address"`
		Email   string `json:"email" validate:"email"`
		Phone   string `json:"phone" validate:"e164"`
	}
)
