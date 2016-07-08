package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type SunnyUserAndGroup struct {
	UserId      int `orm:"column(user_id);pk"`
	UserGroupId int `orm:"column(user_group_id);"`
}

func (t *SunnyUserAndGroup) TableName() string {
	return "sunny_user_and_group"
}
func init() {
	orm.RegisterModel(new(SunnyUserAndGroup))
}

func GetSunnyUserAndGroupById(id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var usergroup []SunnyUserAndGroup
	strquery := "select * from sunny_user_and_group where user_id="
	strquery += strconv.Itoa(id)
	num, _ := o.Raw(strquery).QueryRows(&usergroup)
	if num > 0 {
		for _, v := range usergroup {
			ml = append(ml, v)
		}
		return ml, nil
	} else {
		return nil, errors.New("not find")
	}
}

//only insert one
func AddSunnyUserAndGroup(m *SunnyUserAndGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//insert multi
func AddSunnyUserAndGroups(m []SunnyUserAndGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.InsertMulti(len(m), m)
	return
}

//Delete
func DeleteSunnyUserAndGroup(id int) (err error) {
	o := orm.NewOrm()
	v := SunnyUserAndGroup{UserId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&SunnyUserAndGroup{UserId: id})
	}
	return
}

//Delete SunnyUserAndGroup By UserGroupId
func DeleteSunnyUserAndGroupByUserGroupId(id int) (err error) {
	o := orm.NewOrm()
	v := SunnyUserAndGroup{UserGroupId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&SunnyUserAndGroup{UserGroupId: id})
	}
	return
}
