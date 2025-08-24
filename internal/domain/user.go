package domain

type User struct {
	Id            int    `json:"-"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password_hash string `json:"-"`
}
