package models

import (
	"fmt"
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

func GetStoresByCondition(tagIds []int, city string, district string) []Store {
	o := orm.NewOrm()
	var stores []Store

	sql := "select t1.id,name,address,latitude,longitude from store_tags t0 join store t1 on t1.id = t0.store_id where t1.city = '" + city + "' "
	if district != "" {
		sql += " and t1.district = '" + district + "' "
	}
	for i, value := range tagIds {
		if i == 0 {
			sql += " and tag_id = " + fmt.Sprintf("%d", value)
		} else {
			sql += " and exists(select 1 from store_tags where store_tags.tag_id = " + fmt.Sprintf("%d", value) + " and store_tags.store_id = t0.store_id)"
		}

	}
	o.Raw(sql).QueryRows(&stores)
	// qs := o.QueryTable("store")
	// if district == "" {
	// 	qs.Filter("Tags__Tag__Id__in", tagIds).Filter("city", city).All(&stores, "Id", "Name", "Address", "Latitude", "Longitude")
	// } else {
	// 	qs.Filter("Tags__Tag__Id__in", tagIds).Filter("city", city).Filter("district", district).All(&stores, "Id", "Name", "Address", "Latitude", "Longitude")
	// }
	for index, store := range stores {
		o.LoadRelated(&store, "Tags")
		stores[index] = store
	}
	return stores

}
func GetStoresWithNoTag(city string, district string) []Store {
	o := orm.NewOrm()
	var stores []Store
	qs := o.QueryTable("store")
	if district == "" {
		qs.Filter("city", city).All(&stores, "Id", "Name", "Address", "Latitude", "Longitude")
	} else {
		qs.Filter("city", city).Filter("district", district).All(&stores, "Id", "Name", "Address", "Latitude", "Longitude")
	}
	for index, store := range stores {
		o.LoadRelated(&store, "Tags")
		stores[index] = store
	}
	return stores
}
