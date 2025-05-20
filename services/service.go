package services

import (
	"bubble/db"
	"bubble/models"
	"errors"
	"gorm.io/gorm"
)

func CreateAUser(user *models.User) (err error) {
	err = db.DB.Create(&user).Error
	return
}

// CreateATodo 创建todo
func CreateATodo(todo *models.Todo) (err error) {
	err = db.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*models.Todo, err error) {

	if err = db.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return

}

func GetATodo(id string) (todo *models.Todo, err error) {
	todo = new(models.Todo)
	if err = db.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *models.Todo) (err error) {
	err = db.DB.Save(todo).Error
	return
}

func UpdateTodo(id string, input *models.UpdateTodoInput) (models.Todo, error) {
	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todo, errors.New("纪录未找到")
		}
		// 其他查询错误（连接、权限等）
		return todo, err
	}

	// 使用结构体 Updates：只更新非 nil 字段
	if err := db.DB.Model(&todo).Updates(input).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func DeleteATodo(id string) (err error) {
	err = db.DB.Where("id=?", id).Delete(&models.Todo{}).Error
	return
}
