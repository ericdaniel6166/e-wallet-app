package db

import "e-wallet-app/modules/user/userenum"

func (u User) GetUsername() string {
	return u.Username
}

func (u User) GetFullName() string {
	return u.FullName
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetRole() userenum.Role {
	return userenum.Role(u.Role)
}
