package controllers

import (
	"encoding/json"
	"time"

	"meal/enums"
	"meal/models"
	"meal/utils"

	"fmt"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"github.com/hunterhug/go_image"
)

//MaintainController 维修管理
type MaintainController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *MaintainController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 角色管理首页
func (c *MaintainController) Index() {
	//将页面左边维修的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "maintain/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "maintain/index_footerjs.html"

}

// 获取所有菜谱
func (c *MaintainController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.MaintainQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.MaintainPageList(&params)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}

//Delete 批量删除
func (c *MaintainController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.MaintainBatchDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *MaintainController) UploadImage() {
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
