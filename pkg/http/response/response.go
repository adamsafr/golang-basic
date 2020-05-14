package response

import (
	"encoding/json"
	"net/http"
)

// Response ...
type Response struct {
	w http.ResponseWriter
}

// Create ...
func Create(w http.ResponseWriter) *Response {
	return &Response{w}
}

// ByteToJSON ...
func (res *Response) ByteToJSON(data []byte) {
	res.SetJSONContentType()
	res.w.Write(data)
}

// ToJSON ...
func (res *Response) ToJSON(data interface{}) {
	res.SetJSONContentType()
	json.NewEncoder(res.w).Encode(data)
}

// SetJSONContentType ...
func (res *Response) SetJSONContentType() {
	res.w.Header().Set("Content-Type", "application/json")
}

func (res *Response) Error(error string, code int) {
	res.SetJSONContentType()
	res.w.WriteHeader(code)

	json.NewEncoder(res.w).Encode(struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	}{
		code,
		error,
	})
}

// BadRequest ...
func (res *Response) BadRequest(error string) {
	res.Error(error, http.StatusBadRequest)
}
