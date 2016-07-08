package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type SunnyUsergroupAndPower struct {
	PowerId     int `orm:"column(power_id);pk"`
	UserGroupId int `orm:"column(user_group_id);"`
}

func (t *SunnyUsergroupAndPower) TableName() string {
	return "sunny_usergroup_and_power"
}
func init() {
	orm.RegisterModel(new(SunnyUsergroupAndPower))
}

func GetSunnyUsergroupAndPowerById(id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var usergroupAndPower []SunnyUsergroupAndPower
	strquery := "select * from sunny_usergroup_and_power where user_group_id="
	strquery += strconv.Itoa(id)
	num, _ := o.Raw(strquery).QueryRows(&usergroupAndPower)
	if num > 0 {
		for _, v := range usergroupAndPower {
			ml = append(ml, v)
		}
		return ml, nil
	} else {
		return nil, errors.New("not find")
	}
}

//only insert one
func AddSunnyUsergroupAndPower(m *SunnyUsergroupAndPower) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//Delete
func DeleteSunnyUsergroupAndPower(id int) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("Delete from sunny_usergroup_and_power where  user_group_id = ?", id).Exec()
	return
}
