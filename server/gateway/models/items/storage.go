package items

import (
	"errors"
)

var ErrItemNotFound = errors.New("item not found")

var ErrItemNotInserted = errors.New("failed to insert new item")

var ErrItemNotUpdated = errors.New("failed to update item")

var ErrItemNotDeleted = errors.New("failed to delete item")

type Storage interface {
	//GetByID returns the item with the given ID
	GetById(itemId int64) (*Item, error)
	
	//GetItems returns all items of given store ID
	GetItems(storeId int64) ([]*Item, error)

	//Insert new items and returns inserted item
	Insert(item *Item) (*Item, error)

	//Update applies updates struct item
	//and returns the newly-updated item
	Update(id int64, updates *ItemUpdate) (*Item, error)

	//Delete deletes item with given ID
	Delete(id int64) error
}
