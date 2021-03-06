package api

import (
	"bytes"
	"encoding/json"
	"context"
	"io/ioutil"
	"net/http"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
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

func EncodeError(ctx context.Context, i error, w http.ResponseWriter) {
	var buf bytes.Buffer
	w.Header().Set(`Content-Type`, `application/json`)

	if err := json.NewEncoder(&buf).Encode(i); err != nil {
		jserr := NewError()
		jserr.SetErr(`INTERNAL_ERROR`, err.Error(), 500)
		EncodeError(ctx, jserr, w)
		return
	}

	// TODO: allow to configure status header [how? I still don't know :/]
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(i)
}
