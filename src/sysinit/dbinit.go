package sysinit

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func dbinit() {
	registerDBs("w", "r")
}

//一次注册多个数据库
func registerDBs(aliases ...string) {
	if len(aliases) == 0 {
		registerDB("w")
		return
	}

	for _, alias := range aliases {
		registerDB(alias)
	}
}

//注册单个数据库
func registerDB(alias string) {
	Alias := ""
	if len(alias) == 0 || alias == "default" {
		alias = "w"
	}

	if alias == "w" {
		Alias = "default"
	} else {
		Alias = alias
	}

	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	dbPort := beego.AppConfig.String("db_" + alias + "_port")
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	dbPwd := beego.AppConfig.String("db_" + alias + "_pwd")
	dbName := beego.AppConfig.String("db_" + alias + "_database")

	//root:my123456@tcp(127.0.0.1:3306)/mbook?charset=utf8
	dataSource := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	err := orm.RegisterDataBase(Alias, "mysql", dataSource)
	if err != nil {
		fmt.Println(err)
		panic("注册数据库失败，退出应用")
	}

	//主库需要进行同步
	if alias == "w" {
		orm.RunSyncdb(Alias, false, true)
	}
}
