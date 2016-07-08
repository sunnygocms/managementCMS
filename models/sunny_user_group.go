package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	//	"time"

	"github.com/astaxie/beego/orm"
)

type SunnyUserGroup struct {
	Id          int    `orm:"column(user_group_id);auto"`
	GroupName   string `orm:"column(group_name);size(255);null"`
	EditId      int    `orm:"column(edit_id);null"`
	Description string `orm:"column(description);size(255);null"`
	Active      int    `orm:"column(active)";null`
}

func (t *SunnyUserGroup) TableName() string {
	return "sunny_user_group"
}

func init() {
	orm.RegisterModel(new(SunnyUserGroup))
}

// AddSunnyUserGroup insert a new SunnyUserGroup into database and returns
// last inserted Id on success.
func AddSunnyUserGroup(m *SunnyUserGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunnyUserGroupById retrieves SunnyUserGroup by Id. Returns error if
// Id doesn't exist
func GetSunnyUserGroupById(id int) (v *SunnyUserGroup, err error) {
	o := orm.NewOrm()
	v = &SunnyUserGroup{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//get user_group
func GetAllUserGroup(where string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var usergroup []SunnyUserGroup
	num, _ := o.Raw("select * from sunny_user_group " + where).QueryRows(&usergroup)
	if num > 0 {
		for _, v := range usergroup {
			ml = append(ml, v)
		}

		return ml, nil
	} else {
		return nil, errors.New("not find")
	}

}

// GetAllSunnyUserGroup retrieves all SunnyUserGroup matches certain condition. Returns empty list if
// no records exist
func GetAllSunnyUserGroup(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunnyUserGroup))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SunnyUserGroup
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSunnyUserGroup updates SunnyUserGroup by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunnyUserGroupById(m *SunnyUserGroup) (err error) {
	o := orm.NewOrm()
	v := SunnyUserGroup{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunnyUserGroup deletes SunnyUserGroup by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunnyUserGroup(id int) (err error) {
	o := orm.NewOrm()
	v := SunnyUserGroup{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunnyUserGroup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//判断用户是不是已经存在
func IsExistUsergroupByUsername(name string) (b bool) {
	var v []*SunnyUserGroup
	o := orm.NewOrm()
	user := new(SunnyUserGroup)
	qs := o.QueryTable(user)
	qs.Filter("group_name", name).All(&v)
	if len(v) == 0 {
		b = false
	} else {
		b = true
	}
	return
}
