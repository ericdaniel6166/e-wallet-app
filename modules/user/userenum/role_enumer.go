// Code generated by "enumer -type=Role -json -sql -transform=snake-upper -output=role_enumer.go"; DO NOT EDIT.

package userenum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	_RoleName_0      = "ROLE_ANONYMOUSROLE_USERROLE_SELLER"
	_RoleLowerName_0 = "role_anonymousrole_userrole_seller"
	_RoleName_1      = "ROLE_ADMIN"
	_RoleLowerName_1 = "role_admin"
)

var (
	_RoleIndex_0 = [...]uint8{0, 14, 23, 34}
	_RoleIndex_1 = [...]uint8{0, 10}
)

func (i Role) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _RoleName_0[_RoleIndex_0[i]:_RoleIndex_0[i+1]]
	case i == 4:
		return _RoleName_1
	default:
		return fmt.Sprintf("Role(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _RoleNoOp() {
	var x [1]struct{}
	_ = x[RoleAnonymous-(0)]
	_ = x[RoleUser-(1)]
	_ = x[RoleSeller-(2)]
	_ = x[RoleAdmin-(4)]
}

var _RoleValues = []Role{RoleAnonymous, RoleUser, RoleSeller, RoleAdmin}

var _RoleNameToValueMap = map[string]Role{
	_RoleName_0[0:14]:       RoleAnonymous,
	_RoleLowerName_0[0:14]:  RoleAnonymous,
	_RoleName_0[14:23]:      RoleUser,
	_RoleLowerName_0[14:23]: RoleUser,
	_RoleName_0[23:34]:      RoleSeller,
	_RoleLowerName_0[23:34]: RoleSeller,
	_RoleName_1[0:10]:       RoleAdmin,
	_RoleLowerName_1[0:10]:  RoleAdmin,
}

var _RoleNames = []string{
	_RoleName_0[0:14],
	_RoleName_0[14:23],
	_RoleName_0[23:34],
	_RoleName_1[0:10],
}

// RoleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RoleString(s string) (Role, error) {
	if val, ok := _RoleNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _RoleNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Role values", s)
}

// RoleValues returns all values of the enum
func RoleValues() []Role {
	return _RoleValues
}

// RoleStrings returns a slice of all String values of the enum
func RoleStrings() []string {
	strs := make([]string, len(_RoleNames))
	copy(strs, _RoleNames)
	return strs
}

// IsARole returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Role) IsARole() bool {
	for _, v := range _RoleValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Role
func (i Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Role
func (i *Role) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Role should be a string, got %s", data)
	}

	var err error
	*i, err = RoleString(s)
	return err
}

func (i Role) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Role) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of Role: %[1]T(%[1]v)", value)
	}

	val, err := RoleString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
