package dto

import (
	"fmt"
	"strings"
)

type Code int

const (
	CodeOK Code = iota
	CodeInvalidRequest
	CodeGenerateShortURLFailed
	CodeGetOriginalURLFailed
)

var (
	codeMessageMap = map[Code]string{
		CodeOK:                     "请求成功",
		CodeInvalidRequest:         "请求参数错误",
		CodeGenerateShortURLFailed: "生成短链接失败",
		CodeGetOriginalURLFailed:   "获取原始链接失败",
	}
)

func (c Code) ToMessage(additionalMessage ...string) string {
	message, ok := codeMessageMap[c]
	if !ok {
		message = "未知错误"
	}
	return fmt.Sprintf("%s: %s", message, strings.Join(additionalMessage, " "))
}

type StandardResponse struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
