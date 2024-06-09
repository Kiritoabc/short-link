package shortUrl

import "github.com/Kiritoabc/short-link/backend/pkg/utils/snowflake"

type shortURLHandler struct{}

func (s *shortURLHandler) GetShortLink(url string) (string, error) {
	// snowflake
	id := snowflake.Snowflake.Generate()
	idStr := id.String()
	println(idStr)
	return idStr, nil
}
