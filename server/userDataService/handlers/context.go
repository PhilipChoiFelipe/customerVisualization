package handlers

import (
	"github.com/info441-sp21/final-project/server/userDataService/models/customers"
	"github.com/info441-sp21/final-project/server/userDataService/models/items"
)

type HttpHandler struct {
	CustomerStorage customers.Storage
	ItemStorage     items.Storage
}
