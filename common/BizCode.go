package common

import "errors"

const (
	NotFound          = 10000 //资源不存在
	ErrorCode         = 500   //服务器内部错误
	DataAlreadyExists = 10001
)

func Foo() error {
	return errors.New("自定义错误消息")
}
