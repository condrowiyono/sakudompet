package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func Gateway(action string, fn func(http.ResponseWriter, *http.Request, httprouter.Params) error) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    	if err := fn(w, r, params); err !=nil {
    		//do something wkwkwk
    	} 
    }
}