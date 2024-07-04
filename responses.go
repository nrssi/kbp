package kbp

type KbpResponse struct {
	Header KbpResponseHeader
	Body   *KbpPayload
}

type TaskCreateResponse struct {
	TaskId uint64
}

type TaskGetResponse struct {
	TaskId         uint64
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

type TaskListResponse struct {
	Length uint64
	Tasks  []TaskGetResponse
}

type TaskUpdateResponse struct {
	TaskId         uint64
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

type TaskDeleteResponse struct {
	TaskId         uint64
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

type ProjectCreateResponse struct {
	ProjectId uint64
}

type ProjectGetResponse struct {
	ProjectId      uint64
	projectNameLen byte
	ProjectName    string
	descriptionLen uint64
	Description    string
}

type ProjectListResponse struct {
	Length   uint64
	Projects []ProjectGetResponse
}

type ProjectUpdateResponse struct {
	ProjectId      uint64
	projectNameLen byte
	ProjectName    string
	descriptionLen uint64
	Description    string
}

type ProjectDeleteResponse struct {
	ProjectId      uint64
	projectNameLen byte
	ProjectName    string
	descriptionLen uint64
	Description    string
}

type NoteCreateResponse struct {
	NoteId uint64
}

type NoteGetResponse struct {
	NoteId        uint64
	noteDetailLen uint64
	NoteDetail    string
}

type NoteListResponse struct {
	Length uint64
	Notes  []NoteGetResponse
}

type NoteUpdateResponse struct {
	NoteId        uint64
	noteDetailLen uint64
	NoteDetail    string
}

type NoteDeleteResponse struct {
	NoteId        uint64
	noteDetailLen uint64
	NoteDetail    string
}

type ReminderCreatResponse struct {
	ReminderId uint64
}

type ReminderGetResponse struct {
	ReminderId uint64
	DueTime    string
}

type ReminderListResponse struct {
	Length    uint64
	Reminders []ReminderListResponse
}

type ReminderUpdateResponse struct {
	ReminderId uint64
	DueTime    string
}

type ReminderDeleteResponse struct {
	ReminderId uint64
	DueTime    string
}
