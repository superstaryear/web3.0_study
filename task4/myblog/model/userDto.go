package model

type RegisterRequest struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"required,min=3,max=6" msg:"密码不能为空且长度应为3-6位"`
	Email    string `json:"email" binding:"required,email" msg:"邮箱地址不能为空且格式不正确"`
}

type RegsiterResponse struct {
	UserId   uint
	Username string
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"required,min=3,max=6" msg:"密码不能为空且长度应为3-6位"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
