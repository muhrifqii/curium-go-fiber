package dto

type (
	AuthnRequest struct {
		Identifier string `json:"identifier" validate:"required"`
		Password   string `json:"password" validate:"required"`
	}

	RegisterByEmailRequest struct {
		Username string `json:"username" validate:"required,alphanum,ne_ignore_case=system"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	RegisterByPhoneRequest struct {
		Phone    string `json:"phone" validate:"required,e164"`
		Password string `json:"password" validate:"required"`
	}
)
