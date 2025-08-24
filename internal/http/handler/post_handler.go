package handler

import (
	"database/sql"
	"net/http"
	handlermodel "web/example/internal/http/handler_model"
	"web/example/internal/services"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(db *sql.DB) *PostHandler {
	return &PostHandler{
		postService: services.NewPostService(db),
	}
}

func (np *PostHandler) Create(c *gin.Context) {
	var req handlermodel.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := np.postService.CreatePostService(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)
}

func (np *PostHandler) Delete(c *gin.Context) {
	var req handlermodel.DeletePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := np.postService.DeletePostService(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)
}

func (np *PostHandler) Update(c *gin.Context) {
	var req handlermodel.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := np.postService.UpdatePostService(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)
}

func (np *PostHandler) Read(c *gin.Context) {
	var req handlermodel.ReadPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if post, err := np.postService.ReadPost(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, post)
	}
}

func (np *PostHandler) ReadAll(c *gin.Context) {
	if posts, err := np.postService.ReadAllPosts(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, posts)
	}
}
