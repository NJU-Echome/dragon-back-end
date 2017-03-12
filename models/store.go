package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Store struct {
	Id             uint32
	Name           string `orm:"size(15)"`
	Address        string `orm:"size(63)"`
	Longitude      float32
	Latitude       float32
	City           string    `orm:"size(15)"`
	District       string    `orm:"size(15)"`
	Province       string    `orm:"size(15)"`
	Image_url      string    `orm:"size(255)"`
	Open_time      time.Time `orm:"type(datetime)"`
	Close_time     time.Time `orm:"type(datetime)"`
	Contact_number string    `orm:"size(31)"`
	Create_time    time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time    time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time    time.Time `orm:"type(datetime)"`
	Tags           []*Tag    `orm:"rel(m2m)"`
}

func GetStoresByCondition(tagId int, city string, district string) []Store {
	o := orm.NewOrm()
	var stores []Store
	qs := o.QueryTable("store")
	if district == "" {
		qs.Filter("Tags__Tag__Id", tagId).Filter("city", city).All(&stores, "Id", "Name", "Address", "Latitude", "Longitude")
	} else {
		qs.Filter("Tags__Tag__Id", tagId).Filter("city", city).Filter("district", district).All(&stores, "Id", "Name", "Address", "Latitude", "Longitude")
	}
	for index, store := range stores {
		o.LoadRelated(&store, "Tags")
		stores[index] = store
	}
	return stores

}
