package dto

import (
	"fmt"
	"strings"
)

type Code int

const (

	// CodeOK Code 请求成功
	CodeOK Code = iota
	// CodeInvalidRequest Code 请求参数错误
	CodeInvalidRequest
	// CodeGenerateShortURLFailed Code 生成短链接失败
	CodeGenerateShortURLFailed
	// CodeGetOriginalURLFailed Code 获取原始链接失败
	CodeGetOriginalURLFailed
)

var codeMessageMap = map[Code]string{
	CodeOK:                     "请求成功",
	CodeInvalidRequest:         "请求参数错误",
	CodeGenerateShortURLFailed: "生成短链接失败",
	CodeGetOriginalURLFailed:   "获取原始链接失败",
}

// ToMessage 获取错误信息
//
//	@receiver c Code
//	@param additionalMessage ...string
//	@return string
//	@author centonhuang
//	@update 2024-12-05 16:08:28
func (c Code) ToMessage(additionalMessage ...string) string {
	message, ok := codeMessageMap[c]
	if !ok {
		message = "未知错误"
	}
	return fmt.Sprintf("%s: %s", message, strings.Join(additionalMessage, " "))
}

// StandardResponse 标准响应
//
//	@author centonhuang
//	@update 2024-12-05 16:08:35
type StandardResponse struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
