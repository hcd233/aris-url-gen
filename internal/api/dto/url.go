package dto

type GenerateShortURLRequest struct {
	OriginalURL string `json:"originalURL" validate:"required,min=5,max=512"`
	ExpireDays  int8   `json:"expireDays" validate:"omitempty,min=1"`
}

type GenerateShortURLResponse struct {
	ShortURL string `json:"shortURL"`
}

type GetOriginalURLRequest struct {
	ShortURL string `json:"shortURL" validate:"required,min=5,max=10"`
}

type GetOriginalURLResponse struct {
	OriginalURL string `json:"originalURL"`
}
