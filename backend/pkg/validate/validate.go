package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

type Validator struct {
	Validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		Validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

var ProviderSet = wire.NewSet(NewValidator)
