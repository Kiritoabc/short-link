package shortUrl

func NewShortUrlHandler() ShortURLIntf {
	return &shortURLHandler{}
}

type ShortURLIntf interface {
	GetShortLink(url string) (string, error)
}
