package random

import "testing"
import "fmt"

func Test_GenerateCodeConfirm(t *testing.T) {
	code := NewRandom(4)
	fmt.Println(code)
}