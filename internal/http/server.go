package http

import (
	"database/sql"
	"net/http"
	"web/example/internal/http/handler"
	"web/example/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer(db_connection *sql.DB) {
	r := gin.Default()

	user_handler := handler.NewUserHandler(db_connection)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user", user_handler.Create)
	r.DELETE("/user", middleware.RequireMockToken(), user_handler.Delete)
	r.GET("/user", middleware.RequireMockToken(), user_handler.Get)
	r.PATCH("/user", middleware.RequireMockToken(), user_handler.ChangeUsername)

	r.POST("/user/login", user_handler.Login)

	r.Run()
}
