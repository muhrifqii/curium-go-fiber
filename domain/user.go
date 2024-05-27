package domain

import "time"

type User struct {
	BaseModel

	Username string
	Email    string
	Phone    string

	Password string
	Salt     string

	Status    string
	LastLogin time.Time

	FirstName string
	LastName  string
	Birthday  time.Time
	Gender    string
	Avatar    string
}

func NewUser() User {
	return User{
		ID: 1,
	}
}
