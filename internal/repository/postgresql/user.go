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
	_ = `SELECT * FROM users u WHERE u.username ILIKE ? OR u.email ILIKE ?`
	return domain.User{
		BaseModel: domain.BaseModel{
			ID: 1,
		},
		Username: "muhrifqii",
		Email:    "muh_rif@live.com",
		Password: "123123123",
	}, nil
}

func (r *UserRepository) CreateUser(c context.Context, user domain.User) error {
	return nil
}
