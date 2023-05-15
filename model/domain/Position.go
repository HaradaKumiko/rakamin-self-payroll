package domain

type Position struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Salary int64  `json:"salary"`
}
