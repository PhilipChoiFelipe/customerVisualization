package handlers

import "net/http"

/* TODO: implement a CORS middleware handler, as described
in https://drstearns.github.io/tutorials/cors/ that responds
with the following headers to all requests:

  Access-Control-Allow-Origin: *
  Access-Control-Allow-Methods: GET, PUT, POST, PATCH, DELETE
  Access-Control-Allow-Headers: Content-Type, Authorization
  Access-Control-Expose-Headers: Authorization
  Access-Control-Max-Age: 600
*/

const  ACAllowOrigin = "Access-Control-Allow-Origin"
const  ACAllowMethods = "Access-Control-Allow-Methods"
const  ACAllowHeaders = "Access-Control-Allow-Headers"
const  ACExposeHeaders = "Access-Control-Expose-Headers"
const  ACMaxAge = "Access-Control-Max-Age"

const  ACAllowOriginValue = "*"
const  ACAllowMethodsValue = "GET, PUT, POST, PATCH, DELETE"
const  ACAllowHeadersValue = "Content-Type, Authorization"
const  ACExposeHeadersValue = "Authorization"
const  ACMaxAgeValue = "600"

type CORSHandler struct {
  handler http.Handler
}

func NewCORSHandler(handlerToWrap http.Handler) *CORSHandler {
    return &CORSHandler{handlerToWrap}
}

func (ch *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set(ACAllowOrigin, ACAllowOriginValue)
    w.Header().Set(ACAllowMethods, ACAllowMethodsValue)
    w.Header().Set(ACAllowHeaders, ACAllowHeadersValue)
    w.Header().Set(ACExposeHeaders, ACExposeHeadersValue)
    w.Header().Set(ACMaxAge, ACMaxAgeValue)

    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusOK)
      return
    }

    ch.handler.ServeHTTP(w, r)
}

