package facebook

import (
    fb "github.com/huandu/facebook"
)

type Facebook struct {
	AccessToken string `json:"access_token"`
	TypeRequest bool   `json:"type_request"`
}

func NewFacebook(access_token string) Facebook {
	self := Facebook{AccessToken:access_token}
	return self
}

//Get list friends user
func (self *Facebook) GetFriends() (interface{},error) {
	res, err := fb.Get("me/friends", fb.Params{
    	"access_token": self.AccessToken,
	})
	return res,err
}


func (self *Facebook) Me() (fb.Result, error) {
	res,err := fb.Get("me",fb.Params{
		"access_token": self.AccessToken,
	})
	return res,err
}