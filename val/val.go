package val

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	tagAccountNumber = "account_number"
	tagUsername      = "username"
	tagPassword      = "password"
)

func RegisterCustomValidation() error {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		err := val.RegisterValidation(tagAccountNumber, validAccountNumber)
		if err != nil {
			return err
		}
		err = val.RegisterValidation(tagUsername, validUsername)
		if err != nil {
			return err
		}
		err = val.RegisterValidation(tagPassword, validPassword)
		if err != nil {
			return err
		}
	}
	return nil
}
