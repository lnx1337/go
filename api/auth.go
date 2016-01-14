package api

import (
	"encoding/json"
	"net/http"
	"fmt"
 	"io/ioutil"
)

func Auth(authEndPoint string,next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 		var (
 			accessToken string
 			url string
 			client http.Client
 			err *Err
 		)

 		accessToken = r.Header.Get("Authorization")  
 	 	
 		if accessToken == "" {
 			err = NewError()
 			err.Push(Msg{Error: `ACCESS_TOKEN_NOT_PRESENT`})
 			json.NewEncoder(w).Encode(err)
 			return
 		}

 		url  = fmt.Sprint(authEndPoint, accessToken)
		req, _ := http.NewRequest("GET", url, nil)
		res, errService := client.Do(req)

		if errService != nil {
			err = NewError()
			err.Push(Msg{Error: `AUTH_SERVICE_ERROR`})
			json.NewEncoder(w).Encode(err)
 			return
		}

		defer res.Body.Close()

		// response service
		body, _  := ioutil.ReadAll(res.Body)
		errParse := json.Unmarshal(body,&err)

		// error parse json response 
		if errParse != nil {
			err = NewError()
			err.Push(Msg{Error: `AUTH_SERVICE_RESPONSE_ERROR`})
			json.NewEncoder(w).Encode(err)
 			return
		}

		// access_token_not_valid
		if err.Failed() {
			json.NewEncoder(w).Encode(err)
			return
		}
	    next.ServeHTTP(w, r)
  	})
}