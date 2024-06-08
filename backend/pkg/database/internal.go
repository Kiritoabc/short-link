package database

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

import "gorm.io/driver/mysql"

var drivers = []driver{
	&mysqlDriver{},
}

func getDriver(name string) (driver, error) {
	for _, dri := range drivers {
		if strings.EqualFold(dri.driverName(), name) {
			return dri, nil
		}
	}
	return nil, fmt.Errorf("un support driver of %s", name)
}

type driver interface {
	driverName() string
	openFunc() func(dsn string) gorm.Dialector
}

type mysqlDriver struct{}

func (*mysqlDriver) driverName() string {
	return "mysql"
}

func (*mysqlDriver) openFunc() func(dsn string) gorm.Dialector {
	return mysql.Open
}
