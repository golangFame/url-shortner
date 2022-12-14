package entities

type GenericResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	StatusCode int         `json:"-"`
}
