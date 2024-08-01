package dto

// ShortLinkDTO 短链接数据传输对象
type ShortLinkDTO struct {
	ID          int32  `json:"id"`           // 生成的数据库ID
	OriginalURL string `json:"original_url"` // 原始的URL
	ShortCode   string `json:"short_code"`   // 短链接
}
