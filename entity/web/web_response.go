package web

type WebResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
