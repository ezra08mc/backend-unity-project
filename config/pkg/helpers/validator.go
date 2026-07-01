package helpers

import (
	"strings"
	"time"

	"github.com/ezra08mc/backend-unity-project/config/pkg/errs"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(payload any) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(payload)
	if err != nil {
		return errs.BadRequest(err.Error())
	}

	return nil
}

func Choose(newVal, oldVal string) string {
	if strings.TrimSpace(newVal) != "" {
		return newVal
	}
	return oldVal
}

func ChooseTime(newValue, oldValue time.Time) time.Time {
	if !newValue.IsZero() {
		return newValue
	}
	return oldValue
}
