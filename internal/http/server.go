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
	post_handler := handler.NewPostHandler(db_connection)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// User Handling
	r.POST("/user", user_handler.Create)
	r.DELETE("/user", middleware.RequireMockToken(), user_handler.Delete) // Will also delete all user posts
	r.GET("/user", middleware.RequireMockToken(), user_handler.Get)
	r.PATCH("/user", middleware.RequireMockToken(), user_handler.ChangeUsername)

	// Login handling
	r.POST("/user/login", user_handler.Login)

	// Post handling
	// posts could be accessed also by using
	// `/post/{post_id}` but i prefere to use full json approach
	r.POST("/post", post_handler.Create)
	r.DELETE("/post", middleware.RequireMockToken(), post_handler.Delete)
	r.GET("/post", post_handler.Read) // posts are public
	r.PUT("/post", middleware.RequireMockToken(), post_handler.Update)

	r.GET("/post/all", post_handler.ReadAll) // posts are public

	r.Run()
}
