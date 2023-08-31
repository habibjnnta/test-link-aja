package models

type Account struct {
	AccountNumber  string `gorm:"type:varchar(7);primaryKey;" json:"account_number"`
	Balance        int    `json:"balance"`
	CustomerNumber string `gorm:"type:varchar(7);" json:"customer_number"`
}
