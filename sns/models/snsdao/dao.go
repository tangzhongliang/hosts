package snsdao

import (
	"os"
	"sns/models/beans"
	"sns/util/snserror"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func New() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "user:iser@tcp(172.25.78.80:3306)/sns?charset=utf8")
	snserror.LogAndPanic(err)
	f, err2 := os.Create("sql.log")
	defer f.Close()
	snserror.LogAndPanic(err2)
	//	engine.Logger = xorm.NewSimpleLogger(f)
	//	engine.SetMaxIdleConns(3000)
	//	engine.SetMaxOpenConns(3000)
	return engine
}

func Test() {
	engine, err := xorm.NewEngine("mysql", "root:root@/sns?charset=utf8")
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	err = engine.CreateTables(&beans.SNSEPAccount{})
	snserror.LogAndPanic(err)
	engine.Close()
}
