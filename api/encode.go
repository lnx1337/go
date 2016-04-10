package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	options           string = "OPTIONS"
	allow_origin      string = "Access-Control-Allow-Origin"
	allow_methods     string = "Access-Control-Allow-Methods"
	allow_headers     string = "Access-Control-Allow-Headers"
	allow_credentials string = "Access-Control-Allow-Credentials"
	expose_headers    string = "Access-Control-Expose-Headers"
	credentials       string = "true"
	origin            string = "*"
	methods           string = "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"

	// If you want to expose some other headers add it here
	headers string = "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token"
)


func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	if o := r.Header.Get(origin); o != "" {
		w.Header().Set(allow_origin, o)
	} else {
		w.Header().Set(allow_origin, "*")
	}

		// Set other headers
	w.Header().Set(allow_headers, headers)
	w.Header().Set(allow_methods, methods)
	w.Header().Set(allow_credentials, credentials)
	w.Header().Set(expose_headers, headers)	
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
