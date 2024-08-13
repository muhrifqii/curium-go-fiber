package domain

import "time"

type User struct {
	BaseOrganizationModel

	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`

	Password            string     `json:"-" db:"a_password"`
	LastPasswordUpdated *time.Time `json:"-" db:"last_password_updated"`
	LastLogin           *time.Time `json:"-" db:"last_login"`

	Status string `json:"-" db:"a_status"`

	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Birthday  *time.Time `json:"birth_date"`
	Gender    *string    `json:"gender"`
	Avatar    *string    `json:"avatar"`
}

type UserLoginHistory struct {
	ID        int64      `json:"-" db:"id"`
	UserID    int64      `json:"-" db:"user_id"`
	IpAddress string     `json:"ip_address" db:"ip_address"`
	UserAgent string     `json:"user_agent" db:"user_agent"`
	LoginTime *time.Time `json:"login_time" db:"login_time"`
}

const (
	UserStatusActive   = "active"
	UserStatusPending  = "pending"
	UserStatusInactive = "inactive"
	UserStatusBanned   = "banned"

	UserStatusSystem = "system"
)

type OAuthProvider struct {
	BaseModel
	Name        string `json:"name" db:"provider_name"`
	DisplayName string `json:"display_name" db:"provider_display_name"`
}

type UserOauthAccount struct {
	BaseModel
	UserID          int64  `json:"-" db:"user_id"`
	OAuthProviderID int64  `json:"-" db:"provider_id"`
	OAuthID         string `json:"-" db:"oauth_id"`
	Email           string `json:"-" db:"email"`
	RefreshToken    string `json:"-" db:"refresh_token"`
}

type (
	AuthnRequest struct {
		// Organization string `json:"organization" validate:"required,ascii"`
		Identifier string `json:"identifier" validate:"required,x_username_or_email"`
		Password   string `json:"password" validate:"required"`
	}

	RegisterByEmailRequest struct {
		Username string `json:"username" validate:"required,ne_ignore_case=system,x_username"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	RegisterByPhoneRequest struct {
		Phone    string `json:"phone" validate:"required,e164"`
		Password string `json:"password" validate:"required"`
	}

	AuthnResponse struct {
		AccessToken           string    `json:"access_token"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		RefreshToken          string    `json:"-"`
		RefreshTokenExpiresAt time.Time `json:"-"`
	}
)

type (
	CreateUserRequest struct {
		Username  string  `json:"username" validate:"required,ne_ignore_case=system,x_username"`
		Email     string  `json:"email" validate:"required,email"`
		Password  string  `json:"password" validate:"required"`
		Phone     string  `json:"phone" validate:"e164"`
		FirstName string  `json:"first_name" validate:"required,ascii"`
		LastName  string  `json:"last_name" validate:"ascii"`
		Gender    *string `json:"gender" validate:"oneof=male female"`
	}

	UpdateUserRequest struct {
		FirstName string  `json:"first_name" validate:"required,ascii"`
		LastName  string  `json:"last_name" validate:"ascii"`
		Gender    *string `json:"gender" validate:"oneof=male female"`
		Avatar    *string `json:"avatar"`
	}
)
