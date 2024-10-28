package types

type TaskType string

const (
	GenerateContentTaskType TaskType = "generating-content"
	GeneratePictureTaskType TaskType = "generating-picture"
)

type CreateTaskArgs struct {
	TaskId  string
	Title   string
	Content string
}

type CreateTaskReply struct{}

type GetTaskArgs struct {
	Msg string
}

type GetTaskReply struct {
	TaskType TaskType
}
