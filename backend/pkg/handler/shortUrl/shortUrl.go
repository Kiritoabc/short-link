package shortUrl

import (
	"github.com/Kiritoabc/short-link/backend/pkg/utils"
	"github.com/Kiritoabc/short-link/backend/pkg/utils/snowflake"
)

type shortURLHandler struct{}

func (s *shortURLHandler) GetShortLink(url string) (string, error) {
	// snowflake
	id := snowflake.Snowflake.Generate()
	id62Str := utils.To62RadixString(id.Int64())
	// 存入数据库
	println(id62Str)
	return id62Str, nil
}
