package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Auth(authEndPoint string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			accessToken string
			url         string
			client      http.Client
			err         *Err
			resApi      *Response
		)

		accessToken = r.Header.Get("Authorization")

		if accessToken == "" {
			err = NewError()
			err.Push(Msg{Error: `ACCESS_TOKEN_NOT_PRESENT`})
			resApi = &Response{Errors:err.Errors,Error:true}
			json.NewEncoder(w).Encode(resApi)
			return
		}

		url = fmt.Sprint(authEndPoint, accessToken)
		req, _ := http.NewRequest("GET", url, nil)
		res, errService := client.Do(req)

		if errService != nil {
			err = NewError()
			err.Push(Msg{Error: `AUTH_SERVICE_ERROR`})
			resApi = &Response{Errors:err.Errors,Error:true}
			json.NewEncoder(w).Encode(resApi)
			return
		}

		defer res.Body.Close()

		// response service
		body, _ := ioutil.ReadAll(res.Body)
		errParse := json.Unmarshal(body, &err)

		// error parse json response
		if errParse != nil {
			err = NewError()
			err.Push(Msg{Error: `AUTH_SERVICE_RESPONSE_ERROR`})
			resApi = &Response{Errors:err.Errors, Error:true}
			json.NewEncoder(w).Encode(resApi)
			return
		}

		// access_token_not_valid
		if err.Failed() {
			resApi = &Response{Errors:err.Errors,Error:true}
			json.NewEncoder(w).Encode(resApi)
			return
		}
		next.ServeHTTP(w, r)
	})
}
