package handlers

import (
	"time"

	"github.com/info441-sp21/final-project/server/gateway/models/users"
)

//SessionsState contains authenticated user and session's start time
type SessionState struct {
	SessionTime time.Time  `json:"sessionTime"`
	AuthUser    users.User `json:"authUser"`
}
