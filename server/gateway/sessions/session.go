package sessions

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"
const paramAuthorization = "auth"
const schemeBearer = "Bearer "

//ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + headerAuthorization + " header")

//ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

//getAuthFromReq extracts session id from request header and returns it.
func getAuthFromReq(r *http.Request) (SessionID, error) {
	authVal := ""
	authVal = r.Header.Get(headerAuthorization)
	if len(authVal) < 1 {
		authVal = r.URL.Query().Get(paramAuthorization)
	}
	if !strings.HasPrefix(authVal, "Bearer") {
		return InvalidSessionID, fmt.Errorf("found no Bearer Prefix")
	}
	authVal = strings.Replace(authVal, "Bearer ", "", -1)
	if len(authVal) < 1 {
		return InvalidSessionID, fmt.Errorf("no authorize value in request")
	}
	return SessionID(authVal), nil
}

//BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
//Authorization header to the response with the SessionID, and returns the new SessionID
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {
	if len(signingKey) < 1 {
		return InvalidSessionID, fmt.Errorf("signing key is invalid")
	}
	sid, err := NewSessionID(signingKey)
	if err != nil {
		return InvalidSessionID, nil
	}
	store.Save(sid, sessionState)
	w.Header().Add(headerAuthorization, schemeBearer+string(sid))
	return sid, nil
}

//GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {
	authVal, err := getAuthFromReq(r)
	if err != nil {
		return InvalidSessionID, err
	}
	sid, err := ValidateID(string(authVal), signingKey)
	if err != nil || len(sid) < 1 {
		return InvalidSessionID, err
	}
	return sid, err
}

//GetState extracts the SessionID from the request,
//gets the associated state from the provided store into
//the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {
	sid, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	err = store.Get(sid, sessionState)
	if err != nil {
		return sid, err
	}

	return sid, nil
}

//EndSession extracts the SessionID from the request,
//and deletes the associated data in the provided store, returning
//the extracted SessionID.
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	sid, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	err = store.Delete(sid)
	if err != nil {
		return InvalidSessionID, err
	}
	return sid, nil
}
