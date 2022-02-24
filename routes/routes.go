package routes

import (
	"canyonwan.top/gin_todolist_server/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine) *gin.Engine {
	v1 := e.Group("/v1/api/")
	{
		// 获取全部的待办列表
		v1.POST("todoList")
		// 获取某个待办信息
		v1.GET("detail/:id")
		// 新增一个待办
		v1.POST("addTodo", controller.CreateItem)
		// 删除一个待办
		v1.DELETE("deleteTodo/:id", controller.DeleteItem)
		// 更新一个行办
		v1.PUT("updateTdo")
	}

	return e
}
