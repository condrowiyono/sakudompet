// Package handler manages the data flow from client to appropriate service.
package handler

import (
  "fmt"
  "net/http"
  "strconv"
  "github.com/julienschmidt/httprouter"

  sd "github.com/condrowiyono/sakudompet"
)

// Handler controls request flow from client to service
type Handler struct {
  sakudompet *sd.SakuDompet
}


type Meta struct {
  Offset     int `json:"offset"`
  Limit      int `json:"limit"`
  Total      int `json:"total"`
  HTTPStatus int `json:"http_status"`
}

// NewHandler returns a pointer of Handler instance
func NewHandler(sakudompet *sd.SakuDompet) *Handler {
  return &Handler{
    sakudompet: sakudompet,
  }
}


// Healthz is used to control the flow of GET /healthz endpoint
func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintln(w, "ok")
}

func getOffsetLimit(offset string, limit string) (int, int) {
  var offint, limint int
  var err error
  if offset == "" {
    offint = 0
  } else {
    offint, err = strconv.Atoi(offset)
    if err != nil {
      offint = 0
    }
  }

  if limit == "" {
    limint = 10
  } else {
    limint, err = strconv.Atoi(limit)
    if err != nil || limint > 100 {
      limint = 100
    }
  }

  return offint, limint
}
