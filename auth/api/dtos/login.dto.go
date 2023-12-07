package dtos

type LoginReqDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResDto struct {
	Token string
}
