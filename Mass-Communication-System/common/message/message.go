package message

// 确定消息类型
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// 定义两个消息，后面需要再增加
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户Id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回的状态码 500 表示用户未注册，200 表示成功
	Error string `json:"error"` // 返回错误信息
}
