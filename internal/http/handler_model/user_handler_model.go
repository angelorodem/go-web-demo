package handlermodel

// Create new user model
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Delete user model
type DeleteUserRequest struct {
	Email string `json:"email" binging:"required"`
}

// Login user model
type LoginUserRequest struct {
	Email    string `json:"email" binging:"required"`
	Password string `json:"password" binging:"required"`
}

// Login user model
type GetUserRequest struct {
	Email string `json:"email" binging:"required"`
}

// new username
type ChangeUsernameRequest struct {
	Email       string `json:"email" binging:"required"`
	NewUsername string `json:"newUsername" binging:"required"`
}
