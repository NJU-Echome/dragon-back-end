package models

import (
	"github.com/NJU-Echome/dragon-back-end/util"
	"github.com/astaxie/beego/orm"
	"time"
)

type Sms_send_record struct {
	Id          int
	Phonenumber string    `orm:"size(11)"`
	Send_time   time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetTodaySendCount(phonenumber string) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("sms_send_record").Filter("phonenumber", phonenumber).Filter("send_time__gte", util.GetTodayDate()).Count()
	return cnt
}
func AddRecord(phonenumber string) {
	var record Sms_send_record
	record.Phonenumber = phonenumber
	// record.Send_time = time.Now()
	o := orm.NewOrm()
	o.Insert(&record)

}
