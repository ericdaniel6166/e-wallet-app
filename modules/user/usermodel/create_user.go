package usermodel

import (
	"e-wallet-app/modules/user/userenum"
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,password"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Role     int32
}

type CreateUserResponse struct {
	Username  string        `json:"username"`
	FullName  string        `json:"full_name"`
	Email     string        `json:"email"`
	Role      userenum.Role `json:"role"`
	UpdatedAt time.Time     `json:"updated_at"`
	CreatedAt time.Time     `json:"created_at"`
}

func (req *CreateUserRequest) FillDefault() {
	req.Role = int32(userenum.RoleAnonymous)
}
