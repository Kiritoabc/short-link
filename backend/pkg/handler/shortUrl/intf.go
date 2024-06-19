package shortUrl

import "github.com/Kiritoabc/short-link/backend/pkg/dto"

func NewShortUrlHandler() ShortURLIntf {
	return &shortURLHandler{}
}

type ShortURLIntf interface {
	GetLink(shortLink string) (*dto.ShortLinkDTO, error)
	GenerateShortLink(url string) (*dto.ShortLinkDTO, error)
}
