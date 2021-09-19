package structs

import "time"

type Order struct {
	OrderId      int       `json:"order_id" gorm:"column:order_id; PRIMARY_KEY"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}
