package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"superStar/conf"
	"sync"
)

var (
	masterDb *gorm.DB
	slaveDb  *gorm.DB
	lock     sync.Mutex
)

func InstanceMaster() *gorm.DB {
	if masterDb != nil {
		return masterDb
	}

	lock.Lock()
	defer lock.Unlock()

	if masterDb != nil {
		return masterDb
	}

	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf-8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)

	masterDb, err := gorm.Open(conf.DriverName, driveSource)
	if err != nil {
		fmt.Println("masterDb conn err=", err)
		return nil
	}
	return masterDb
}

func InstanceSlave() *gorm.DB {
	if slaveDb != nil {
		return slaveDb
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveDb != nil {
		return slaveDb
	}

	c := conf.SlaveDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf-8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	slaveDb, err := gorm.Open(conf.DriverName, driveSource)

	if err != nil {
		fmt.Println("slaveDb conn err=", err)
		return nil
	}
	return slaveDb
}
