package controller

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error"`
}

type ResponseWithToken struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}
