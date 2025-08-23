package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	hm "web/example/internal/http/handler_model"
	"web/example/internal/repository"
	"web/example/internal/services"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		db,
	}
}

// Create new user
func (uh *UserHandler) Create(c *gin.Context) {
	var req hm.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// since creating a user is long we create a service for it.
	if err := services.CreateUserService(&req, uh.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusAccepted)
}

func (uh *UserHandler) Delete(c *gin.Context) {
	var req hm.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := repository.DeleteUser(uh.DB, req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusAccepted)
}

func (uh *UserHandler) Login(c *gin.Context) {
	var req hm.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// since logging a user is long we create a service for it.
	if token, err := services.LoginUserService(&req, uh.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (uh *UserHandler) Get(c *gin.Context) {
	var req hm.GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if usr, err := repository.GetUser(uh.DB, req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, usr)
	}

}

func (uh *UserHandler) ChangeUsername(c *gin.Context) {
	var req hm.ChangeUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := repository.UpdateUsername(uh.DB, req.Email, req.NewUsername)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusAccepted)
}
