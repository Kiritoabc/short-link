package dto

type ShortLinkDTO struct {
	ID          int32  `json:"id"`           // 生成的数据库ID
	OriginalURL string `json:"original_url"` // 原始的URL
	ShortCode   string `json:"short_code"`   // 短链接
}
