package transaction

type Transaction struct {
	Id               int     `json:"id" gorm:"primary_key"`
	Transaction_type string  `json:"type" binding:"required,oneof=income expense"`
	Amount           float64 `json:"amount" binding:"required,gt=0"`
	Category         string  `json:"category" binding:"required"`
	Description      string  `json:"description"`
	Date             int64   `json:"date" binding:"max=2147483647"`
}
