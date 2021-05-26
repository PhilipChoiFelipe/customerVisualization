package customers

import (
	"fmt"
	"time"

	// "golang.org/x/text/date"
	"final-project/server/gateway/models/items"
)

type Customer struct {
	ID         int64      `json:"id"`
	UserID     int64      `json:"-"` //never JSON encoded/decoded
	StoreID    int64      `json:"-"` //never JSON encoded/decoded
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Ethnicity  string     `json:"ethnicity"`
	Gender     string     `json:"gender"`
	Birthday   time.Time  `json:"birthday"` // TODO: better datatype?
	PostalCode int64      `json:"postalCode"`
	LastVisted time.Time  `json:"lastVisited"`
	DisChannel string     `json:"disChannel"`
	FavItem    items.Item `json:"favItem"` // TODO: item struct type better as json object? wbout sql?
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

// TODO: do we need update struct?
func (c *Customer) ApplyFavItemUpdates(item *items.Item) error {
	if item == nil {
		return fmt.Errorf("empty updating value")
	}
	c.FavItem = item
	return nil
}

func (c *Customer) ApplyLastVistedUpdates(date time.Time) error {
	if date.IsZero() {
		return fmt.Errorf("empty updating value")
	}
	c.LastVisted = date
	return nil
}
