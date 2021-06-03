package customers

import (
	// "fmt"
	
)

type Customer struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Ethnicity   string `json:"ethnicity"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	PostalCode  int64  `json:"postalCode"`
	LastVisited string `json:"lastVisited"`
	DisChannel  string `json:"disChannel"`
	FavItem     int64  `json:"favItem"`
}


type NameUpdates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Updates struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Ethnicity   string `json:"ethnicity"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	PostalCode  int64  `json:"postalCode"`
	LastVisited string `json:"lastVisited"`
	DisChannel  string `json:"disChannel"`
	FavItem     int64  `json:"favItem"`
}

func (c *Customer) ApplyNameUpdates(updates *NameUpdates) error {
	if updates.FirstName == "" || updates.LastName == "" {
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

// func (c *Customer) ApplyLastVistedUpdates(date string) error {
// 	if len(date) == 0 {
// 		return fmt.Errorf("empty updating value")
// 	}
// 	c.LastVisited = date
// 	return nil
// }
