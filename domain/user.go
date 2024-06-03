package domain

import "time"

type User struct {
	BaseModel

	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`

	Password string `json:"-"`
	Salt     string `json:"-"`

	Status      string    `json:"-"`
	LastLogin   time.Time `json:"-" db:"last_login"`
	LastLoginIp string    `json:"-" db:"last_login_ip"`

	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Birthday  time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Avatar    string    `json:"avatar"`
}

const (
	UserStatusActive   = "active"
	UserStatusPending  = "pending"
	UserStatusInactive = "inactive"
	UserStatusBanned   = "banned"
)
