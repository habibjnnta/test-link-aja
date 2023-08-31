package models

type Customer struct {
	CustomerNumber string  `gorm:"type:varchar(7);primaryKey;" json:"customer_number"`
	Name           string  `json:"customer_name"`
	Accounts       Account `gorm:"foreignKey:customer_number"`
}
