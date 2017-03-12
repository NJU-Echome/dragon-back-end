package models

import (
	"time"
)

type Tag struct {
	Id          uint32
	Name        string    `orm:"size(15)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
	Stores      []*Store  `orm:"reverse(many)"`
}
