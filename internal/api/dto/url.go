package dto

// GenerateShortURLRequest 生成短URL请求
//
//	author centonhuang
//	update 2024-12-05 16:09:04
type GenerateShortURLRequest struct {
	OriginalURL string `json:"originalURL" validate:"required,min=5,max=512"`
	ExpireDays  int8   `json:"expireDays" validate:"omitempty,min=1"`
}

// GenerateShortURLResponse 生成短URL响应
//
//	author centonhuang
//	update 2024-12-05 16:09:00
type GenerateShortURLResponse struct {
	ShortURL string `json:"shortURL"`
}

// GetOriginalURLRequest 获取原始URL请求
//
//	author centonhuang
//	update 2024-12-05 16:08:54
type GetOriginalURLRequest struct {
	ShortURL string `json:"shortURL" validate:"required,min=5,max=10"`
}

// GetOriginalURLResponse 获取原始URL响应
//
//	author centonhuang
//	update 2024-12-05 16:08:49
type GetOriginalURLResponse struct {
	OriginalURL string `json:"originalURL"`
}
