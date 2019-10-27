package controllers

import (
	"encoding/json"
	"meal/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"meal/utils"
	"os"
	"github.com/hunterhug/go_image"
	"strings"
	"time"
	"meal/enums"
	"strconv"
	"fmt"
	"github.com/astaxie/beego/orm"
)

//CarouselController 菜单管理
type CarouselController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *CarouselController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid",  "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 菜谱管理首页
func (c *CarouselController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//获取菜单类别
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "carousel/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "carousel/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("CarouselController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("CarouselController", "Delete")
	beego.Info("CarouselController,",c.Data["canEdit"],c.Data["canEdit"])
}

// 获取所有菜谱
func (c *CarouselController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.MealCarouselQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.MealCarousePageList(&params)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}

//Edit 添加、编辑菜谱界面
func (c *CarouselController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt64(":id", 0)
	m := models.Carouse{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["hasImg"] = len(m.Img) > 0
	c.Data["m"] = m

	c.setTpl("carousel/edit.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "carousel/edit_headcssjs.html"
	c.LayoutSections["footerjs"] = "carousel/edit_footerjs.html"

	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("CarouselController.Index")
}

//Save 添加、编辑页面 保存
func (c *CarouselController) Save() {
	var err error
	m := models.Carouse{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "提交表单数据失败，可能原因："+err.Error(), m.Id)
	}
	if m.Img == "" {
		c.jsonResult(enums.JRCodeFailed, "图片地址不能为空"+err.Error(), m.Id)
	}
	logs.Info("Carouse",m)
	o := orm.NewOrm()
	if m.Id == 0 {

		m.Time = time.Now().Unix()
		if _, err = o.Insert(&m); err == nil {
			c.jsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因："+err.Error(), m.Id)
		}

	} else {
		oM, err := models.CarouseOne(m.Id)
		oM.Img = m.Img
		oM.Time = time.Now().Unix()
		if _, err = o.Update(oM); err == nil {
			c.jsonResult(enums.JRCodeSucc, "更新成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "更新失败", m.Id)
		}
	}

}

//Delete 批量删除
func (c *CarouselController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.MealCarouseBatchDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *CarouselController) UpdateSeq() {
	Id, _ := c.GetInt64("pk", 0)
	oM, err := models.CarouseOne(Id)
	if err != nil || oM == nil {
		c.jsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}
	value, _ := c.GetInt32("value", 0)
	oM.Seq = value
	o := orm.NewOrm()
	if _, err := o.Update(oM); err == nil {
		c.jsonResult(enums.JRCodeSucc, "修改成功", oM.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, "修改失败", oM.Id)
	}
}

func (c *CarouselController) UploadImage() {
	//这里type没有用，只是为了演示传值
	beego.Info("body",c.Input().Get("Img"))
	stype, _ := c.GetInt32("type", 0)
	if stype > 0 {
		f, h, err := c.GetFile("fileImg")
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "上传失败", "")
		}
		defer f.Close()
		filepreix := "static/upload/"
		exist := utils.IsExist(filepreix)
		if exist {
			logs.Info("filepreix",filepreix)
		} else {
			logs.Error("file path not exist",filepreix)
			// 创建文件夹
			err := os.Mkdir(filepreix, os.ModePerm)
			if err != nil {
				logs.Error("create file err",err)
			} else {
				logs.Error("mkdir success!\n")
			}
		}
		filePath := filepreix + h.Filename
		c.SaveToFile("fileImg", filePath)
		realfilename, err := go_image.RealImageName(filePath)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "上传失败"+err.Error(), "")
		}
		sp := strings.Split(realfilename,".")
		if len(sp) != 2 {
			c.jsonResult(enums.JRCodeFailed, "上传失败", "")
		}
		newfile := filepreix+time.Unix(utils.GetNow(),0).Format("2006-01-02")+utils.RandomString(10)+sp[1]
		// 保存位置在 static/upload, 没有文件夹要先创建
		err = go_image.ScaleF2F(filePath, newfile,600 )
		if err != nil {
			panic(err)
		}
		logs.Info("oldfile",filePath)
		logs.Info("newfile",newfile)
		c.jsonResult(enums.JRCodeSucc, "上传成功", "/"+newfile)
	} else {
		c.jsonResult(enums.JRCodeFailed, "上传失败", "")
	}
}