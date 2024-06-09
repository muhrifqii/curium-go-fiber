package domain

import "time"

type User struct {
	BaseModel

	Username     string `json:"username"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`

	Password            string     `json:"-" db:"a_password"`
	LastPasswordUpdated *time.Time `json:"-" db:"last_password_updated"`

	Status string `json:"-" db:"a_status"`

	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Birthday  *time.Time `json:"birth_date"`
	Gender    string     `json:"gender"`
	Avatar    string     `json:"avatar"`
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
