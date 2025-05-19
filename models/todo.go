package models

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status" gorm:"default:0"`
}

type TodoResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Todo   `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
