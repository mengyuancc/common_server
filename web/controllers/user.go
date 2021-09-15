/**
 * 用户相关
 * http://localhost:8080/user
 */
package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"iris_server/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

// 获取用户
func (c *UserController) Get() {
	datalist := c.Service.GetAll()
	fmt.Println(datalist)
	data :=  map[string]interface{}{
		"code":0,
		"data":datalist,
	}
	c.Ctx.JSONP(data)
}

// 根据id获取用户
func (c *UserController) GetBy(id int) {
	if id < 1 { 
		panic("id 不存在")
	}
	data := c.Service.Get(id)
	c.Ctx.JSON(data)
}

// 根据名称获取用户
func (c *UserController) GetSearch() {
	country := c.Ctx.URLParam("name")
	datalist := c.Service.Search(country)
	// set the model and render the view template.
	c.Ctx.JSON(datalist)
}