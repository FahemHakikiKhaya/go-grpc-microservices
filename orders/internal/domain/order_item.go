package domain

type OrderItem struct {
	MenuID    string  `json:"menu_id" bson:"menu_id"`
	MenuName  string  `json:"menu_name" bson:"menu_name"`
	Quantity  int32   `json:"quantity" bson:"quantity"`
	UnitPrice float64 `json:"unit_price" bson:"unit_price"`
}