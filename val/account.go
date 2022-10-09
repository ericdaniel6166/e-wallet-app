package val

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	isValidAccountNumber = regexp.MustCompile(`^[0-9]+$`).MatchString
)

var validAccountNumber validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if acctNo, ok := fieldLevel.Field().Interface().(string); ok {
		return isValidAccountNumber(acctNo)
	}
	return false
}
