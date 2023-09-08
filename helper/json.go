package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromReqBody(r *http.Request, result interface {})  {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResBody(w http.ResponseWriter, res interface {})  {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	PanicIfError(err)
}