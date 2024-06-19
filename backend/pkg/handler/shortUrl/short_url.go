package shortUrl

import (
	"github.com/Kiritoabc/short-link/backend/pkg/dao"
	"github.com/Kiritoabc/short-link/backend/pkg/dto"
	"github.com/Kiritoabc/short-link/backend/pkg/model"
	"github.com/Kiritoabc/short-link/backend/pkg/utils"
	"github.com/Kiritoabc/short-link/backend/pkg/utils/snowflake"
	"github.com/jinzhu/copier"
)

type shortURLHandler struct{}

func (s *shortURLHandler) GetLink(shortLink string) (*dto.ShortLinkDTO, error) {
	// 从数据库中获取
	link, err := dao.ShortLink.Select(dao.ShortLink.OriginalURL).Where(dao.ShortLink.ShortCode.Eq(shortLink)).First()
	if err != nil {
		return nil, err
	}
	var shortLinkDTO dto.ShortLinkDTO
	if err = copier.Copy(&shortLinkDTO, link); err != nil {
		return nil, err
	}
	return &shortLinkDTO, nil
}

func (s *shortURLHandler) GenerateShortLink(url string) (*dto.ShortLinkDTO, error) {
	// snowflake
	id := snowflake.Snowflake.Generate()
	id62Str := utils.To62RadixString(id.Int64())
	// 存入数据库
	//writeDB := dao.Q.WriteDB()
	err := dao.ShortLink.Create(&model.ShortLink{
		OriginalURL: url,
		ShortCode:   id62Str,
		CreatedAt:   utils.GetCurrentTime(),
	})
	if err != nil {
		return nil, err
	}
	return &dto.ShortLinkDTO{
		OriginalURL: url,
		ShortCode:   id62Str,
	}, nil
}
