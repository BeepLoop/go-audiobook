package types

type Admin struct {
	Username string `json:"username" form:"username" db:"username" binding:"required"`
	Password string `json:"password" form:"password" db:"password" binding:"required"`
}
