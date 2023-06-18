package dto

type RegisterDto struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
