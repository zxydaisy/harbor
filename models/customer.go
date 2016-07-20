package models

type Customer struct {
	Id int    `orm:"column(id);auto"`
	Name  string `orm:"column(name);size(32)"`
	Tag  string `orm:"column(tag);size(32)"`
}
