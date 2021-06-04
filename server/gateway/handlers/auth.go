package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/info441-sp21/final-project/server/gateway/models/users"
	"github.com/info441-sp21/final-project/server/gateway/sessions"
	"golang.org/x/crypto/bcrypt"
)

//UserHandler validates inserts new user and creates new session
func (hh *HttpHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var nu users.NewUser
		rHeader := r.Header.Get("Content-Type")
		if !strings.HasPrefix(rHeader, "application/json") {
			http.Error(w, "ERROR: request body must have json format", http.StatusUnsupportedMediaType)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&nu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := nu.ToUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insertedUser, err := hh.UserStorage.Insert(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sessionState := SessionState{SessionTime: time.Now(), AuthUser: *insertedUser}
		_, err = sessions.BeginSession(hh.SigningKey, hh.SessionStore, sessionState, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		buffer, err := json.Marshal(insertedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if strings.Contains(string(buffer), "PassHash") || strings.Contains(string(buffer), "Email") {
			http.Error(w, "ERROR: returning value error", http.StatusConflict)
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

//SpecificUserHandler validates request's session and respond processed user
func (hh *HttpHandler) SpecificUserHandler(w http.ResponseWriter, r *http.Request) {
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, hh.SigningKey, hh.SessionStore, sessionState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	baseId := path.Base(r.URL.Path)
	if baseId == "me" {
		baseId = strconv.Itoa(int(sessionState.AuthUser.ID))
	}
	userId, err := strconv.ParseInt(baseId, 10, 64)
	if err != nil {
		http.Error(w, "ERROR: invalid user ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		user, err := hh.UserStorage.GetById(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	case "PATCH":
		var updates users.Updates
		if userId != sessionState.AuthUser.ID {
			http.Error(w, "ERROR: unauthorized user", http.StatusForbidden)
			return
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "ERROR: request body must have json format", http.StatusUnsupportedMediaType)
			return
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := hh.UserStorage.Update(userId, &updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	case "DELETE":
		err := hh.CustomerStorage.DeleteAllbyUserId(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = hh.ItemStorage.DeleteAllbyUserId(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = hh.UserStorage.Delete(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
		}
		w.Write([]byte("successfully deleted user"))

	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}
}

//SessionHandler handles user's login and creates new session. Respond with requested user
func (hh *HttpHandler) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rHeader := r.Header.Get("Content-Type")
		if !strings.HasPrefix(rHeader, "application/json") {
			http.Error(w, "ERROR: request body must have json format", http.StatusUnsupportedMediaType)
			return
		}
		var uc users.Credentials
		err := json.NewDecoder(r.Body).Decode(&uc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, emailErr := hh.UserStorage.GetByEmail(uc.Email)
		if emailErr != nil {
			dummyHash := []byte("dummyHash123!@#")
			bcrypt.CompareHashAndPassword(dummyHash, []byte("dummydummywrong"))
			http.Error(w, "ERROR: invalid credentials", http.StatusUnauthorized)
			return
		}
		passwordErr := user.Authenticate(uc.Password)
		if passwordErr != nil {
			http.Error(w, "ERROR: invalid credentials", http.StatusUnauthorized)
			return
		}
		sessionState := &SessionState{SessionTime: time.Now(), AuthUser: *user}
		_, err = sessions.BeginSession(hh.SigningKey, hh.SessionStore, sessionState, w)
		if err != nil {
			http.Error(w, "ERROR: failed to begin session", http.StatusBadGateway)
			return
		}
		clientIP := r.RemoteAddr
		userID := user.ID
		err = hh.UserStorage.InsertSignIn(userID, clientIP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		err = hh.UserStorage.LogSignIns(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}
}

//SpecificSessionHandler validates current user's session and ends it.
func (hh *HttpHandler) SpecificSessionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		params := mux.Vars(r)
		selectedSession := params["session_id"]
		if selectedSession != "mine" {
			http.Error(w, "ERROR: need appropriate sesssion call", http.StatusForbidden)
			return
		}
		_, err := sessions.EndSession(r, hh.SigningKey, hh.SessionStore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Signed Out"))
	default:
		http.Error(w, "ERROR: wrong request method", http.StatusMethodNotAllowed)
		return
	}
}
