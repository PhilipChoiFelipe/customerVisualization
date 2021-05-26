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
	GetById(itemId int64) (*Item, error)cusomer of current user
	GetCustomers() ([]*Item, error)
itemtitem
	//Inseritems new customer and returns inserted customer
	Insert(customer *Item) (*Item, error)
item
	//Update applies updates struct item
	//and returns the newly-upItemd customer
	Update(id int64, updates *NameUpdates) (*Item, error)
item
	//Delete deletes customer with given ID
	Delete(id int64) error
}
