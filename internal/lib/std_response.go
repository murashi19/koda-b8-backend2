package lib

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token string `json:"token,omitempty"`
	Result  any    `json:"result,omitempty"`
}
