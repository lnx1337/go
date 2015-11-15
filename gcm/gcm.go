package gcm

import (
	"github.com/alexjlockwood/gcm"
)
// API KEY GOOGLE CLOUD MESSAGE
const APIKEY = "AIzaSyCGeQYbgUkQDLsTbimC8r2J1ztUr17yAcg"

type Gcm struct {
	RegIds []string
	Data   map[string]interface{}
}

func NewPushGcm(regIDs []string, data map[string]interface{}) *Gcm {
	self := &Gcm{}
	self.RegIds = regIDs
	self.Data = data
	return self
}
// Send Push
func (self *Gcm) Send() (interface{}, error) {
	msg := gcm.NewMessage(self.Data, self.RegIds...)
	// Create a Sender to send the message.
	sender := &gcm.Sender{ApiKey: APIKEY}
	// Send the message and receive the response after at most two retries.
	response, err := sender.Send(msg, 2)
	if err != nil {
		return response,err
	}
	return response, nil
}
