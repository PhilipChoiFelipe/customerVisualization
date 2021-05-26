package customers

import (
	"errors"
)

var ErrCustomerNotFound = errors.New("customer not found")

var ErrCustomerNotInserted = errors.New("failed to insert new customer")

var ErrCustomerNotUpdated = errors.New("failed to update customer")

var ErrCusomerNotDeleted = errors.New("failed to delete customer")

type Storage interface {
	//GetByID returns the customer with the given ID
	GetById(customerId int64) (*Customer, er
	//GetByItemId returns the customer
	GetByItemId(itemId int64) ([]*Customer, error)

	//GetCustomers returns all the cusomer of current user
	GetCustomers() ([]*Customer, error)

	//Insert inserts new customer and returns inserted customer
	Insert(customer *Customer) (*Customer, error)

	//Update applies updates struct to the given customer ID
	//and returns the newly-updated customer
	Update(id int64, updates *NameUpdates) (*Customer, error)

	//Delete deletes customer with given ID
	Delete(id int64) error
}
