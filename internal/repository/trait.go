package repository

import (
	"context"
	"time"

	"github.com/muhrifqii/curium_go_fiber/domain"
)

type (
	UserRepository interface {
		GetByEmail(ctx context.Context, email string) (domain.User, error)
		GetByUsername(ctx context.Context, username string) (domain.User, error)
		CreateUser(ctx context.Context, user *domain.User) error
		IsUserExistByIdentifier(c context.Context, email, username string) (bool, error)

		OnUserLoggedIn(ctx context.Context, id int64, time time.Time, ip, ua string) error
	}
)
