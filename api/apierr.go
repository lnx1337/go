// Provee structs JsonErr y JsonMsg para comunicar errores y mensajes de un API
// a un consumer JSON.
package api

import (
	"fmt"
)

type Err struct {
	HasErrors  bool `json:"hasErrors"` 
	Errors []Msg `json:"errors"`
}

type Msg struct {
	Field string `json:"field,omitempty"`
	Error string `json:"error,omitempty"`
	Code  int    `json:"code,omitempty"`
}

func (self *Err) Push(message Msg) error {
	self.HasErrors = true
	self.Errors = append(self.Errors, message)
	return nil
}

func (self *Err) SetErr(_field string, _message string, _code int) error {
	msg := Msg{
		Field: _field,
		Error: _message,
	}
	self.Errors = append(self.Errors, msg)
	return nil
}

func (self *Err) Len() int {
	return len(self.Errors)
}

func (self *Err) Failed() bool {
	return len(self.Errors) > 0
}

// Error implements error interface, joining each Msg field, error and code
func (e *Err) Error() string {
	var err string
	for _, m := range e.Errors {
		err = fmt.Sprintf("%s %s: %s code: %d\n", err, m.Field, m.Error, m.Code)
	}

	return err
}

func NewError() *Err {
	self := &Err{
		Errors: []Msg{},
	}
	return self
}
