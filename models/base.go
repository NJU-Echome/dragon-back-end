package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db_user := beego.AppConfig.String("mysqluser")
	db_pass := beego.AppConfig.String("mysqlpass")
	db_name := beego.AppConfig.String("mysqldb")
	db_urls := beego.AppConfig.String("mysqlurls")
	db_max_idle_conn, _ := beego.AppConfig.Int("db_max_idle_conn")
	db_max_open_conn, _ := beego.AppConfig.Int("db_max_open_conn")
	orm.RegisterDataBase("default", "mysql", db_user+":"+db_pass+"@tcp("+db_urls+")/"+db_name+"?charset=utf8&loc=Local", db_max_idle_conn, db_max_open_conn)

	orm.Debug = true
	// register model
	orm.RegisterModel(
		new(Sms), new(Sms_send_record), new(Store), new(Tag),
	)

	// create table
	// orm.RunSyncdb("default", false, true)
}
