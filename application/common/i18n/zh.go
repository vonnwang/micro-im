package i18n

const (
	ErrParam  string = "参数错误"
	ErrServer string = "服务器忙碌，请稍后再试"
	TimLayOut string = "2006-01-02 15:04:05"
)

var ZhMessage = map[string]string{
	"LoginRequest.Username.required": "用户名不能为空",
	"LoginRequest.Password.required": "密码不能为空",


	"RegisterRequest.Username.required": "用户名不能为空",
	"RegisterRequest.Password.required": "密码不能为空",

	//SendRequest struct {
	//		FromToken string    `json:"fromToken"  binding:"required"`
	//		ToToken   string    `json:"toToken"  binding:"required"`
	//		Body      string    `json:"body"  binding:"required"`
	//		Timestamp time.Time `json:"timestamp"`
	//	}

	"SendRequest.FromToken.required": "TOKEN不能为空",
	"SendRequest.ToToken.required":   "发送人TOKEN不能为空",
	"SendRequest.Body.required":      "消息体不能为空",
}
