package val

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	isValidLengthPassword                 = regexp.MustCompile(`^.{8,20}$`).MatchString
	isHaveLowercaseCharacter              = regexp.MustCompile(`[a-z]+`).MatchString
	isHaveUppercaseCharacter              = regexp.MustCompile(`[A-Z]+`).MatchString
	isHaveNumberCharacter                 = regexp.MustCompile(`[0-9]+`).MatchString
	isHaveSpecialCharacter                = regexp.MustCompile(`[!@#$%^&*()_+={}\[\]|\\:;"'<,>.?/-]+`).MatchString
	isValidLengthUsername                 = regexp.MustCompile(`^[a-zA-Z0-9][._-a-zA-Z0-9]{6,18}[a-zA-Z0-9]$`).MatchString
	isSpecialCharacterAppearConsecutively = regexp.MustCompile(`[._-][._-]`).MatchString
)

// validUsername checks:
// 1. Username consists of alphanumeric characters (a-zA-Z0-9), lowercase, or uppercase.
// 2. Username allowed of the dot (.), underscore (_), and hyphen (-).
// 3. The dot (.), underscore (_), or hyphen (-) must not be the first or last character.
// 4. The dot (.), underscore (_), or hyphen (-) does not appear consecutively, e.g., java..regex
// 5. The number of characters must be between 8 and 20.
var validUsername validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if username, ok := fieldLevel.Field().Interface().(string); ok {
		return isValidLengthUsername(username) && !isSpecialCharacterAppearConsecutively(username)
	}
	return false
}

// validPassword checks:
// 1. Password must contain at least one digit [0-9].
// 2. Password must contain at least one lowercase Latin character [a-z].
// 3. Password must contain at least one uppercase Latin character [A-Z].
// 4. Password must contain at least one special character like ! @ # & ( ) etc.
// 5. Password must contain a length of at least 8 characters and a maximum of 20 characters.
var validPassword validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if password, ok := fieldLevel.Field().Interface().(string); ok {
		length := isValidLengthPassword(password)
		lower := isHaveLowercaseCharacter(password)
		upper := isHaveUppercaseCharacter(password)
		number := isHaveNumberCharacter(password)
		special := isHaveSpecialCharacter(password)
		return length && lower && upper && number && special
	}
	return false
}
