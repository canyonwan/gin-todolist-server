package controller

import (
	"canyonwan.top/gin_todolist_server/dao"
	"canyonwan.top/gin_todolist_server/model"
	"canyonwan.top/gin_todolist_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var todo model.TodoItem

func CreateItem(c *gin.Context) {
	// 声明模型，用于接收数据
	// 获取前端传过来的参数
	err := c.BindJSON(&todo)
	if err != nil {
		return
	}
	daoErr := dao.CreateATodo(&todo)
	if daoErr != nil {
		c.JSON(http.StatusOK, util.CommonResponse(http.StatusInternalServerError, daoErr.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, util.CommonResponse(http.StatusOK, "创建成功", todo))
}

func UpdateItem(c *gin.Context) {
	err := c.BindJSON(&todo)
	if err != nil {
		return
	}
	updateErr := dao.UpdateATodo(&todo)
	if updateErr != nil {
		return
	}
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	err := dao.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, util.CommonResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, util.CommonResponse(http.StatusOK, "删除成功", todo))
}
