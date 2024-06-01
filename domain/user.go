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
	LastLogin   time.Time `json:"-"`
	LastLoginIp string    `json:"-"`

	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Avatar    string    `json:"avatar"`
}
