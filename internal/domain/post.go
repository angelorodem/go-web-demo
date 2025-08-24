package domain

type Post struct {
	Id        int    `json:"id"`
	UserId    int    `json:"-"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	CreatedAt string `json:"createdAt"`
}
