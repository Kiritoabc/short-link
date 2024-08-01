package shortUrl

import "github.com/Kiritoabc/short-link/backend/pkg/dto"

// NewShortUrlHandler 创建短链接处理器
func NewShortUrlHandler() ShortURLIntf {
	return &shortURLHandler{}
}

// ShortURLIntf 短链接处理器接口
type ShortURLIntf interface {
	GetLink(shortLink string) (*dto.ShortLinkDTO, error)
	GenerateShortLink(url string) (*dto.ShortLinkDTO, error)
}
