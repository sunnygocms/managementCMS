package models

type SunnyUserAndGroup struct {
	UserId      int `orm:"column(user_id)"`
	UserGroupId int `orm:"column(user_group_id)"`
}
