package models

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type TodoResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Todo   `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

///*
//	Todo这个Model的增删改查操作都放在这里
//*/
//// CreateATodo 创建todo
//func CreateATodo(todo *Todo) (err error) {
//	err = db.DB.Create(&todo).Error
//	return
//}
//
//func GetAllTodo() (todoList []*Todo, err error) {
//
//	if err = db.DB.Find(&todoList).Error; err != nil {
//		return nil, err
//	}
//	return
//
//}
//
//func GetATodo(id string) (todo *Todo, err error) {
//	todo = new(Todo)
//	if err = db.DB.Where("id=?", id).First(todo).Error; err != nil {
//		return nil, err
//	}
//	return
//}
//
//func UpdateATodo(todo *Todo) (err error) {
//	err = db.DB.Save(todo).Error
//	return
//}
//
//func DeleteATodo(id string) (err error) {
//	err = db.DB.Where("id=?", id).Delete(&Todo{}).Error
//	return
//}
