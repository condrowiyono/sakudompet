package handler

import (
  "fmt"
  "net/http"
  "errors"
  "strconv"

  "github.com/julienschmidt/httprouter"
)

func (h *Handler) GetPasses(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx := r.Context()
	select {
		case <-ctx.Done():
			fmt.Print(errors.New("Timeout"))
			return errors.New("Timeout")
		default:
	}

	result, _ := h.sakudompet.GetPasses(ctx)
	meta := Meta{
		HTTPStatus: 200,
	}

	writeSuccess(w, result, meta)
	return nil
}

func (h *Handler) FindPass(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx := r.Context()
	select {
		case <-ctx.Done():
			fmt.Print(errors.New("Timeout"))
			return errors.New("Timeout")
		default:
	}
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writeError(w, err)
		return err
	}

	result, err := h.sakudompet.FindPass(ctx, uint(id))
	
	if err != nil {
		writeError(w, err)
		return err
	}

	meta := Meta{
		HTTPStatus: 200,
	}

	writeSuccess(w, result, meta)
	return nil
}

func (h *Handler) DeletePass(w http.ResponseWriter, r *http.Request, params httprouter.Params) error{
	ctx := r.Context()
	select {
		case <-ctx.Done():
			fmt.Print(errors.New("Timeout"))
		default:
	}

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writeError(w, err)
		return err
	}

	_, err = h.sakudompet.DeletePass(ctx,uint(id))
	if err != nil {
		writeError(w, err)
		return err
	}

	data := map[string]interface{} {
		"message" : "Berhasil di hapus",
	}

	meta := Meta{
		HTTPStatus: 200,
	}
	
	writeSuccess(w, data, meta)
	return nil
}