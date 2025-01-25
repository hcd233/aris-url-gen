package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/Aris-url-gen/internal/api/dto"
	"github.com/hcd233/Aris-url-gen/internal/api/service"
)

type shortURLHandler struct {
	service service.ShortURLService
}

// ShortURLHandlerOption 短URL接口选项
//
//	param handler *shortURLHandler
//	author centonhuang
//	update 2024-12-05 16:09:55
type ShortURLHandlerOption func(handler *shortURLHandler)

// WithShortURLService 设置短URL服务
//
//	param service service.ShortURLService
//	return ShortURLHandlerOption
//	author centonhuang
//	update 2024-12-05 16:09:59
func WithShortURLService(service service.ShortURLService) ShortURLHandlerOption {
	return func(handler *shortURLHandler) {
		handler.service = service
	}
}

// NewShortURLHandler 创建短URL接口
//
//	param handlerOptions ...ShortURLHandlerOption
//	return ShortURLHandler
//	author centonhuang
//	update 2024-12-05 16:12:10
func NewShortURLHandler(handlerOptions ...ShortURLHandlerOption) ShortURLHandler {
	handler := &shortURLHandler{}
	for _, option := range handlerOptions {
		option(handler)
	}
	return handler
}

// GenerateShortURL 生成短URL
//
// @Summary 生成短URL
// @Description 将长URL转换为短URL
// @Tags 短链接
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body dto.GenerateShortURLRequest true "生成短URL请求"
// @Success 200 {object} dto.StandardResponse{data=dto.GenerateShortURLResponse} "成功"
// @Failure 400 {object} dto.StandardResponse "请求参数错误"
// @Failure 500 {object} dto.StandardResponse "服务器内部错误"
// @Router /v1/shortURL [post]
func (h *shortURLHandler) GenerateShortURL(ctx *fiber.Ctx) error {
	request := new(dto.GenerateShortURLRequest)
	if err := ctx.BodyParser(request); err != nil {
		code := dto.CodeInvalidRequest
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.StandardResponse{
			Code:    code,
			Message: code.ToMessage(err.Error()),
		})
	}

	response, err := h.service.GenerateShortURL(request)
	if err != nil {
		code := dto.CodeGenerateShortURLFailed
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.StandardResponse{
			Code:    code,
			Message: code.ToMessage(err.Error()),
		})
	}

	return ctx.JSON(dto.StandardResponse{
		Code:    dto.CodeOK,
		Message: dto.CodeOK.ToMessage(),
		Data:    response,
	})
}

// GetOriginalURL 获取原始URL
//
// @Summary 获取原始URL
// @Description 通过短URL获取原始URL并重定向
// @Tags 短链接
// @Accept json
// @Produce json
// @Param shortURL path string true "短URL"
// @Success 302 {string} string "重定向到原始URL"
// @Failure 400 {object} dto.StandardResponse "请求参数错误"
// @Failure 500 {object} dto.StandardResponse "服务器内部错误"
// @Router /v1/s/{shortURL} [get]
func (h *shortURLHandler) GetOriginalURL(ctx *fiber.Ctx) error {
	request := new(dto.GetOriginalURLRequest)
	if err := ctx.ParamsParser(request); err != nil {
		code := dto.CodeInvalidRequest
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.StandardResponse{
			Code:    code,
			Message: code.ToMessage(err.Error()),
		})
	}

	response, err := h.service.GetOriginalURL(request)
	if err != nil {
		code := dto.CodeGetOriginalURLFailed
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.StandardResponse{
			Code:    code,
			Message: code.ToMessage(err.Error()),
		})
	}

	return ctx.Redirect(response.OriginalURL)
}
