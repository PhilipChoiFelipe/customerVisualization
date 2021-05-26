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

//GetByID returns the item with the given ID
func (is *ItemStorage) GetByID(id int64) (*Item, error) {
	row := is.DB.QueryRow("select * from items where id=?", id)
	item := &Item{}
	if err := row.Scan(&item.ID, &item.StoreID, &item.ItemName, &item.Price); err != nil {
		return nil, fmt.Errorf("%v: %v", ErrItemNotFound, err)
	}
	return item, nil
}

// Returns all items with given store ID
func (is *ItemStorage) GetItems(storeId int64) ([]*Item, error) {
	rows, err := is.DB.Query("select * from items")
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrItemNotFound, err)
	}

	defer rows.Close()
	var items []*Item

	for rows.Next() {
		item := &Item{}
		err := rows.Scan(&item.ID, &item.StoreID, &item.ItemName, &item.Price)
		if err != nil {
			return nil, fmt.Errorf("%v: %v", ErrItemNotFound, err)
		}
		items = append(items, item)
	}
	return items, nil
}

//Insert new items and returns inserted item
func (is *ItemStorage) Insert(item *Item) (*Item, error) {
	insq := "insert into items(store_id, item_name, price) values (?,?,?)"

	res, err := is.DB.Exec(insq, item.StoreID, item.ItemName, item.Price)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrItemNotInserted, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrItemNotInserted, err)
	}

	item.ID = id
	return item, nil
}

//Update applies updates struct item
//and returns the newly-updated item
func (is *ItemStorage) Update(id int64, updates *ItemUpdate) (*Item, error) {
	updq := "update items set item_name = ?, price = ? where id = ?"
	_, err := is.DB.Exec(updq, updates.ItemName, updates.Price, id)
	if updates.ItemName == "" || updates.Price == 0 {
		return nil, fmt.Errorf("%v: %s", ErrItemNotUpdated, "Contains empty field")
	}

	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrItemNotUpdated, err)
	}

	return is.GetByID(id)
}

//Delete deletes item with given ID
func (is *ItemStorage) Delete(id int64) error {
	delq := "delete from items where id = ?"
	_, err := is.DB.Exec(delq, id)
	if err != nil {
		return fmt.Errorf("%v: %v", ErrItemNotDeleted, err)
	}

	return nil
}
