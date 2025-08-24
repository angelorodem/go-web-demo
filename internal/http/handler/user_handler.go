package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	hm "web/example/internal/http/handler_model"
	"web/example/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(db),
	}
}

// Create new user
func (uh *UserHandler) Create(c *gin.Context) {
	var req hm.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// since creating a user is long we create a service for it.
	if err := uh.userService.CreateUserService(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

func (uh *UserHandler) Delete(c *gin.Context) {
	var req hm.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uh.userService.DeleteUser(req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

func (uh *UserHandler) Login(c *gin.Context) {
	var req hm.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// since logging a user is long we create a service for it.
	if token, err := uh.userService.LoginUserService(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (uh *UserHandler) Get(c *gin.Context) {
	var req hm.GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if usr, err := uh.userService.ReadUser(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, usr)
	}

}

func (uh *UserHandler) ChangeUsername(c *gin.Context) {
	var req hm.ChangeUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uh.userService.UpdateUsername(req.Email, req.NewUsername)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}
