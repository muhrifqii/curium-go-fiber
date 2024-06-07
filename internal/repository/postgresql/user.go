package postgresql

import (
	"context"
	"time"

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

func (r *UserRepository) GetByUsername(c context.Context, username string) (domain.User, error) {
	var user domain.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	return user, err
}

func (r *UserRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}

func (r *UserRepository) CreateUser(c context.Context, user *domain.User) error {

	_, err := r.DB.NamedExec("INSERT INTO users (username, email, phone, a_password, a_status, first_name, last_name) VALUES (:username, :email, :phone, :a_password, :a_status, :first_name, :last_name)", user)

	return err
}

func (r *UserRepository) IsUserExistByIdentifier(c context.Context, email, username string) (bool, error) {
	var exist bool
	err := r.DB.Get(&exist, `SELECT EXISTS(SELECT 1 FROM users u WHERE LOWER(u.username) = LOWER($1) OR LOWER(u.email) = LOWER($2))`, username, email)
	return exist, err
}

func (r *UserRepository) OnUserLoggedIn(c context.Context, id int64, time time.Time, ip, ua string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_ = r.DB.MustExec("INSERT INTO user_login_history (user_id, login_time, ip_address, user_agent) VALUES ($1, $2, $3, $4)", id, time, ip, ua)
	_ = r.DB.MustExec("UPDATE users SET last_login = $1 WHERE id = $2", time, id)

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
