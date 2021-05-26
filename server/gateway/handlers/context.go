package handlers

import (
	"github.com/info441-sp21/final-project/server/gateway/models/customers"
	"github.com/info441-sp21/final-project/server/gateway/models/users"
	"github.com/info441-sp21/final-project/server/gateway/sessions"
)

type HttpHandler struct {
	SigningKey      string
	SessionStore    sessions.Store
	UserStorage     users.Storage
	CustomerStorage customers.Storage
}
