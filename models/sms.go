package models

import (
	"github.com/NJU-Echome/dragon-back-end/util"
	"github.com/astaxie/beego/orm"
	"time"
)

type Sms struct {
	Id          int
	Phonenumber string `orm:"size(11)"`
	Code        string `orm:"size(6)"`
	Expire_time int
	State       string    `orm:default('NOT_VER')`
	Ver_time    time.Time `orm:"type(datetime)"`
	Send_time   time.Time `orm:"auto_now_add;type(datetime)"`
}

func AddSms(phonenumber string, code string, expire_time int) {
	var sms Sms
	sms.Code = code
	sms.Phonenumber = phonenumber
	sms.Expire_time = expire_time
	o := orm.NewOrm()
	o.Insert(&sms)
}

func Verify(phonenumber string, code string) (bool, string) {
	// SELECT LAST RECORD
	var sms Sms
	o := orm.NewOrm()
	err := o.QueryTable("sms").Filter("phonenumber", phonenumber).OrderBy("-id").Limit(1).One(&sms)
	if err == orm.ErrNoRows {
		return false, "未发送验证码"
	}
	interval := util.GetIntervalToNow(sms.Send_time)
	if interval >= sms.Expire_time {
		return false, "验证码已过期"
	}
	if code == sms.Code {
		sms.Ver_time = time.Now()
		sms.State = "SUCCESS"
		o.Update(&sms)
		return true, ""
	}
	sms.Ver_time = time.Now()
	sms.State = "FAIL"
	o.Update(&sms)
	return false, "验证码错误"
}
