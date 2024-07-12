package kbp

type KbpResponse struct {
	Body   *KbpPayload
	Header KbpResponseHeader
}

type TaskCreateResponse struct {
	TaskId uint64
}

type TaskGetResponse struct {
	TaskName    string
	Description string
	ProjectName string
	DueDate     string
	TaskId      uint64
	TaskStatus  TaskStatus
}

type TaskListResponse struct {
	Tasks  []TaskGetResponse
	length uint64
}

type TaskUpdateResponse struct {
	TaskName    string
	Description string
	ProjectName string
	DueDate     string
	TaskId      uint64
	TaskStatus  TaskStatus
}

type TaskDeleteResponse struct {
	TaskName    string
	Description string
	ProjectName string
	DueDate     string
	TaskId      uint64
	TaskStatus  TaskStatus
}

type ProjectCreateResponse struct {
	ProjectId uint64
}

type ProjectGetResponse struct {
	ProjectName string
	Description string
	ProjectId   uint64
}

type ProjectListResponse struct {
	Projects []ProjectGetResponse
	length   uint64
}

type ProjectUpdateResponse struct {
	ProjectName string
	Description string
	ProjectId   uint64
}

type ProjectDeleteResponse struct {
	ProjectName string
	Description string
	ProjectId   uint64
}

type NoteCreateResponse struct {
	NoteId uint64
}

type NoteGetResponse struct {
	NoteDetail string
	NoteId     uint64
}

type NoteListResponse struct {
	Notes  []NoteGetResponse
	length uint64
}

type NoteUpdateResponse struct {
	NoteDetail string
	NoteId     uint64
}

type NoteDeleteResponse struct {
	NoteDetail string
	NoteId     uint64
}

type ReminderCreateResponse struct {
	ReminderId uint64
}

type ReminderGetResponse struct {
	DueTime    string
	ReminderId uint64
}

type ReminderListResponse struct {
	Reminders []ReminderGetResponse
	length    uint64
}

type ReminderUpdateResponse struct {
	DueTime    string
	ReminderId uint64
}

type ReminderDeleteResponse struct {
	DueTime    string
	ReminderId uint64
}
