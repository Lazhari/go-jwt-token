package models

// RequestError the error response
type RequestError struct {
	Message          string `json:"message"`
	StatusCode       int    `json:"status"`
	ValidationErrors error  `json:"errors,omitempty"`
}

func (r *RequestError) Error() string {
	return r.Message
}
