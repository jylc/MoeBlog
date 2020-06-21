package resp

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
}

const (
	//CodeSuccess 注册成功
	CodeSuccess = 200
	//CodeSFailed 注册失败
	CodeFailed = 201
	//CodeNotFullSuccess 未完全成功
	CodeNotFullSuccess = 203
	//CodeCheckLogin 未登录
	CodeCheckLogin = 401
	//CodeNotFound 资源未找到
	CodeNotFound = 404
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

//参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}

//通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
		Err:  err.Error(),
	}

	return res
}
