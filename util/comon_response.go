package util

type CommonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CommonResponse 统一返回
func CommonResponse(code int, message string, data interface{}) *CommonResp {
	return &CommonResp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
