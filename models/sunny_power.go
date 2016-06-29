package models

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
)

//type UserPower struct {
//	Element map[string][]string
//}
var mycache, _ = cache.NewCache("memory", `{"interval":600}`)

type SunnyPower struct {
	Id         int    `orm:"column(power_id);auto"`
	PowerName  string `orm:"column(power_name);size(255);null"`
	Controller string `orm:"column(controller);size(255);null"`
	Action     string `orm:"column(action);size(255);null"`
}

func (t *SunnyPower) TableName() string {
	return "sunny_power"
}

func init() {

	orm.RegisterModel(new(SunnyPower))
}

// AddSunnyPower insert a new SunnyPower into database and returns
// last inserted Id on success.
func AddSunnyPower(m *SunnyPower) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunnyPowerById retrieves SunnyPower by Id. Returns error if
// Id doesn't exist
func GetSunnyPowerById(id int) (v *SunnyPower, err error) {
	o := orm.NewOrm()
	v = &SunnyPower{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//得到权限map
func GetEditorPowersById(id int) (userpower interface{}, err error) { //
	key := "power" + strconv.Itoa(id)
	if mycache.IsExist(key) {
		return mycache.Get(key), nil
	} else {
		o := orm.NewOrm()
		var power []SunnyPower
		var buf bytes.Buffer
		var result = make(map[string][]string)
		buf.WriteString("SELECT p.power_id AS id,p.power_name AS power_name,p.controller AS controller,p.action as action ")
		buf.WriteString("FROM sunny_power AS p RIGHT JOIN ( ")
		buf.WriteString("SELECT DISTINCT(gp.power_id) AS power_id ")
		buf.WriteString("FROM sunny_user_and_group ug ")
		buf.WriteString("RIGHT JOIN sunny_usergroup_and_power gp ON ug.user_group_id = gp.user_group_id ")
		buf.WriteString("WHERE ug.user_id =  ")
		buf.WriteString(strconv.Itoa(id))
		buf.WriteString(") k ON p.power_id = k.power_id ")
		num, _ := o.Raw(buf.String()).QueryRows(&power)
		if num > 0 {
			for _, p := range power {
				result[p.Controller] = append(result[p.Controller], p.Action)
			}
			mycache.Put(key, result, 600*time.Second)
			return result, nil
		} else {
			return nil, errors.New("not find")
		}
	}

}

// GetAllSunnyPower retrieves all SunnyPower matches certain condition. Returns empty list if
// no records exist
func GetAllSunnyPower(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunnyPower))
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

	var l []SunnyPower
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

// UpdateSunnyPower updates SunnyPower by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunnyPowerById(m *SunnyPower) (err error) {
	o := orm.NewOrm()
	v := SunnyPower{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunnyPower deletes SunnyPower by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunnyPower(id int) (err error) {
	o := orm.NewOrm()
	v := SunnyPower{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunnyPower{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//get power array
func GetEditorPowers(editorID int) (ml []interface{}, err error) {
	var num int64
	b := bytes.Buffer{}
	var powMap map[string]interface{}
	powMap = make(map[string]interface{})
	b.WriteString("SELECT p.power_id AS id,p.power_name AS name,p.controller AS controller,p.action ")
	b.WriteString("FROM sunny_power AS p RIGHT JOIN ( ")
	b.WriteString("SELECT DISTINCT(gp.power_id) AS power_id ")
	b.WriteString("FROM sunny_user_and_group ug ")
	b.WriteString("RIGHT JOIN sunny_usergroup_and_power gp ON ug.user_group_id = gp.user_group_id ")
	b.WriteString("WHERE ug.user_id =  ")
	b.WriteString(strconv.Itoa(editorID))
	b.WriteString(") k ")
	b.WriteString("ON p.power_id = k.power_id ")
	//	var fields []string
	//	var power []SunnyPower
	o := orm.NewOrm()
	var maps []orm.Params
	num, err = o.Raw(b.String()).Values(&maps)
	_ = num
	stmp := ""
	var powerinterface []interface{}
	i := 0
	if err == nil {
		for num, v := range maps {
			if v["controller"].(string) == stmp {
				powerinterface = append(powerinterface, v["action"].(string))
			} else {
				if len(stmp) > 0 {
					powMap[stmp] = powerinterface
					ml = append(ml, powMap)
				}
				stmp = v["controller"].(string)
				powMap = make(map[string]interface{})
				powerinterface = powerinterface[:0]
				powerinterface = append(powerinterface, v["action"].(string))
			}
			i++
			if i == num {
				powMap[stmp] = powerinterface
				ml = append(ml, powMap)
			}
		}
		return ml, nil
	} else {
		return nil, err
	}

}
