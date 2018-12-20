package handler

import (
  "fmt"
  "net/http"
  "errors"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
)

func (h *Handler) FindAllDebits(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	select {
		case <-ctx.Done():
			fmt.Print(errors.New("Timeout"))
		default:
	}
	w.Header().Set("Content-Type", "application/json")

	result, _ := h.sakudompet.GetDebits(ctx)
  	
  	json.NewEncoder(w).Encode(result)
}