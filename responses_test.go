package kbp

import "testing"

func TestTaskCreateRequestResponse(t *testing.T) {
	tRequest := TaskCreateResponse{
		TaskId: 123,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskGetResponse(t *testing.T) {
	tRequest := TaskGetResponse{
		TaskName:    "some name",
		Description: "some description, some details about the task",
		ProjectName: "some project name",
		DueDate:     "2023-12-14",
		TaskId:      123,
		TaskStatus:  Done,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskListResponse(t *testing.T) {
	tRequest := TaskListResponse{
		Tasks: []TaskGetResponse{
			{
				TaskName:    "some name 1",
				Description: "some description, some details about the task",
				ProjectName: "some project name",
				DueDate:     "2023-12-14",
				TaskId:      123,
				TaskStatus:  Done,
			},
			{
				TaskName:    "some name 2",
				Description: "some description, some details about the task",
				ProjectName: "some project name",
				DueDate:     "2023-12-14",
				TaskId:      123,
				TaskStatus:  Done,
			},
			{
				TaskName:    "some name 3",
				Description: "some description, some details about the task",
				ProjectName: "some project name",
				DueDate:     "2023-12-14",
				TaskId:      123,
				TaskStatus:  Done,
			},
			{
				TaskName:    "some name 4",
				Description: "some description, some details about the task",
				ProjectName: "some project name",
				DueDate:     "2023-12-14",
				TaskId:      123,
				TaskStatus:  Done,
			},
		},
		length: 4,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskUpdateResponse(t *testing.T) {
	tRequest := TaskUpdateResponse{
		TaskName:    "some name",
		Description: "some description, some details about the task",
		ProjectName: "some project name",
		DueDate:     "2023-12-14",
		TaskId:      123,
		TaskStatus:  Done,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskDeleteResponse(t *testing.T) {
	tRequest := TaskUpdateResponse{
		TaskName:    "some name",
		Description: "some description, some details about the task",
		ProjectName: "some project name",
		DueDate:     "2023-12-14",
		TaskId:      123,
		TaskStatus:  Done,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectCreateResponse(t *testing.T) {
	tRequest := ProjectCreateResponse{
		ProjectId: 123,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectGetResponse(t *testing.T) {
	tRequest := ProjectGetResponse{
		ProjectName: "some project name",
		Description: "some project description",
		ProjectId:   23,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectListResponse(t *testing.T) {
	tRequest := ProjectListResponse{
		Projects: []ProjectGetResponse{
			{
				ProjectName: "project 1",
				Description: "some project description",
				ProjectId:   23,
			},
			{
				ProjectName: "project 2",
				Description: "some project description",
				ProjectId:   23,
			},
			{
				ProjectName: "project 3",
				Description: "some project description",
				ProjectId:   23,
			},
			{
				ProjectName: "project 4",
				Description: "some project description",
				ProjectId:   23,
			},
		},
		length: 4,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectUpdateResponse(t *testing.T) {
	tRequest := ProjectUpdateResponse{
		ProjectName: "some project name",
		Description: "some project description",
		ProjectId:   23,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectDeleteResponse(t *testing.T) {
	tRequest := ProjectGetResponse{
		ProjectName: "some project name",
		Description: "some project description",
		ProjectId:   23,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteCreateResponse(t *testing.T) {
	tRequest := NoteCreateResponse{
		NoteId: 234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteGetResponse(t *testing.T) {
	tRequest := NoteGetResponse{
		NoteDetail: "some note detail",
		NoteId:     234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteListResponse(t *testing.T) {
	tRequest := NoteListResponse{
		Notes: []NoteGetResponse{
			{
				NoteDetail: "some detail for note 1",
				NoteId:     123,
			},
			{
				NoteDetail: "some detail for note 2",
				NoteId:     123,
			},
			{
				NoteId:     123,
				NoteDetail: "some detail for note 3",
			},
			{
				NoteDetail: "some detail for note 4",
				NoteId:     123,
			},
		},
		length: 0,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteUpdateResponse(t *testing.T) {
	tRequest := NoteUpdateResponse{
		NoteDetail: "some note detail",
		NoteId:     234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteDeleteResponse(t *testing.T) {
	tRequest := NoteDeleteResponse{
		NoteDetail: "some note detail",
		NoteId:     234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderCreateResponse(t *testing.T) {
	tRequest := ReminderCreateResponse{
		ReminderId: 234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderGetResponse(t *testing.T) {
	tRequest := ReminderGetResponse{
		DueTime:    "23-12-2024",
		ReminderId: 234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderListResponse(t *testing.T) {
	tRequest := ReminderListResponse{
		Reminders: []ReminderGetResponse{
			{
				DueTime:    "23-12-2024",
				ReminderId: 12,
			},
			{
				DueTime:    "23-12-2024",
				ReminderId: 12,
			},
			{
				DueTime:    "23-12-2024",
				ReminderId: 12,
			},
			{
				DueTime:    "23-12-2024",
				ReminderId: 12,
			},
		},
		length: 0,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderUpdateResponse(t *testing.T) {
	tRequest := ReminderUpdateResponse{
		DueTime:    "23-12-2024",
		ReminderId: 234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderDeleteResponse(t *testing.T) {
	tRequest := ReminderDeleteResponse{
		DueTime:    "23-12-2024",
		ReminderId: 234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}
