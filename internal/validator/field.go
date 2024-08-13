package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomFieldValidators(validate *validator.Validate) {
	registerJsonTags(validate)
	validate.RegisterValidation("x_username", validateUsername)
	validate.RegisterValidation("x_username_or_email", validateUsernameOrEmail(validate))
}

func registerJsonTags(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9._]{3,18}$`)
	return re.MatchString(username)
}

func validateUsernameOrEmail(validate *validator.Validate) validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		emailErr := validate.Var(value, "email")
		return validateUsername(fl) || emailErr == nil
	}
}
