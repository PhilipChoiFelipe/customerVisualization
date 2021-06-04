// Director is the director used for routing to microservices
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/info441-sp21/final-project/server/gateway/handlers"
	"github.com/info441-sp21/final-project/server/gateway/sessions"
)

type Director func(r *http.Request)

// CustomDirector forwards to the microservice and passes it the current user.
func CustomDirector(targets []*url.URL, ctx *handlers.HttpHandler) Director {
	var counter int32
	counter = 0
	mutex := sync.Mutex{}
	log.Println("Director being created!!")
	return func(r *http.Request) {
		log.Println("Director called!")
		mutex.Lock()
		defer mutex.Unlock()
		targ := targets[counter%int32(len(targets))]
		atomic.AddInt32(&counter, 1)
		r.Header.Add("X-Forwarded-Host", r.Host)
		r.Header.Del("X-User")
		sessionState := &handlers.SessionState{}
		_, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, sessionState)
		log.Printf("derector line 34) SESSION STATE:%v", sessionState)
		// If there is an error, forward it to the API to deal with it.
		if err != nil {
			r.Header.Add("X-User", "")
		} else {
			user := sessionState.AuthUser
			userJSON, err := json.Marshal(user)
			if err != nil {
				r.Header.Add("X-User", "")
			} else {
				r.Header.Add("X-User", string(userJSON))
			}
		}
		log.Printf("derector line 49) TARGET for request: %v", targ)
		log.Printf("derector line 49) TARGET for request: %v", targ.String())
		r.Host = targ.String()
		r.URL.Host = targ.String()
		r.URL.Scheme = "http"
	}
}

func GetURLs(addrString string) []*url.URL {
	addrsSplit := strings.Split(addrString, ",")
	URLs := make([]*url.URL, len(addrsSplit))
	for i, c := range addrsSplit {
		URL, err := url.Parse(c)
		if err != nil {
			log.Fatal(fmt.Printf("Failure to parse url %v", err))
		}
		URLs[i] = URL
	}
	return URLs
}
