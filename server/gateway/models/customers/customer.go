package customers

import (
	"fmt"
	// "golang.org/x/text/date"
)

type Customer struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"userId"`
	StoreID    int64  `json:"storeId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Ethnicity  string `json:"ethnicity"`
	Gender     string `json:"gender"`
	Birthday   string `      json:"birthday"` // TOASK: better datatype?
	PostalCode int64  `json:"postalCode"`
	LastVisted string `json:"lastVisited"`
	DisChannel string `json:"disChannel"`
	FavItem    int64  `json:"favItem"` // TOASK: item struct type better as json object? wbout sql?
}

// https://stackoverflow.com/questions/47335697/golang-decode-json-request-in-nested-struct-and-insert-in-db-as-blob

type NameUpdates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (c *Customer) ApplyNameUpdates(updates *NameUpdates) error {
	if updates.FirstName == "" || updates.LastName == "" {
		return fmt.Errorf("empty updating value")
	}
	c.FirstName = updates.FirstName
	c.LastName = updates.LastName
	return nil
}

// // TOASK: do we need update struct?
// func (c *Customer) ApplyFavItemUpdates(item *items.Item) error {
// 	if item == nil {
// 		return fmt.Errorf("empty updating value")
// 	}
// 	c.FavItem = item
// 	return nil
// }

func (c *Customer) ApplyLastVistedUpdates(date string) error {
	if len(date) == 0 {
		return fmt.Errorf("empty updating value")
	}
	c.LastVisted = date
	return nil
}
