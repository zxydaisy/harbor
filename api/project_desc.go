package api

import (
	"strconv"
	"github.com/vmware/harbor/models"
	"github.com/vmware/harbor/dao"
)

type ProjectDescController struct {
	BaseAPI
}

/*
 Method: Put
 https://registry.51yixiao.com/api/project_desc/3
 param: name
 */
func (c *ProjectDescController) UpdateProject() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	m := models.ProjectDesc{}
	m.ProjectId = id
	m.Name = c.GetString("name")

	if err := dao.UpdateProjectById(m); err == nil {
			c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err
	}

	c.ServeJSON()
}
