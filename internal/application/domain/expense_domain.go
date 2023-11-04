package domain

import "time"

type ExpenseDomain struct {
	ID              string
	Description     string
	Price     		float64
	CompanyName     string
	DateTime     	time.Time
}