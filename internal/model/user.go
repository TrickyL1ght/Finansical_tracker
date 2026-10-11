package model

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Login    string `json:"login" binding:"required" gorm:"unique"`
	Password string `json:"password"`
}

type AuthorisationRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
