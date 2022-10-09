package val

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	tagAccountNumber = "account_number"
)

func RegisterCustomValidation() error {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		err := val.RegisterValidation(tagAccountNumber, validAccountNumber)
		if err != nil {
			return err
		}
	}
	return nil
}
