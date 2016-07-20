package api

import (
	"github.com/vmware/harbor/models"
	"strconv"
	"net/http"
	"github.com/vmware/harbor/dao"
	"github.com/vmware/harbor/utils/log"
)

type CustomerController struct {
	BaseAPI
}

/*
	Method: Post
 	https://registry.51yixiao.com/api/customer
	param: name 客户中文名
	param: tag 客户标签
 */

func (c *CustomerController) PostCustomer() {
	name := c.GetString("name")
	if len(name) == 0 {
		c.CustomAbort(http.StatusBadRequest, "name is nil")
	}

	tag := c.GetString("tag")
	if len(tag) == 0 {
		c.CustomAbort(http.StatusBadRequest, "tag is nil")
	}

	customer := models.Customer{Name: name, Tag: tag}
	if res,err := dao.AddCustomer(customer); err == nil {
		//StatusCreated
		if res == true {
			c.CustomAbort(http.StatusCreated, "add success")
		}else{
			c.CustomAbort(http.StatusOK, "customer is exist ")
		}
		c.Data["json"] = res
	}else{
		c.CustomAbort(http.StatusInternalServerError,"Failed to insert customer to db")
	}

	c.ServeJSON()
}

/*
	Method: Get
 	https://registry.51yixiao.com/api/customer/3
	ret: 返回客户信息
 */
func (c *CustomerController) GetOneCustomer() {
	idStr := c.Ctx.Input.Param(":id")

	log.Infof("idStr: %+v",idStr )
	id, _ := strconv.Atoi(idStr)

	v, err := dao.GetCustomerById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

/*
	Method: Post
 	https://registry.51yixiao.com/api/customer/list
	param: project_name 项目名称
	ret: 返回客户列表
 */
func (c *CustomerController) GetListCustomer() {
	projectName := c.GetString("project_name")
	log.Infof("projectName: %+v",projectName )
	l, err := dao.GetProjectAllCustomer(projectName)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	log.Infof("user: %+v",l)
	c.ServeJSON()
}

/*
 Method: Put
 https://registry.51yixiao.com/api/customer/3
 param: name 招商银行
 param: tag CMM
 */
func (c *CustomerController) UpdateCustomer() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	m := models.Customer{}
	m.Id = id
	m.Name = c.GetString("name")
	m.Tag = c.GetString("tag")

	log.Infof("user: %+v",m)

	if err := dao.UpdateCustomerById(m); err == nil {
			c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

/*
Method: Delete
https://registry.51yixiao.com/api/customer/3
 */
func (c *CustomerController) DeleteCustomer() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := dao.DeleteCustomer(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
