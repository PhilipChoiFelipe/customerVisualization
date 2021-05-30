package stores

import (
	"errors"
)

var ErrStoreNotFound = errors.New("store not found")

var ErrStoreNotInserted = errors.New("failed to insert new store")

var ErrStoreNotUpdated = errors.New("failed to update store")

var ErrStoreNotDeleted = errors.New("failed to delete store")

type Storage interface {
	//GetByID returns the store with the given ID
	GetById(storeId int64) (*Store, error)

	//GetStores returns all stores of given user ID
	GetStores(userId int64) ([]*Store, error)

	//Insert new stores and returns inserted store
	Insert(store *Store) (*Store, error)

	//Update applies updates struct store
	//and returns the newly-updated store
	Update(id int64, updates *StoreUpdate) (*Store, error)

	//Delete deletes store with given ID
	Delete(id int64) error

	//DeleteAllbyUserId deletes all customers reltated to user_id
	DeleteAllbyUserId(user_id int64) error
}
