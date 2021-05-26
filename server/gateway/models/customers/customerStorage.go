package customers

import (
	"database/sql"
	"time"

	"github.com/info441-sp21/final-project/server/gateway/models/items"
)

type CustomerStorage struct {
	sqlsess *sql.DB
}

//Customers Storage implementation

//NewSqlStorage returns new sql connection instance
func NewSqlStorage(sqlsess *sql.DB) *CustomerStorage {
	if sqlsess == nil {
		panic("nil sql session")
	}
	return &CustomerStorage{sqlsess}
}

var (
	ID         int64
	UserID     int64
	StoreID    int64
	FirstName  string
	LastName   string
	Ethnicity  string
	Gender     string
	Birthday   time.Time
	PostalCode int64
	LastVisted time.Time
	DisChannel string
	ItemId     int64
	FavItem    *items.Item
	result     []*Customer
)

//GetById finds id of customers in DB and returns the customer
func (cs *CustomerStorage) GetById(customerId int64) (*Customer, error) {

	query := "select id, user_id, store_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item from customers where id = ?"
	err := cs.sqlsess.QueryRow(query, customerId).Scan(&ID, &UserID, &StoreID, &FirstName, &LastName, &Ethnicity, &Gender, &Birthday, &PostalCode, &LastVisted, &DisChannel, &ItemId)
	if err != nil {
		return nil, err
	}
	itemQuery := "select * from items where id = ?"
	err = cs.sqlsess.QueryRow(itemQuery, ItemId).Scan(&FavItem)
	if err != nil {
		return nil, err
	}
	return &Customer{ID, UserID, StoreID, FirstName, LastName, Ethnicity, Gender, Birthday, PostalCode, LastVisted, DisChannel, FavItem}, nil
}

//GetByItemId finds itemId of customers in DB and returns the customers with certain favorite ids
func (cs *CustomerStorage) GetByItemId(itemId int64) ([]*Customer, error) {
	query := "select id, user_id, store_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item from customers where fav_item = ?"
	rows, err := cs.sqlsess.Query(query, itemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&ID, &UserID, &StoreID, &FirstName, &LastName, &Ethnicity, &Gender, &Birthday, &PostalCode, &LastVisted, &DisChannel, &ItemId)
		itemQuery := "select * from items where id = ?"
		err = cs.sqlsess.QueryRow(itemQuery, itemId).Scan(&FavItem)
		if err != nil {
			return nil, err
		}
		returnedCus := &Customer{ID, UserID, StoreID, FirstName, LastName, Ethnicity, Gender, Birthday, PostalCode, LastVisted, DisChannel, FavItem}
		result = append(result, returnedCus)
	}
	if len(result) == 0 {
		return nil, ErrCustomerNotFound
	}
	return result, nil
}

//GetByItemId returns the all customers in DB
func (cs *CustomerStorage) GetCustomers() ([]*Customer, error) {
	query := "select id, user_id, store_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item from customers"
	rows, err := cs.sqlsess.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&ID, &UserID, &StoreID, &FirstName, &LastName, &Ethnicity, &Gender, &Birthday, &PostalCode, &LastVisted, &DisChannel, &ItemId)
		itemQuery := "select * from items where id = ?"
		err = cs.sqlsess.QueryRow(itemQuery, ItemId).Scan(&FavItem)
		if err != nil {
			return nil, err
		}
		returnedCus := &Customer{ID, UserID, StoreID, FirstName, LastName, Ethnicity, Gender, Birthday, PostalCode, LastVisted, DisChannel, FavItem}
		result = append(result, returnedCus)
	}
	if len(result) == 0 {
		return nil, ErrCustomerNotFound
	}
	return result, nil
}

//Insert inserts new customer into database and returns inserted customer
func (cs *CustomerStorage) Insert(customer *Customer) (*Customer, error) {
	query := "insert into customers (user_id, store_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := cs.sqlsess.Exec(query, customer.UserID, customer.StoreID, customer.FirstName, customer.LastName, customer.Ethnicity, customer.Gender, customer.Birthday, customer.PostalCode, customer.LastVisted, customer.DisChannel, customer.FavItem.ID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	customer.ID = id
	return customer, nil
}

//Update updates existing customer with given id and returns updated customer
func (cs *CustomerStorage) Update(customerId int64, updates *NameUpdates) (*Customer, error) {
	query := "update customers set first_name = ?, last_name = ? where id = ?"
	_, err := cs.sqlsess.Query(query, updates.FirstName, updates.LastName, customerId)
	if err != nil {
		return nil, err
	}
	customer, err := cs.GetById(customerId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

//Delete deletes customer with given customer id
func (cs *CustomerStorage) Delete(customerId int64) error {
	query := "delete customers where id = customerId"
	_, err := cs.sqlsess.Query(query, customerId)
	if err != nil {
		return err
	}
	return nil
}
