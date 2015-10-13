package gcm

import "testing"

func Test_GetFriends(t *testing.T) {
	regIDs := []string{"APA91bECa-xQqTrNSJNSxgPjgzjmOnc0128EQksWfE7sI1dh2W97ffWKV_Z-hJkGlYpprKXBpq36WZjJUho97ecjSz9Y9N4sh5eHX0i9Y_QWvhY6GWqXRl3NxNUCoit5LQQOBRBxSFV34sLNPvvIa-Y4xkEwhrk9Qw"}
	data := map[string]interface{}{"data": "Send GO test"}
	self := NewPushGcm(regIDs, data)
	_, err := self.Send()
	if err != nil {
		t.Error(err)
	}
}
