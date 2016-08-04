package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunnyNavigation struct {
	Id         int    `orm:"column(id);auto"`
	Level      int    `orm:"column(level)"`
	ParentId   int    `orm:"column(parent_id)"`
	Controller string `orm:"column(controller);size(32);null"`
	Action     string `orm:"column(action);size(32);null"`
	Name       string `orm:"column(name);size(32)"`
	Sort       int    `orm:"column(sort)"`
	Display    int8   `orm:"column(display)"`
}

func (t *SunnyNavigation) TableName() string {
	return "sunny_navigation"
}

func init() {
	orm.RegisterModel(new(SunnyNavigation))
}

// AddSunnyNavigation insert a new SunnyNavigation into database and returns
// last inserted Id on success.
func AddSunnyNavigation(m *SunnyNavigation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunnyNavigationById retrieves SunnyNavigation by Id. Returns error if
// Id doesn't exist
func GetSunnyNavigationById(id int) (v *SunnyNavigation, err error) {
	o := orm.NewOrm()
	v = &SunnyNavigation{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunnyNavigation retrieves all SunnyNavigation matches certain condition. Returns empty list if
// no records exist
func GetAllSunnyNavigation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunnyNavigation))
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

	var l []SunnyNavigation
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

// UpdateSunnyNavigation updates SunnyNavigation by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunnyNavigationById(m *SunnyNavigation) (err error) {
	o := orm.NewOrm()
	v := SunnyNavigation{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunnyNavigation deletes SunnyNavigation by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunnyNavigation(id int) (err error) {
	o := orm.NewOrm()
	v := SunnyNavigation{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunnyNavigation{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetNavigationAll() (ml []SunnyNavigation, err error) {
	o := orm.NewOrm()
	var data []SunnyNavigation
	num, _ := o.Raw("select * from sunny_navigation ORDER BY level asc,parent_id asc,sort asc,id asc").QueryRows(&data)
	if num > 0 {
		for _, v := range data {
			ml = append(ml, v)
		}
		return ml, nil
	} else {
		return nil, errors.New("not find")
	}
}
