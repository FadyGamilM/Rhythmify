package dtos

type SignupReqDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignupResDto struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type LoginReqDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResDto struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}
