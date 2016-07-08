package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunnyEditor struct {
	Id          int    `orm:"column(id);auto"`
	Username    string `orm:"column(username);size(20)"`
	Password    string `orm:"column(password);size(32)"`
	Power       string `orm:"column(power);size(50)"`
	Description string `orm:"column(description);size(500);null"`
	Avatar      string `orm:"column(avatar);size(255);null"`
	Status      int    `orm:"column(status)"`
}

func (t *SunnyEditor) TableName() string {
	return "sunny_editor"
}

func init() {
	orm.RegisterModel(new(SunnyEditor))
}

// AddSunnyEditor insert a new SunnyEditor into database and returns
// last inserted Id on success.
func AddSunnyEditor(m *SunnyEditor) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunnyEditorById retrieves SunnyEditor by Id. Returns error if
// Id doesn't exist
func GetSunnyEditorById(id int) (v *SunnyEditor, err error) {
	o := orm.NewOrm()
	v = &SunnyEditor{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetSunnyEditorById retrieves SunnyEditor by username and password. Returns error if
// Id doesn't exist
func GetSunnyEditorByUsernameAndPwd(username string, password string) (v []*SunnyEditor) {
	o := orm.NewOrm()
	user := new(SunnyEditor)
	qs := o.QueryTable(user)
	qs.Filter("username", username).Filter("password", password).All(&v)
	if len(v) == 0 {
		return nil
	}

	return v
}

//get Editor by username
func GetSunnyEditorByUsername(username string) (v []*SunnyEditor) {
	o := orm.NewOrm()
	user := new(SunnyEditor)
	qs := o.QueryTable(user)
	qs.Filter("username", username).All(&v)
	if len(v) == 0 {
		return nil
	}

	return v
}

//判断用户是不是已经存在
func IsExistEditorByUsername(username string) (b bool) {
	var v []*SunnyEditor
	o := orm.NewOrm()
	user := new(SunnyEditor)
	qs := o.QueryTable(user)
	qs.Filter("username", username).All(&v)
	if len(v) == 0 {
		b = false
	} else {
		b = true
	}
	return
}

// GetSunnyEditorById retrieves SunnyEditor by id and password. Returns error if
// Id doesn't exist
func GetSunnyEditorByIdAndPwd(id int, password string) (v []*SunnyEditor) {
	o := orm.NewOrm()
	user := new(SunnyEditor)
	qs := o.QueryTable(user)
	qs.Filter("id", id).Filter("password", password).All(&v)
	if len(v) == 0 {
		return nil
	}

	return v
}

// GetAllSunnyEditor retrieves all SunnyEditor matches certain condition. Returns empty list if
// no records exist
func GetAllSunnyEditor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunnyEditor))
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

	var l []SunnyEditor
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

//取得用户列表
func GetAllEditor(fields []string, offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunnyEditor))

	var l []SunnyEditor
	if _, err := qs.Filter("Status", 1).Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateSunnyEditor updates SunnyEditor by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunnyEditorById(m *SunnyEditor) (err error) {
	o := orm.NewOrm()
	v := SunnyEditor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Password", "Description"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//Delete editor  用户删除实际就是修改用户的状态
func UpdateDelSunnyEditorById(m *SunnyEditor) (err error) {
	o := orm.NewOrm()
	v := SunnyEditor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Status"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return

}

//only update password
func UpdatePasswordById(m *SunnyEditor) (err error) {
	o := orm.NewOrm()
	v := SunnyEditor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "password"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunnyEditor deletes SunnyEditor by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunnyEditor(id int) (err error) {
	o := orm.NewOrm()
	v := SunnyEditor{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunnyEditor{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
