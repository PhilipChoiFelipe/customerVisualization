package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/info441-sp21/final-project/server/gateway/models/customers"
	"github.com/info441-sp21/final-project/server/gateway/sessions"
)

/*
handling customers api
Method: GET, POST
Endpoint: version/user/{id}/customers
*/

func (hh *HttpHandler) CustomersHandler(w http.ResponseWriter, r *http.Request) {
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, hh.SigningKey, hh.SessionStore, sessionState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case "GET":
		customers, err := hh.CustomerStorage.GetCustomers(sessionState.AuthUser.ID)
		if err != nil {
			http.Error(w, "Error: failed to find customers by current user", http.StatusBadRequest)
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

	sessionState := &SessionState{}
	_, err = sessions.GetState(r, hh.SigningKey, hh.SessionStore, sessionState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
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
		var updates customers.NameUpdates
		if userId != sessionState.AuthUser.ID {
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
		w.Write([]byte("successfully deleted customer"))
	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}
}

// Helper functions

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
