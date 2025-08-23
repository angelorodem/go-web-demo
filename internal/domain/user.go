package domain

type User struct {
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password_hash string `json:"-"`
}
