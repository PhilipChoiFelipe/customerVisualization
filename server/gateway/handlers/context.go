package handlers

import (
	"github.com/info441-sp21/final-project/server/gateway/models/customers"
	"github.com/info441-sp21/final-project/server/gateway/models/users"

	"github.com/info441-sp21/final-project/server/gateway/sessions"

	"github.com/info441-sp21/final-project/server/gateway/models/items"
	"github.com/info441-sp21/final-project/server/gateway/models/stores"
)

type HttpHandler struct {
	SigningKey      string
	SessionStore    sessions.Store
	UserStorage     users.Storage
	CustomerStorage customers.Storage
	ItemStorage     items.Storage
	StoreStorage    stores.Storage
}
