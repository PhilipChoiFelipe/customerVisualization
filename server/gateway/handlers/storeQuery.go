package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/info441-sp21/final-project/server/gateway/models/stores"
	"github.com/info441-sp21/final-project/server/gateway/sessions"
)

/*
handling user's stores api
Method:
	GET: Get all the stores
	POST: Save new stores
Endpoint: version/user/{user_id}/stores
*/

func (hh *HttpHandler) StoreHandler(w http.ResponseWriter, r *http.Request) {
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, hh.SigningKey, hh.SessionStore, sessionState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case "GET":
		stores, err := hh.StoreStorage.GetStores(sessionState.AuthUser.ID)
		if err != nil {
			http.Error(w, "Error: failed to find stores by current user", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stores)
	case "POST":
		var store stores.Store
		rHeader := r.Header.Get("Content-Type")
		if !strings.HasPrefix(rHeader, "application/json") {
			http.Error(w, "ERROR: request body must have json format", http.StatusUnsupportedMediaType)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&store)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insertedStore, err := hh.StoreStorage.Insert(&store)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		buffer, err := json.Marshal(insertedStore)
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
handling specific store api
Method:
	GET: get specific store
	PATCH: update specific store
	DELETE: Delete specific store by current user
Endpoint: version/user/{user_id}/stores/{store_id}
*/

func (hh *HttpHandler) SpecificStoreHandler(w http.ResponseWriter, r *http.Request) {
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, hh.SigningKey, hh.SessionStore, sessionState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	storeId, err := convertsId(r, "store_id")
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
		store, err := hh.StoreStorage.GetById(storeId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(store)
	case "PATCH":
		var updates stores.StoreUpdate
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
		updatedStore, err := hh.StoreStorage.Update(storeId, &updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedStore)
	case "DELETE":
		err := hh.StoreStorage.Delete(storeId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("successfully deleted store"))
	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}

}
