package httplib

type DefaultResponse struct {
	Code    int         `json:"stat_code"`
	Message string      `json:"stat_msg"`
	Data    interface{} `json:"data"`
}
