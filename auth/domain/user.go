package domain

type User struct {
	Id             int64  `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}
