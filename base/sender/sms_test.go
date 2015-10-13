package sender

import "testing"

const telephone = "+525544963762"

func Test_GetFriends(t *testing.T) {
	self := NewSms(telephone,"send golang")
	self.Send()
}