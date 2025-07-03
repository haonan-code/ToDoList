package model

// Todo Model
type Todo struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title  string `json:"title" gorm:"not null"`
	// 0：未完成 1:已完成
	Status int `json:"status" gorm:"default:0"`
	// 0：优先级高 1：优先级中 2：优先级低
	Priority int `json:"priority" gorm:"default:2"`
}

// 用于更新（可选字段）
type UpdateTodoInput struct {
	Title    *string `json:"title,omitempty"`
	Status   *int    `json:"status,omitempty"`
	Priority *int    `json:"priority,omitempty"`
}

//type TodoResponse struct {
//	Status int    `json:"status"`
//	Msg    string `json:"msg"`
//	Data   Todo   `json:"data"`
//}
//
//type ErrorResponse struct {
//	Error string `json:"error"`
//}
