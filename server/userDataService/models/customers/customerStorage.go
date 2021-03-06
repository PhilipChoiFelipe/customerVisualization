package customers

import (
	"database/sql"
	"fmt"
	"log"
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
	ID          int64
	UserID      int64
	FirstName   string
	LastName    string
	Ethnicity   string
	Gender      string
	Birthday    string
	PostalCode  int64
	LastVisited string
	DisChannel  string
	ItemId      int64
)

//GetById finds id of customers in DB and returns the customer
func (cs *CustomerStorage) GetById(customerId int64) (*Customer, error) {

	query := "select id, user_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item from customers where id = ?"
	err := cs.sqlsess.QueryRow(query, customerId).Scan(&ID, &UserID, &FirstName, &LastName, &Ethnicity, &Gender, &Birthday, &PostalCode, &LastVisited, &DisChannel, &ItemId)
	if err != nil {
		return nil, err
	}
	return &Customer{ID, UserID, FirstName, LastName, Ethnicity, Gender, Birthday, PostalCode, LastVisited, DisChannel, ItemId}, nil
}

//GetByItemId finds itemId of customers in DB and returns the customers with certain favorite ids
func (cs *CustomerStorage) GetByItemId(itemId int64) ([]*Customer, error) {
	var result []*Customer
	query := "select id, user_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item from customers where fav_item = ?"
	rows, err := cs.sqlsess.Query(query, itemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&ID, &UserID, &FirstName, &LastName, &Ethnicity, &Gender, &Birthday, &PostalCode, &LastVisited, &DisChannel, &ItemId)
		returnedCus := &Customer{ID, UserID, FirstName, LastName, Ethnicity, Gender, Birthday, PostalCode, LastVisited, DisChannel, ItemId}
		result = append(result, returnedCus)
	}
	if len(result) == 0 {
		return nil, ErrCustomerNotFound
	}
	return result, nil
}

//GetCustomers returns the all customers with given user Id
func (cs *CustomerStorage) GetCustomers(user_id int64, queryCase string, col_name string, reverse string, beforeDate string, afterDate string) ([]*Customer, error) {
	var result []*Customer

	query := "select id, user_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item from customers where user_id = ?"
	var rows *sql.Rows
	var err error

	switch queryCase {
	case "default":
		rows, err = cs.sqlsess.Query(query, user_id)
		if err != nil {
			return nil, err
		}
	case "sort":
		if reverse == "true" {
			query = fmt.Sprintf("select * from customers where user_id = ? order by %s DESC", col_name)

		} else {
			query = fmt.Sprintf("select * from customers where user_id = ? order by %s ASC", col_name)
		}
		rows, err = cs.sqlsess.Query(query, user_id)
		if err != nil {
			return nil, err
		}
	}
	log.Println("query:", query)

	if err != nil {
		log.Println(fmt.Sprintf("err: %v", err.Error()))
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&ID, &UserID, &FirstName, &LastName, &Ethnicity, &Gender, &Birthday, &PostalCode, &LastVisited, &DisChannel, &ItemId)
		returnedCus := &Customer{ID, UserID, FirstName, LastName, Ethnicity, Gender, Birthday, PostalCode, LastVisited, DisChannel, ItemId}
		result = append(result, returnedCus)
	}
	if len(result) == 0 {
		return nil, ErrCustomerNotFound
	}
	return result, nil
}

//Insert inserts new customer into database and returns inserted customer
func (cs *CustomerStorage) Insert(customer *Customer) (*Customer, error) {
	query := "insert into customers (user_id, first_name, last_name, ethnicity, gender, birthday, postal_code, last_visited, dis_channel, fav_item) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := cs.sqlsess.Exec(query, customer.UserID, customer.FirstName, customer.LastName, customer.Ethnicity, customer.Gender, customer.Birthday, customer.PostalCode, customer.LastVisited, customer.DisChannel, customer.FavItem)
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
func (cs *CustomerStorage) Update(customerId int64, updates *Updates) (*Customer, error) {
	query := "update customers set first_name = ?, last_name = ?, ethnicity = ?, gender = ?, birthday = ?, postal_code = ?, last_visited = ?, dis_channel = ?, fav_item = ? where id = ?"
	_, err := cs.sqlsess.Exec(query, updates.FirstName, updates.LastName, updates.Ethnicity, updates.Gender, updates.Birthday, updates.PostalCode, updates.LastVisited, updates.DisChannel, updates.FavItem, customerId)
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
	query := "delete from customers where id = ?"
	_, err := cs.sqlsess.Exec(query, customerId)
	if err != nil {
		return err
	}
	return nil
}

//Delete deletes customer with given user id
func (cs *CustomerStorage) DeleteAllbyUserId(userId int64) error {
	delq := "delete from customers where user_id = ?"
	_, err := cs.sqlsess.Exec(delq, userId)
	if err != nil {
		return fmt.Errorf("%v: %v", ErrCustomerNotDeleted, err)
	}

	return nil
}
