package facebook

import "testing"

const access_token = ""

func Test_GetFriends(t *testing.T) {
	self := NewFacebook(access_token)
	_, err :=self.GetFriends()

	if err != nil {
		t.Error(err)
	}

}
