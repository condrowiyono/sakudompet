// Package handler manages the data flow from client to appropriate service.
package handler

import (
  "fmt"
  "net/http"
  "strconv"
  "encoding/json"
  "os"
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


var (
  ErrUnknown = CustomError{
    Message:  "Unknow Error",
    Code:     999,
    HTTPCode: http.StatusInternalServerError,
  }
)

// SuccessBody holds data for success response
type SuccessBody struct {
  Data interface{} `json:"data"`
  Meta interface{} `json:"meta"`
}

// ErrorBody holds data for error response
type ErrorBody struct {
  Errors []ErrorInfo `json:"errors"`
  Meta   interface{} `json:"meta"`
}

// MetaInfo holds meta data
type MetaInfo struct {
  HTTPStatus int `json:"http_status"`
}

// ErrorInfo holds error detail
type ErrorInfo struct {
  Message string `json:"message"`
  Code    int    `json:"code"`
  Field   string `json:"field,omitempty"`
}

// CustomError holds data for customized error
type CustomError struct {
  Message  string
  Field    string
  Code     int
  HTTPCode int
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (c CustomError) Error() string {
  return c.Message
}

// BuildSuccess is a function to create SuccessBody
func BuildSuccess(data, meta interface{}) SuccessBody {
  return SuccessBody{
    Data: data,
    Meta: meta,
  }
}

// BuildError is a function to create ErrorBody
func BuildError(errors error) ErrorBody {
  var (
    ce CustomError
    ok bool
  )

  err := errors
  ce, ok = err.(CustomError)
  if !ok {
    ce = CustomError {
      Message:  errors.Error(),
      Code:     999,
      HTTPCode: http.StatusInternalServerError,
    }
  }

  return ErrorBody{
    Errors: []ErrorInfo{
      {
        Message: ce.Message,
        Code:    ce.Code,
        Field:   ce.Field,
      },
    },
    Meta: MetaInfo{
      HTTPStatus: ce.HTTPCode,
    },
  }
}

// Write is a function to write data in json format
func Write(w http.ResponseWriter, result interface{}, status int) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  json.NewEncoder(w).Encode(result)
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

func (h *Handler) BasicAuth (fn func(http.ResponseWriter, *http.Request, httprouter.Params) error) (func(http.ResponseWriter, *http.Request, httprouter.Params) error) {
  return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
    user, password, hasAuth := r.BasicAuth()

    if hasAuth && user == os.Getenv("BASIC_USERNAME") && password == os.Getenv("BASIC_PASSWORD") {
       return fn(w, r, params)
    } else {
      err := CustomError{ Message: http.StatusText(http.StatusUnauthorized), Code: 8502, HTTPCode: http.StatusUnauthorized }
      writeError(w, err)
      return err
    }
  }
}

func writeError(w http.ResponseWriter, err error) {
  res := BuildError(err)
  Write(w, res, 200)
}

func writeSuccess(w http.ResponseWriter, data interface{}, meta interface{}) {
  res := BuildSuccess(data, meta)
  Write(w, res, 200)
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
