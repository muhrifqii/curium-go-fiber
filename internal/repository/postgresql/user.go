package postgresql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/muhrifqii/curium_go_fiber/domain"
)

type (
	UserRepository struct {
		DB *sqlx.DB
	}
)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetByIdentifier(c context.Context, identifier string) (domain.User, error) {
	return domain.User{
		BaseModel: domain.BaseModel{
			ID: 1,
		},
		Username: identifier,
		Email:    "muh_rif@live.com",
		Password: "123123123",
	}, nil
}

func (r *UserRepository) CreateUser(c context.Context, user domain.User) error {
	_, err := r.DB.NamedExec("INSERT INTO user (username, email, phone, password, status, first_name, last_name) VALUES (:username, :email, :phone, :password, :status, :first_name, :last_name)", user)
	return err
}

func (r *UserRepository) IsUserExistByIdentifier(c context.Context, email, username string) (bool, error) {
	var exist bool
	err := r.DB.Get(&exist, `SELECT EXISTS(SELECT 1 FROM users u WHERE LOWER(u.username) = LOWER($1) OR LOWER(u.email) = LOWER($1))`, username, email)
	if err != nil {
		return false, err
	}
	return true, nil
}
