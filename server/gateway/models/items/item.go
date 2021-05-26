package items

type Item struct {
	ID       int64  `json:"id"`
	StoreID  int64  `json:"storeId"`
	ItemName string `json:"itemName"`
	Price    int64  `json:"price"`
}

type ItemUpdate struct {
	ItemName string `json:"itemName"`
	Price    int64  `json:"price"`
}
