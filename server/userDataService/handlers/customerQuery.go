package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/info441-sp21/final-project/server/gateway/models/users"
	"github.com/info441-sp21/final-project/server/userDataService/models/customers"
)

/*
handling customers api
Method: GET, POST
Endpoint: version/user/{id}/customers
*/

func (hh *HttpHandler) CustomersHandler(w http.ResponseWriter, r *http.Request) {
	//User
	//type User struct {
	// 	ID        int64  `json:"id"`
	// 	Email     string `json:"-"` //never JSON encoded/decoded
	// 	UserName  string `json:"userName"`
	// 	PassHash  []byte `json:"-"` //never JSON encoded/decoded
	// 	FirstName string `json:"firstName"`
	// 	LastName  string `json:"lastName"`
	// 	StoreName string `json:"storeName"`
	// }
	var authUser users.User
	authHeader := r.Header.Get("X-User")
	if len(authHeader) == 0 {
		http.Error(w, "current user it not authorized", http.StatusUnauthorized)
		return
	}
	err := json.Unmarshal([]byte(authHeader), &authUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		queryCase := "default"
		// params := mux.Vars(r)
		query := r.URL.Query()

		col_name := query.Get("sort")
		reverse := query.Get("reverse")
		if reverse != "" {
			queryCase = "sort"
		}
		beforeDate := query.Get("before")
		if beforeDate != "" {
			queryCase = "sortBefore"
		}
		afterDate := query.Get("after")
		if afterDate != "" {
			queryCase = "sortAfter"
		}

		customers, err := hh.CustomerStorage.GetCustomers(authUser.ID, queryCase, col_name, reverse, beforeDate, afterDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	case "POST":
		var customer customers.Customer
		rHeader := r.Header.Get("Content-Type")
		if !strings.HasPrefix(rHeader, "application/json") {
			http.Error(w, "ERROR: request body must have json format", http.StatusUnsupportedMediaType)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insertedCustomer, err := hh.CustomerStorage.Insert(&customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		buffer, err := json.Marshal(insertedCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(buffer)
	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}
}

/*
handling specific customer api
Method:
	GET: get specific customer
	PATCH: update specific customer info
	DELETE: elete specific customer info
Endpoint: version/user/{user_id}/customers/{customer_id}
*/

func (hh *HttpHandler) SpecificCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var authUser users.User
	authHeader := r.Header.Get("X-User")
	if len(authHeader) == 0 {
		http.Error(w, "current user it not authorized", http.StatusUnauthorized)
		return
	}
	err := json.Unmarshal([]byte(authHeader), &authUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customerId, err := convertsId(r, "customer_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := convertsId(r, "user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		customer, err := hh.CustomerStorage.GetById(customerId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	case "PATCH":
		var updates customers.Updates
		if userId != authUser.ID {
			http.Error(w, "ERROR: unauthorized user", http.StatusForbidden)
			return
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "ERROR: request body must have json format", http.StatusUnsupportedMediaType)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedCustomer, err := hh.CustomerStorage.Update(customerId, &updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedCustomer)
	case "DELETE":
		err := hh.CustomerStorage.Delete(customerId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("successfully deleted customer"))
	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}
}

// Converts mux param into id
func convertsId(r *http.Request, param string) (int64, error) {
	params := mux.Vars(r)
	idStr := params[param]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
