package main

import (
	"errors"
	"fmt"
	"github.com/Kiritoabc/short-link/backend/pkg/dao"
	"github.com/Kiritoabc/short-link/backend/pkg/database"
	"github.com/Kiritoabc/short-link/backend/pkg/model"
	"gorm.io/gorm"
	"testing"
)

func TestName(t *testing.T) {
	database.Init()
	dao.SetDefault(database.ShortLinkDB)
	// 查询
	first, err := dao.ShortLink.Select(dao.ShortLink.ID, dao.ShortLink.OriginalURL).Where(dao.ShortLink.ID.Eq(1)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("not found")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		panic(err)
	}

	// 创建
	dao.ShortLink.Create(&model.ShortLink{})

	println(first)
}
