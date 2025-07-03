package todo

import (
	"bubble/internal/global"
	"bubble/internal/model"
	"errors"
	"gorm.io/gorm"
)

func GetAllTodo(userID uint) (todoList []*model.Todo, err error) {

	if err = global.DB.Where("user_id = ?", userID).Find(&todoList).Error; err != nil {
		return nil, err
	}
	return

}

// CreateATodo 创建todo
func CreateATodo(userID uint, todo *model.Todo) (err error) {
	todo.UserID = userID
	err = global.DB.Create(&todo).Error
	return
}

func UpdateTodo(id string, input *model.UpdateTodoInput) (model.Todo, error) {
	var todo model.Todo
	if err := global.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todo, errors.New("记录未找到")
		}
		// 其他查询错误（连接、权限等）
		return todo, err
	}

	// 使用结构体 Updates：只更新非 nil 字段
	if err := global.DB.Model(&todo).Updates(input).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func DeleteATodo(id string) (err error) {
	err = global.DB.Where("id = ?", id).Delete(&model.Todo{}).Error
	return
}
