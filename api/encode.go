package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func EncodeRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func EncodeError(w http.ResponseWriter, i error) {
	var buf bytes.Buffer
	w.Header().Set(`Content-Type`, `application/json`)

	if err := json.NewEncoder(&buf).Encode(i); err != nil {
		jserr := NewError()
		jserr.SetErr(`INTERNAL_ERROR`, err.Error(), 500)
		EncodeError(w, jserr)
		return
	}

	// TODO: allow to configure status header [how? I still don't know :/]
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(i)
}
