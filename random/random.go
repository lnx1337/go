package random

import "crypto/rand"

type Random struct {
	Number int
}

func NewRandom(number int) string {
	self := Random{Number:number}
	return self.Generate()
}

func (self *Random) Generate() string {
    const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, self.Number)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b%byte(len(alphanum))]
    }
    key := string(bytes)
    return key
}