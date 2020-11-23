package controller

type Response struct {
	Code  int
	Msg   string
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type ResponseWithToken struct {
	Code  int
	Token string
}
