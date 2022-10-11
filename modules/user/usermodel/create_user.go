package usermodel

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,password"`
	//Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
