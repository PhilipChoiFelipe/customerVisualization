package items

import (
	"database/sql"
	"fmt"
)

type ItemStorage struct {
	DB *sql.DB
}

var (
	id        int64
	store_id  int64
	item_name string
	price     int64
)

func (is *ItemStorage) GetByID(id int64) (*Item, error) {
	row := is.DB.QueryRow("select * from items where id=?", id)
	item := &Item{}
	if err := row.Scan(&item.ID, &item.StoreID, &item.ItemName, &item.Price); err != nil {
		return nil, fmt.Errorf("%v: %v", ErrItemNotFound, err)
	}

	return item, nil
}
