package common

import "e-wallet-app/modules/user/userenum"

type Requester interface {
	GetUsername() string
	GetFullName() string
	GetEmail() string
	GetRole() userenum.Role
}
