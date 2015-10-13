// Respuesta standard de API a consumidor.
package api

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
	Error  bool `json:"error"`
}
