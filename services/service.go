package services

import (
	"bubble/db"
	"bubble/models"
)

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

func DeleteATodo(id string) (err error) {
	err = db.DB.Where("id=?", id).Delete(&models.Todo{}).Error
	return
}
