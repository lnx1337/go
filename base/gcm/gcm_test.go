package gcm

import "testing"

func Test_GetFriends(t *testing.T) {
	regIDs := []string{"API_KEY"}
	data := map[string]interface{}{"data": "Send GO test"}
	self := NewPushGcm(regIDs, data)
	_, err := self.Send()
	if err != nil {
		t.Error(err)
	}
}
