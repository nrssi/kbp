package kbp

type KbpRequest[T KbpPayload] struct {
	Header KbpRequestHeader
	Body   T
}

type TaskCreateRequest struct {
	taskNameLen    byte
	TaskName       string
	descriptionLen uint64
	Description    string
	projectNameLen byte
	ProjectName    string
	dueDateLen     byte
	DueDate        string
	taskStatusLen  byte
	TaskStatus     TaskStatus
}

type TaskGetRequest struct {
	TaskId uint64
}

type TaskListRequest struct {
	taskNameLen    byte
	TaskName       string
	descriptionLen uint64
	Description    string
	projectNameLen byte
	ProjectName    string
}

type TaskUpdateRequest struct {
	TaskId         uint64
	taskNameLen    byte
	TaskName       string
	descriptionLen uint64
	Description    string
	projectNameLen byte
	ProjectName    string
}

type TaskDeleteRequest struct {
	TaskId uint64
}

type ProjectCreatRequest struct {
	projectNameLen        byte
	ProjectName           string
	projectDescriptionLen uint64
	Description           string
}

type ProjectGetRequest struct {
	ProjectId uint64
}

type ProjectListRequest struct {
	projectNameLen        byte
	ProjectName           string
	projectDescriptionLen uint64
	Description           string
}

type ProjectDeleteRequest struct {
	ProjectId uint64
}

type NoteCreateRequest struct {
	TaskId        uint64
	noteDetailLen uint64
	NoteDetail    string
}

type NoteListRequest struct {
	TaskId uint64
}

type NoteUpdateRequest struct {
	NoteId        uint64
	noteDetailLen uint64
	NoteDetail    string
}

type NoteDeleteRequest struct {
	NoteId uint64
}

type CreateReminderRequest struct {
	time string
}

type GetReminderRequest struct {
	ReminderId uint64
}

type ListReminderRequest struct {
	DueTime string
}

type DeleteReminderRequest struct {
	ReminderId uint64
}
