// Package dto 数据传输对象
//
//	@update 2024-12-06 19:57:57
package dto

import (
	"fmt"
	"strings"
)

// Code 错误码
//
//	@author centonhuang
//	@update 2024-12-06 19:57:39
type Code int

const (
	// CodeUnauthorized Code 未认证
	CodeUnauthorized Code = -1001
	// CodeForbidden Code 禁止访问
	CodeForbidden Code = -1002
	// CodeOK Code 请求成功
	CodeOK Code = 0
	// CodeInvalidRequest Code 请求参数错误
	CodeInvalidRequest Code = 1001
	// CodeGenerateShortURLFailed Code 生成短链接失败
	CodeGenerateShortURLFailed Code = 1002
	// CodeGetOriginalURLFailed Code 获取原始链接失败
	CodeGetOriginalURLFailed Code = 1003
)

var codeMessageMap = map[Code]string{
	CodeUnauthorized:           "未认证",
	CodeForbidden:              "禁止访问",
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

	if len(additionalMessage) > 0 {
		return fmt.Sprintf("%s: %s", message, strings.Join(additionalMessage, " "))
	}

	return message
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
