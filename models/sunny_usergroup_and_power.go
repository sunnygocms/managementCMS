package models

type SunnyUsergroupAndPower struct {
	PowerId     int `orm:"column(power_id)"`
	UserGroupId int `orm:"column(user_group_id)"`
}
