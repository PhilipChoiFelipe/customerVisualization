package items

import (
	"database/sql"
	"fmt"
)

type ItemStorage struct {
	DB *sql.DB
}

func NewSqlStorage(sqlsess *sql.DB) *ItemStorage {
	if sqlsess == nil {
		panic("nil sql session")
	}
	return &ItemStorage{sqlsess}
}

var (
	id        int64
	user_id   int64
	item_name string
	price     int64
)

//GetByID returns the item with the given ID
func (is *ItemStorage) GetById(id int64) (*Item, error) {
	row := is.DB.QueryRow("select * from items where id=?", id)
	item := &Item{}
	if err := row.Scan(&item.ID, &item.UserID, &item.ItemName, &item.Price); err != nil {
		return nil, fmt.Errorf("%v: %v", ErrItemNotFound, err)
	}
	return item, nil
}

// Returns all items with given user ID
func (is *ItemStorage) GetItems(userId int64, queryCase string, col_name string, reverse string) ([]*Item, error) {

	query := "select * from items where user_id = ?"
	// rows, err := is.DB.Query("select * from items where user_id = ?", userId)
	// if err != nil {
	// 	return nil, fmt.Errorf("%v: %v", ErrItemNotFound, err)
	// }

	var values []interface{}
	var items []*Item
	switch queryCase {
	case "default":
		values = append(values, user_id)
	case "sort":
		if reverse == "true" {
			query = "select * from items where user_id = ? order by ? DESC"
		} else {
			query = "select * from items where user_id = ? order by ? ASC"
		}
		values = append(values, user_id, col_name)
	}

	rows, err := is.DB.Query(query, values)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := &Item{}
		err := rows.Scan(&item.ID, &item.UserID, &item.ItemName, &item.Price)
		if err != nil {
			return nil, fmt.Errorf("%v: %v", ErrItemNotFound, err)
		}
		items = append(items, item)
	}
	return items, nil
}

//Insert new items and returns inserted item
func (is *ItemStorage) Insert(item *Item) (*Item, error) {
	insql := "insert into items(user_id, item_name, price) values (?,?,?)"

	res, err := is.DB.Exec(insql, item.UserID, item.ItemName, item.Price)
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

	return is.GetById(id)
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

func (is *ItemStorage) DeleteAllbyUserId(userId int64) error {
	delq := "delete from items where user_id = ?"
	_, err := is.DB.Exec(delq, userId)
	if err != nil {
		return fmt.Errorf("%v: %v", ErrItemNotDeleted, err)
	}
	return nil
}
