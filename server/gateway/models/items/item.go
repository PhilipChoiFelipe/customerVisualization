package items

import "fmt"

// Item represents a product or service provided in a store
type Item struct {
	ID       int64  `json:"id"`
	StoreID  int64  `json:"storeId"`
	UserID   int64  `json:"userId"`
	ItemName string `json:"itemName"`
	Price    int64  `json:"price"`
}

type ItemUpdate struct {
	ItemName string `json:"itemName"`
	Price    int64  `json:"price"`
}

func (i *Item) ApplyUpdates(updates *ItemUpdate) error {
	if updates.ItemName == "" || updates.Price == 0 {
		return fmt.Errorf("empty updating value")
	}
	i.ItemName = updates.ItemName
	i.Price = updates.Price
	return nil
}
