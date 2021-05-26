package handlers

import (
	"final-project/server/gateway/models/customers"
	"final-project/server/gateway/models/users"
	"final-project/server/gateway/sessions"
)

type HttpHandler struct {
	SigningKey      string
	SessionStore    sessions.Store
	UserStorage     users.Storage
	CustomerStorage customers.Storage
}
