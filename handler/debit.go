package handler

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "errors"
  "strconv"
  "encoding/json"

  "github.com/julienschmidt/httprouter"
  
  "github.com/condrowiyono/sakudompet/saku"
  "github.com/condrowiyono/sakudompet/pass"
)

func (h *Handler) FindAllDebits(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx := r.Context()
	select {
		case <-ctx.Done():
			fmt.Print(errors.New("Timeout"))
			return errors.New("Timeout")
		default:
	}

	result, _ := h.sakudompet.GetDebits(ctx)
	meta := Meta{
		HTTPStatus: 200,
	}
	
	for i := 0; i < len(result); i++ {
		pass, err := h.sakudompet.FindPass(ctx, result[i].PassId)
		if err == nil {
			result[i].Pass = &pass
		}
    }

	writeSuccess(w, result, meta)
	return nil
}

func (h *Handler) FindDebit(w http.ResponseWriter, r *http.Request, params httprouter.Params) error{
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

	result, err := h.sakudompet.FindDebit(ctx,uint(id))
	if err != nil {
		writeError(w, err)
		return err
	} else {
		var pass pass.Pass
		pass, err = h.sakudompet.FindPass(ctx, result.PassId)

		if err == nil {
			result.Pass = &pass
		}
		
		meta := Meta{
			HTTPStatus: 200,
		}
		writeSuccess(w, result, meta)
		return nil
	}
}

func (h *Handler) CreateDebit(w http.ResponseWriter, r *http.Request, params httprouter.Params) error{
	ctx := r.Context()
	select {
		case <-ctx.Done():
			fmt.Print(errors.New("Timeout"))
		default:
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
		return err
	}
	
	var debit saku.Debit

	queryValues := r.URL.Query()
  	isPass, err := strconv.Atoi(queryValues.Get("pass"))

	if err != nil {
		isPass = 0
	}
	
	err = json.Unmarshal(body, &debit)
	if err != nil {
		writeError(w, err)
		return err
	}

	result, err := h.sakudompet.CreateDebit(ctx, debit)
  	if err != nil {
		writeError(w, err)
		return err
	} else {
		if (isPass==1) {
			primary := saku.DebitPrimaryField {
				Key: 	"Name",
				Value:	debit.Name,
			}
			secondary := saku.DebitSecondaryField {
				Key: 	"Name",
				Value:	"",
			}
			aux := saku.DebitAuxiliaryField {
				Key: 	"Name",
				Value:	"",
			}

			primaryField, _ := json.Marshal(primary)
			secondaryField, _ := json.Marshal(secondary)
			auxiliaryField, _ := json.Marshal(aux) 

			pass := pass.Pass{
				Logo: 					debit.IssuedBy,
				LogoText:				debit.IssuedBy,
				HeaderFields:			debit.IssuedBy,
				PrimaryField:			string(primaryField),
				SecondaryField:			string(secondaryField),
				AuxiliaryField:			string(auxiliaryField),
				Background:				"",
				Thumbnail:				"",
				BarcodeType:			"",
				BarcodeMessage:			"",
			}
			resultPass, err := h.sakudompet.CreatePass(ctx, pass);
			if err != nil {
				writeError(w, err)
				return err
			}
			result.PassId = resultPass.Id
			result.Pass = &resultPass
			
			_,err = h.sakudompet.PutDebit(ctx, result.Id, result)
			
			if err != nil {
				writeError(w, err)
				return err
			}
		}

		meta := Meta{
			HTTPStatus: 200,
		}
		writeSuccess(w, result, meta)
		return nil
	}
}

func (h *Handler) PutDebit(w http.ResponseWriter, r *http.Request, params httprouter.Params) error{
	ctx := r.Context()
	select {
		case <-ctx.Done():
			fmt.Print(errors.New("Timeout"))
		default:
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
		return err
	}
	
	var debit saku.Debit

	err = json.Unmarshal(body, &debit)
	if err != nil {
		writeError(w, err)
		return err
	}

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writeError(w, err)
		return err
	}

	result, err := h.sakudompet.PutDebit(ctx, uint(id), debit)
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

func (h *Handler) DeleteDebit(w http.ResponseWriter, r *http.Request, params httprouter.Params) error{
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

	_, err = h.sakudompet.DeleteDebit(ctx,uint(id))
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
