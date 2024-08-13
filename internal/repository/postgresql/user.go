package postgresql

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/sqler"
	"go.uber.org/zap"
)

type (
	UserRepository struct {
		db *sqler.SqlxWrapper
	}
)

func NewUserRepository(db *sqlx.DB, zap *zap.Logger) *UserRepository {
	return &UserRepository{
		db: sqler.NewSqlxWrapper(db, zap),
	}
}

func (r *UserRepository) GetByID(c context.Context, id int64) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	return user, err
}

func (r *UserRepository) GetByUsername(c context.Context, username string) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	return user, err
}

func (r *UserRepository) GetByOrgUsername(c context.Context, orgID, username string) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1 AND organization_id = $2", username, orgID)
	return user, err
}

func (r *UserRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}

func (r *UserRepository) GetByOrgEmail(c context.Context, orgID, email string) (domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1 AND organization_id = $2", email, orgID)
	return user, err
}

func (r *UserRepository) CreateUser(c context.Context, user *domain.User) error {

	_, err := r.db.NamedExec("INSERT INTO users (organization_id, username, email, phone, a_password, a_status, first_name, last_name) VALUES (:organization_id, :username, :email, :phone, :a_password, :a_status, :first_name, :last_name)", user)

	return err
}

func (r *UserRepository) IsOrgUserExistByIdentifier(c context.Context, orgID, email, username string) (bool, error) {
	var exist bool
	err := r.db.Get(&exist, `SELECT EXISTS(SELECT 1 FROM users u WHERE u.organization_id = $1 AND (LOWER(u.username) = LOWER($2) OR LOWER(u.email) = LOWER($3)))`, orgID, username, email)
	return exist, err
}

func (r *UserRepository) OnUserLoggedIn(c context.Context, id int64, time time.Time, ip, ua string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_ = r.db.MustExec("INSERT INTO user_login_history (user_id, login_time, ip_address, user_agent) VALUES ($1, $2, $3, $4)", id, time, ip, ua)
	_ = r.db.MustExec("UPDATE users SET last_login = $1 WHERE id = $2", time, id)

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(c context.Context, user *domain.User) error {
	_, err := r.db.NamedExec("UPDATE users SET username = :username, email = :email, phone = :phone, first_name = :first_name, last_name = :last_name, gender = :gender, avatar = :avatar, updated_at = NOW() WHERE id = :id", user)
	return err
}
