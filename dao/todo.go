package dao

import (
	"canyonwan.top/gin_todolist_server/model"
	"canyonwan.top/gin_todolist_server/setting"
)

// CreateATodo 新增一条待办
func CreateATodo(item *model.TodoItem) (err error) {
	err = setting.DB.Create(&item).Error
	return
}

func UpdateATodo(item *model.TodoItem) (err error) {
	err = setting.DB.Save(item).Error
	return
}

// DeleteATodo 删除一条
func DeleteATodo(id string) (err error) {
	err = setting.DB.Where("id=?", id).Delete(&model.TodoItem{}).Error
	return
}

// 获取一条

// 获取全部
