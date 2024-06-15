package main

import (
	"errors"
	"fmt"
	"github.com/Kiritoabc/short-link/backend/pkg/database"
	"github.com/Kiritoabc/short-link/backend/pkg/query"
	"gorm.io/gorm"
	"testing"
)

func TestName(t *testing.T) {
	database.Init()
	query.SetDefault(database.ShortLinkDB)
	first, err := query.ShortLink.Select(query.ShortLink.ID, query.ShortLink.OriginalURL).Where(query.ShortLink.ID.Eq(1)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("not found")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		panic(err)
	}

	println(first)
}
