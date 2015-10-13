package facebook

import "testing"

const access_token = "CAACEdEose0cBACrN1VN7Ou2C3AlIGQFPezdwduBGLDPXp2j7YoFyjfBwNaZBVFz2chvrHbdxrGP5bqQQe9y9WwpmgDAIR2BYOsyLGjY6KzInW5RlZAaE7X5kDqxoysLmIejR6JbWadZCVGGhZCDhnCLWG8Q9sYQX9ZAmIr2Mv0fskFmCNcYM9APcCPDmRY6ZChKzseEBVugtYt0t89xbunmHearhlFsu4ZD"

func Test_GetFriends(t *testing.T) {
	self := NewFacebook(access_token)
	_, err :=self.GetFriends()

	if err != nil {
		t.Error(err)
	}

}