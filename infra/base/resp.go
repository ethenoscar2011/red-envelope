package base

type Resp struct {
	Code  int64  `json:"code" desc:"错误码"`
	Msg   string `json:"msg" desc:"错误信息"`
	Data  interface{} `json:"data" desc:"接收数据类型"`
}

type ResCode int

const (
	ResCodeOk                 = 1000
	ResCodeValidationError    = 2000
	ResCodeRequestParamsError = 2100
	ResCodeInnerServerError   = 5000
	ResCodeBizError           = 6000
)