package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/prometheus/common/log"
	"superStar/models"
	"superStar/services"
	"time"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	datalist := c.Service.GetAll()
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	}
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后天",
			"info":  data,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	if err != nil {
		log.Fatal(err)
	}
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info)
	} else {
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}
