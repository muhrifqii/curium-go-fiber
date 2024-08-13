package domain

import (
	"time"
)

type OAuthClient struct {
	BaseModel
	ClientID               string `db:"client_id" json:"client_id" validate:"required"`
	ClientSecret           string `db:"client_secret" json:"client_secret" validate:"required"`
	ClientName             string `db:"client_name" json:"client_name" validate:"required"`
	GrantTypes             string `db:"grant_types" json:"grant_types" validate:"required"`
	Scope                  string `db:"scope" json:"scope"`
	RedirectURIs           string `db:"redirect_uris" json:"redirect_uris" validate:"required"`
	PostLogoutRedirectURIs string `db:"post_logout_redirect_uris" json:"post_logout_redirect_uris"`
}

type OAuthAuthorizationCode struct {
	Code                string    `db:"code"`
	UserID              int64     `db:"user_id"`
	ClientID            string    `db:"client_id"`
	CodeChallenge       string    `db:"code_challenge"`
	CodeChallengeMethod string    `db:"code_challenge_method"`
	RedirectURI         string    `db:"redirect_uri"`
	ExpiresAt           time.Time `db:"expires_at"`
	CreatedAt           time.Time `db:"created_at"`
}

type OAuthRefreshToken struct {
	Token     string    `db:"token"`
	UserID    int64     `db:"user_id"`
	ClientID  string    `db:"client_id"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
}

type OAuthOidcClaim struct {
	UserID int64  `db:"user_id"`
	Name   string `db:"claim_name"`
	Value  string `db:"claim_value"`
}

type (
	OAuthAuthorizeRequest struct {
		ClientID            string `form:"client_id" validate:"required"`
		RedirectURI         string `form:"redirect_uri" validate:"required"`
		CodeChallenge       string `form:"code_challenge"`
		CodeChallengeMethod string `form:"code_challenge_method"`
		State               string `form:"state"`
		ResponseType        string `form:"response_type" validate:"required,oneof=code token id_token"`

		LoginRequest
	}

	OAuthTokenRequest struct {
		GrantType string `form:"grant_type" validate:"required,oneof=authorization_code refresh_token password"`
		ClientID  string `form:"client_id" validate:"required"`
		Code      string `form:"code" validate:"required"`

		CodeVerifier string `form:"code_verifier"`

		RefreshToken string `form:"refresh_token"`

		LoginRequest
	}

	OAuthTokenResponse struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenType    string `json:"token_type"`
		IDToken      string `json:"id_token"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}

	LoginRequest struct {
		Identifier string `form:"identifier"`
		OrgID      string `form:"org_id"`
		Password   string `form:"password"`
	}

	UserInfoResponse struct {
		Sub     int64  `json:"sub"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		OrgID   string `json:"org"`
		Profile string `json:"profile"`
	}
)
