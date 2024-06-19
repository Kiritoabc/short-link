package database

import (
	"fmt"
	"github.com/Kiritoabc/short-link/backend/pkg/config"
	"github.com/Kiritoabc/short-link/backend/pkg/dao"
	"gorm.io/gorm"
	"strings"
)

var ShortLinkDB *gorm.DB

// Init  初始化数据库
func Init() error {
	var databaseDriver driver
	var err error
	databaseDriver, err = getDriver(config.DatabaseDriver.Value)
	if err != nil {
		return err
	}
	ShortLinkDB, err = getGormDB(getShortLinkDatabaseConfig, databaseDriver)
	if err != nil {
		return err
	}
	// 设置数据库
	dao.SetDefault(ShortLinkDB)
	return nil
}

func getShortLinkDatabaseConfig() string {
	var username, password, host, port, scheme, args string
	username = config.SLDatabaseUsername.Value
	password = config.SLDatabasePassword.Value
	host = config.SLDatabaseHost.Value
	port = config.SLDatabasePort.Value
	scheme = config.SLDatabaseScheme.Value
	args = config.SLDatabaseArgs.Value
	return parse2Dsn(username, password, host, port, scheme, args)
}

func getGormDB(getDsnFunc func() string, databaseDriver driver) (*gorm.DB, error) {
	dsn := getDsnFunc()
	return gorm.Open(databaseDriver.openFunc()(dsn), &gorm.Config{})
}

func parse2Dsn(username, password, host, port, schema, args string) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, schema)
	if strings.EqualFold(args, "") {
		return dsn
	}
	return fmt.Sprintf("%s?%s", dsn, args)
}
