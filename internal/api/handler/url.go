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
//	@param handler *shortURLHandler
//	@author centonhuang
//	@update 2024-12-05 16:09:55
type ShortURLHandlerOption func(handler *shortURLHandler)

// WithShortURLService 设置短URL服务
//
//	@param service service.ShortURLService
//	@return ShortURLHandlerOption
//	@author centonhuang
//	@update 2024-12-05 16:09:59
func WithShortURLService(service service.ShortURLService) ShortURLHandlerOption {
	return func(handler *shortURLHandler) {
		handler.service = service
	}
}

// NewShortURLHandler 创建短URL接口
//
//	@param handlerOptions ...ShortURLHandlerOption
//	@return ShortURLHandler
//	@author centonhuang
//	@update 2024-12-05 16:12:10
func NewShortURLHandler(handlerOptions ...ShortURLHandlerOption) ShortURLHandler {
	handler := &shortURLHandler{}
	for _, option := range handlerOptions {
		option(handler)
	}
	return handler
}

// GenerateShortURL 生成短URL
//
//	@receiver h *shortURLHandler
//	@param ctx *fiber.Ctx
//	@return error
//	@author centonhuang
//	@update 2024-12-05 16:12:27
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
//	@receiver h *shortURLHandler
//	@param ctx *fiber.Ctx
//	@return error
//	@author centonhuang
//	@update 2024-12-05 16:12:37
func (h *shortURLHandler) GetOriginalURL(ctx *fiber.Ctx) error {
	request := new(dto.GetOriginalURLRequest)
	if err := ctx.ParamsParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
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
