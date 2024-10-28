package server

type BlogDetailJSON struct {
	TaskId  string `json:"taskid" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
