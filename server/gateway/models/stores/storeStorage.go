package stores

import (
	"database/sql"
	"fmt"
)

type StoreStorage struct {
	DB *sql.DB
}

func NewSqlStorage(sqlsess *sql.DB) *StoreStorage {
	if sqlsess == nil {
		panic("nil sql session")
	}
	return &StoreStorage{sqlsess}
}

var (
	id             int64
	user_id        int64
	store_name     string
	store_location string
)

//GetByID returns the store with the given store ID
func (ss *StoreStorage) GetById(id int64) (*Store, error) {
	row := ss.DB.QueryRow("select * from stores where id=?", id)
	store := &Store{}
	if err := row.Scan(&store.ID, &store.UserID, &store.StoreName, &store.StoreLocation); err != nil {
		return nil, fmt.Errorf("%v: %v", ErrStoreNotFound, err)
	}
	return store, nil
}

//GetStores returns all stores of given user ID
func (ss *StoreStorage) GetStores(userId int64) ([]*Store, error) {
	rows, err := ss.DB.Query("select * from stores where user_id=?", userId)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrStoreNotFound, err)
	}

	defer rows.Close()
	var stores []*Store

	for rows.Next() {
		store := &Store{}
		err := rows.Scan(&store.ID, &store.UserID, &store.StoreName, &store.StoreLocation)
		if err != nil {
			return nil, fmt.Errorf("%v: %v", ErrStoreNotFound, err)
		}
		stores = append(stores, store)
	}
	return stores, nil
}

//Insert new stores and returns inserted store
func (ss *StoreStorage) Insert(store *Store) (*Store, error) {
	insq := "insert into stores(user_id, store_name, store_location) values (?,?,?)"

	res, err := ss.DB.Exec(insq, store.UserID, store.StoreName, store.StoreLocation)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrStoreNotInserted, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrStoreNotInserted, err)
	}

	store.ID = id
	return store, nil
}

//Update applies updates struct store
//and returns the newly-updated store
func (ss *StoreStorage) Update(id int64, updates *StoreUpdate) (*Store, error) {
	updq := "update stores set store_name = ?, store_location = ? where id = ?"
	_, err := ss.DB.Exec(updq, updates.StoreName, updates.StoreLocation, id)
	if updates.StoreName == "" || updates.StoreLocation == "" {
		return nil, fmt.Errorf("%v: %s", ErrStoreNotUpdated, "Contains empty field")
	}

	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrStoreNotUpdated, err)
	}

	return ss.GetById(id)
}

//Delete deletes store with given ID
func (ss *StoreStorage) Delete(id int64) error {
	delq := "delete from stores where id = ?"
	_, err := ss.DB.Exec(delq, id)
	if err != nil {
		return fmt.Errorf("%v: %v", ErrStoreNotDeleted, err)
	}

	return nil
}

//DeleteAllbyUserId deletes every stores related to user
func (ss *StoreStorage) DeleteAllbyUserId(userId int64) error {
	delq := "delete from items where user_id = ?"
	_, err := ss.DB.Exec(delq, userId)
	if err != nil {
		return fmt.Errorf("%v: %v", ErrStoreNotDeleted, err)
	}
	return nil
}
