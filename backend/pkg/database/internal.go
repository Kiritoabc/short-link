package database

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

import "gorm.io/driver/mysql"

// drivers
var drivers = []driver{
	&mysqlDriver{},
}

// getDriver
func getDriver(name string) (driver, error) {
	for _, dri := range drivers {
		if strings.EqualFold(dri.driverName(), name) {
			return dri, nil
		}
	}
	return nil, fmt.Errorf("un support driver of %s", name)
}

// driver
type driver interface {
	driverName() string
	openFunc() func(dsn string) gorm.Dialector
}

// mysqlDriver
type mysqlDriver struct{}

// driverName
func (*mysqlDriver) driverName() string {
	return "mysql"
}

// openFunc
func (*mysqlDriver) openFunc() func(dsn string) gorm.Dialector {
	return mysql.Open
}
