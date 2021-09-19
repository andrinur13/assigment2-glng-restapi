package structs

type Item struct {
	ItemId      int    `json:"item_id" gorm:"column:order_id; PRIMARY_KEY"`
	ItemCode    int    `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}
