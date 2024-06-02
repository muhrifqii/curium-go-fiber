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
	return nil
}

func (r *UserRepository) IsUserExistByIdentifier(c context.Context, identifier stirng) (domain.User, error) {
	q := r.DB.Query(`SELECT EXISTS(SELECT 1 FROM users u WHERE LOWER(u.username) = LOWER(?) OR LOWER(u.email) = LOWER(?))`, identifier)

}
