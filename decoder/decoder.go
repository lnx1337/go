// Package decoder implements common decoding patterns
// for use with microservices
// FIXME: this package is missing tests
package decoder

import (
	"encoding/json"
	api "github.com/lnx1337/go/api"
	"golang.org/x/net/context"
	"net/http"
)

type (
	// IDSetter defines types that allow to set their ID
	IDSetter interface {
		SetID(string) interface{}
	}

	// IDSetter defines types that allow to retrieve their ID
	IDGetter interface {
		GetID() interface{}
	}

	// IDSetterGettter
	IDSetterGettter interface {
		IDSetter
		IDGetter
	}
)

// DecodeCreateRequest decodes a POST request into a provided interface
func DecodeCreateRequest(req interface{}, field string) func(context.Context, *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		if len(r.URL.Query().Get(field)) != 0 {
			return req, nil
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			jserr := api.NewError()
			jserr.Push(api.Msg{Error: err.Error()})
			return nil, jserr
		}

		return req, nil
	}
}

// DecodeCreateIDRequest creates a decoder that assings an ID to the returned interface
// for a POST request
func DecodeCreateIDRequest(req IDSetter, field string) func(context.Context, *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		if r.ContentLength <= 0 && len(r.URL.Query().Get(field)) == 0 {
			return req, nil
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			jserr := api.NewError()
			jserr.Push(api.Msg{Error: err.Error()})
			return nil, jserr
		}

		return req.SetID(r.URL.Query().Get(field)), nil
	}
}

// DecodeReadIDRequest creates a decoder that assings an ID to the returned interface
// for a GET request
func DecodeReadIDRequest(req IDSetter, field string) func(context.Context, *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		if r.ContentLength != 0 || len(r.URL.Query().Get(field)) == 0 {
			return req, nil
		}

		return req.SetID(r.URL.Query().Get(field)), nil
	}
}

// DecodeUpdateRequest creates a decoder for a PUT request
func DecodeUpdateRequest(req IDSetter, field string) func(context.Context, *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		if r.ContentLength <= 0 && len(r.URL.Query().Get(field)) == 0 {
			return req, nil
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// return nil, err
			return err, nil
		}

		return req.SetID(r.URL.Query().Get(field)), nil
	}

}
