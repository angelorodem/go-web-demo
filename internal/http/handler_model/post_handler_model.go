package handlermodel

// Create new post
type CreatePostRequest struct {
	UserEmail string `json:"userEmail" binding:"required"` // We use this as mock to get the user ID since our token does not hold claims
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

// Update the post
type UpdatePostRequest struct {
	Id         int    `json:"id" binding:"required"`
	UserEmail  string `json:"userEmail" binding:"required"` // We use this as mock to get the user ID since our token does not hold claims
	NewTitle   string `json:"newTitle" binding:"required"`
	NewContent string `json:"newContent" binding:"required"`
}

// Read the post (all posts are public no need email/claim)
type ReadPostRequest struct {
	Id int `json:"id" binding:"required"`
}

// Delete the post
type DeletePostRequest struct {
	Id        int    `json:"id" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"` // We use this as mock to get the user ID since our token does not hold claims
}
