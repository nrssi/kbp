package kbp

type KbpRequest[T KbpPayload] struct {
	Body   T
	Header KbpRequestHeader
}

type TaskCreateRequest struct {
	TaskName    string
	Description string
	ProjectName string
	DueDate     string
	Status      TaskStatus
}

type TaskGetRequest struct {
	TaskId uint64
}

type TaskListRequest struct {
	TaskName    string
	Description string
	ProjectName string
	DueDate     string
	Status      TaskStatus
}

type TaskUpdateRequest struct {
	TaskName    string
	Description string
	ProjectName string
	DueDate     string
	TaskId      uint64
	Status      TaskStatus
}

type TaskDeleteRequest struct {
	TaskId uint64
}

type ProjectCreateRequest struct {
	ProjectName string
	Description string
}

type ProjectGetRequest struct {
	ProjectId uint64
}

type ProjectListRequest struct {
	ProjectName string
	Description string
}

type ProjectUpdateRequest struct {
	ProjectName string
	Description string
	ProjectId   uint64
}
type ProjectDeleteRequest struct {
	ProjectId uint64
}

type NoteCreateRequest struct {
	NoteDetail string
	TaskId     uint64
}

type NoteListRequest struct {
	TaskId uint64
}

type NoteUpdateRequest struct {
	NoteDetail string
	NoteId     uint64
}

type NoteDeleteRequest struct {
	NoteId uint64
}

type ReminderCreateRequest struct {
	Time string
}

type ReminderGetRequest struct {
	ReminderId uint64
}

type ReminderListRequest struct {
	DueTime string
}

type ReminderUpdateRequest struct {
	DueTime    string
	ReminderId uint64
}
type ReminderDeleteRequest struct {
	ReminderId uint64
}
