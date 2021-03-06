package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/info441-sp21/final-project/server/gateway/models/users"
	"github.com/info441-sp21/final-project/server/userDataService/models/items"
)

/*
handling user's spec items api
Method:
	GET: Get all the items
	POST: Save new item
Endpoint: "/v1/user/{user_id}/items"
*/

func (hh *HttpHandler) ItemsHandler(w http.ResponseWriter, r *http.Request) {

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
	userId, err := convertsId(r, "user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		queryCase := "default"
		query := r.URL.Query()

		col_name := query.Get("sort")
		reverse := query.Get("reverse")
		if reverse != "" {
			queryCase = "sort"
		}

		items, err := hh.ItemStorage.GetItems(userId, queryCase, col_name, reverse)
		if err != nil {
			http.Error(w, "Error: failed to find customers by current user", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
	case "POST":
		var item items.Item
		rHeader := r.Header.Get("Content-Type")
		if !strings.HasPrefix(rHeader, "application/json") {
			http.Error(w, "ERROR: request body must have json format", http.StatusUnsupportedMediaType)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insertedItem, err := hh.ItemStorage.Insert(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		buffer, err := json.Marshal(insertedItem)
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
handling specific item api
Method:
	GET: get specific item
	PATCH: update specific item
	DELETE: Delete specific item by current user
Endpoint: "/v1/user/{user_id}/items/{item_id}"
*/

func (hh *HttpHandler) SpecificItemHandler(w http.ResponseWriter, r *http.Request) {
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

	userId, err := convertsId(r, "user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	itemId, err := convertsId(r, "item_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		item, err := hh.ItemStorage.GetById(itemId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	case "PATCH":
		var updates items.ItemUpdate
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
		updatedItem, err := hh.ItemStorage.Update(itemId, &updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedItem)
	case "DELETE":
		err := hh.ItemStorage.Delete(itemId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("successfully deleted item"))
	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}

}
