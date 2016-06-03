package controllers

//"github.com/astaxie/beego"
import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sunnygocms/managementCMS/models"
)

type IndexController struct {
	BaseController
}
type Alertpwd struct {
	Id        int    `form:"-"`
	Sourcepwd string `form:"sourcepwd"`
	Pwd       string `form:"pwd"`
	Repwd     string `form:"repwd"`
}
type BaseTree struct {
	Id         int
	ParentId   int
	Controller string
	Action     string
	Name       string
}
type Tree struct {
	Id         int
	ParentId   int
	Controller string
	Action     string
	Name       string
	Children   []Tree
}

// @router /index
func (this *IndexController) Index() {
	sess_username, _ := this.GetSession("editor_username").(string)

	if len(sess_username) == 0 {
		//		this.Html(sess_username + "00000000000-------")
		this.Ctx.Redirect(302, "/login")
	} else {
		sess_power, _ := this.GetSession("editor_power").(string)
		this.Data["editor_username"] = sess_username
		this.Data["editor_power"] = sess_power
		this.Data["tree"] = this.GetAllNagigation()
		this.TplName = "index/index.html"
	}

}
func (this *IndexController) Welcome() {
	this.TplName = "index/welcome.html"
}

//修改密码
//1）三个输入框都不能为空
//2）新密码的两个密码应该相同
//3）原密码需要正确
func (this *IndexController) Alertpwd() {

	if this.IsSubmit() {
		editorID := this.GetEditorId()
		this.Info("-------ID:" + strconv.Itoa(editorID))
		alp := Alertpwd{}
		if err := this.ParseForm(&alp); err != nil {
			this.Info(err)
		} else {
			if len(alp.Repwd) == 0 || len(alp.Pwd) == 0 || len(alp.Sourcepwd) == 0 {
				this.Error("密码不能够为空", "-1", 4)
			} else if alp.Pwd != alp.Repwd {
				this.Error("新的密码两次输入不同", "-1", 4)
			} else {
				v := models.GetSunnyEditorByIdAndPwd(editorID, this.SunnyMd5(alp.Sourcepwd))
				if v == nil {
					this.Error("原密码不正确，不能够修改！", "-1", 4)
				} else {
					var se models.SunnyEditor
					se.Id = editorID
					se.Password = this.SunnyMd5(alp.Pwd)
					if err = models.UpdatePasswordById(&se); err != nil {
						this.Info(err) //e10adc3949ba59abbe56e057f20f883e
					} else {
						this.Success("成功了", "-1", 4)
					}

				}
			}
		}
	} else {
		this.TplName = "index/alertPWD.html"
	}
}

func (this *IndexController) GetAllNagigation() string {
	var fields []string
	var sortby []string
	var order []string
	var snTree []Tree
	var query map[string]string = make(map[string]string)
	var limit int64 = 0 //if Limit <= 0 then Limit will be set to default limit ,eg 1000
	var offset int64 = 0
	v := "display:1"
	for _, cond := range strings.Split(v, ",") {
		kv := strings.Split(cond, ":")
		if len(kv) != 2 {
			this.Data["json"] = errors.New("Error: invalid query key/value pair")
			return ""
		}
		k, v := kv[0], kv[1]
		query[k] = v
	}
	//id,parent_id,module,action,name
	fields = strings.Split("Id,ParentId,Controller,Action,Name", ",")
	//level asc,parent_id asc,sort asc,id asc
	sortby = strings.Split("level,parent_id,sort,id", ",")
	order = strings.Split("asc,asc,asc,asc", ",")
	l, err := models.GetAllSunnyNavigation(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		for _, v := range l {
			sn := v.(map[string]interface{})
			if sn["ParentId"].(int) == 0 {
				var tree Tree
				tree.Id = sn["Id"].(int)
				tree.ParentId = sn["ParentId"].(int)
				tree.Action = sn["Action"].(string)
				tree.Controller = sn["Controller"].(string)
				tree.Name = sn["Name"].(string)
				tree.Children = nil
				snTree = append(snTree, tree)
			} else {
				var btree Tree
				btree.Id = sn["Id"].(int)
				btree.ParentId = sn["ParentId"].(int)
				btree.Action = sn["Action"].(string)
				btree.Controller = sn["Controller"].(string)
				btree.Name = sn["Name"].(string)
				btree.Children = nil
				for i, s := range snTree {
					if btree.ParentId == s.Id {
						snTree[i].Children = append(s.Children, btree)
						break
					}
				}
			}
		}
		this.Data["json"] = snTree
	}
	fmt.Println(snTree)
	fmt.Println(fetchTree(snTree))
	return fetchTree(snTree)
}

//把navigation拼接成html
func fetchTree(children []Tree) string {
	html := ""
	var navHtml, childrenHtml string
	for _, nav := range children {
		childrenHtml = ""
		if nav.Children != nil {
			childrenHtml = fetchTree(nav.Children)
			if len(childrenHtml) > 0 {
				childrenHtml = "<ul>" + childrenHtml + "</ul>"
			}
		}
		navHtml = ""
		navHtml = "<li><span"
		if len(nav.Controller) > 0 && len(nav.Action) > 0 {
			//权限管理
			//	            if(!checkPower($nav["module"], $nav["action"])){
			//	                continue;
			//	            }
			navHtml = navHtml + " href='/" + nav.Controller + "/" + nav.Action + "'"
		}
		navHtml = navHtml + ">" + nav.Name + "</span>" + childrenHtml + "</li>"
		html = html + navHtml
	}
	return html
}
