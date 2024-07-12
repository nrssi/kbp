package kbp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func decodeAndComp(p1 KbpPayload) error {
	var err error
	var encodedBytes []byte
	if encodedBytes, err = p1.Encode(); err != nil {
		return err
	}
	var decodedStruct KbpPayload
	switch p1.(type) {
	case *TaskCreateRequest:
		decodedStruct = &TaskCreateRequest{}
	case *TaskGetRequest:
		decodedStruct = &TaskGetRequest{}
	case *TaskListRequest:
		decodedStruct = &TaskListRequest{}
	case *TaskUpdateRequest:
		decodedStruct = &TaskUpdateRequest{}
	case *TaskDeleteRequest:
		decodedStruct = &TaskDeleteRequest{}
	case *ProjectCreateRequest:
		decodedStruct = &ProjectCreateRequest{}
	case *ProjectGetRequest:
		decodedStruct = &ProjectGetRequest{}
	case *ProjectListRequest:
		decodedStruct = &ProjectListRequest{}
	case *ProjectUpdateRequest:
		decodedStruct = &ProjectUpdateRequest{}
	case *ProjectDeleteRequest:
		decodedStruct = &ProjectDeleteRequest{}
	case *NoteCreateRequest:
		decodedStruct = &NoteCreateRequest{}
	case *NoteListRequest:
		decodedStruct = &NoteListRequest{}
	case *NoteUpdateRequest:
		decodedStruct = &NoteUpdateRequest{}
	case *NoteDeleteRequest:
		decodedStruct = &NoteDeleteRequest{}
	case *ReminderCreateRequest:
		decodedStruct = &ReminderCreateRequest{}
	case *ReminderGetRequest:
		decodedStruct = &ReminderGetRequest{}
	case *ReminderListRequest:
		decodedStruct = &ReminderListRequest{}
	case *ReminderUpdateRequest:
		decodedStruct = &ReminderUpdateRequest{}
	case *ReminderDeleteRequest:
		decodedStruct = &ReminderDeleteRequest{}
	case *TaskCreateResponse:
		decodedStruct = &TaskCreateResponse{}
	case *TaskGetResponse:
		decodedStruct = &TaskGetResponse{}
	case *TaskListResponse:
		decodedStruct = &TaskListResponse{}
	case *TaskUpdateResponse:
		decodedStruct = &TaskUpdateResponse{}
	case *TaskDeleteResponse:
		decodedStruct = &TaskDeleteResponse{}
	case *ProjectCreateResponse:
		decodedStruct = &ProjectCreateResponse{}
	case *ProjectGetResponse:
		decodedStruct = &ProjectGetResponse{}
	case *ProjectListResponse:
		decodedStruct = &ProjectListResponse{}
	case *ProjectUpdateResponse:
		decodedStruct = &ProjectUpdateResponse{}
	case *ProjectDeleteResponse:
		decodedStruct = &ProjectDeleteResponse{}
	case *NoteCreateResponse:
		decodedStruct = &NoteCreateResponse{}
	case *NoteGetResponse:
		decodedStruct = &NoteGetResponse{}
	case *NoteListResponse:
		decodedStruct = &NoteListResponse{}
	case *NoteUpdateResponse:
		decodedStruct = &NoteUpdateResponse{}
	case *NoteDeleteResponse:
		decodedStruct = &NoteDeleteResponse{}
	case *ReminderCreateResponse:
		decodedStruct = &ReminderCreateResponse{}
	case *ReminderGetResponse:
		decodedStruct = &ReminderGetResponse{}
	case *ReminderListResponse:
		decodedStruct = &ReminderListResponse{}
	case *ReminderUpdateResponse:
		decodedStruct = &ReminderUpdateResponse{}
	case *ReminderDeleteResponse:
		decodedStruct = &ReminderDeleteResponse{}
	default:
		panic("Unknow type, check the implementation")
	}
	if err := decodedStruct.Decode(encodedBytes); err != nil {
		return err
	}
	if !reflect.DeepEqual(p1, decodedStruct) {
		return fmt.Errorf("expected p1 : \n%s\n\ngot decodedStruct : \n%s \n\n", spew.Sdump(p1), spew.Sdump(decodedStruct))
	}
	return nil
}

func TestTaskCreateRequest(t *testing.T) {
	tRequest := TaskCreateRequest{
		"Sample Task Name",
		"some very important description about the task",
		"Dream project",
		"23-06-2024",
		Done,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskGetRequest(t *testing.T) {
	tRequest := TaskGetRequest{
		235,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskListRequest(t *testing.T) {
	tRequest := TaskListRequest{
		TaskName:    "some task name",
		Description: "some description",
		ProjectName: "important project",
		DueDate:     "23-12-2024",
		Status:      Todo,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskUpdateRequest(t *testing.T) {
	tRequest := TaskUpdateRequest{
		TaskName:    "some of task",
		Description: "some description",
		ProjectName: "some project",
		DueDate:     "some due date",
		TaskId:      12334,
		Status:      Done,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestTaskDeleteRequest(t *testing.T) {
	tRequest := TaskDeleteRequest{
		TaskId: 123,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectCreateRequest(t *testing.T) {
	tRequest := ProjectCreateRequest{
		ProjectName: "project name",
		Description: "some description",
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectGetRequest(t *testing.T) {
	tRequest := ProjectGetRequest{
		ProjectId: 234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectListRequest(t *testing.T) {
	tRequest := ProjectListRequest{
		ProjectName: "sdfsdfsdffsdfsdf",
		Description: "sdfsdfglkkajsdfaljs dlfkasjslfgkjaslkfjadffgjlkadssjflaks alsdkfjaslkfjalskfjaskljfaklsfjaslk jalskfj lsadkjfadklsfjaklsdjf ladfks",
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectUpdateRequest(t *testing.T) {
	tRequest := ProjectUpdateRequest{
		ProjectName: "sdflaksdfjalsdkfjaslKFJASDLK JLAKSDFNLAS",
		Description: "sdflajsfljasdlf alsdfjlksdf jlksdjfklsflksdflksdf klsddfj slkjsdrflkjsdf lkdjflkdfslkjdslkdflkjdfs  kse mdflkj sdfjkldfskljdfskljsdfjkl",
		ProjectId:   234234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestProjectDeleteRequest(t *testing.T) {
	tRequest := ProjectDeleteRequest{
		ProjectId: 234,
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteCreateRequest(t *testing.T) {
	tRequest := NoteCreateRequest{
		NoteDetail: "sdflkkjasdflkajsdflk jadslfgkjasdlkfjasldkfjasldkfjaklsdfj alsjasdlkfjafjaskfklasdjfkladjsf kjasdl FJALSKDFJLKASDJF LKADDJFGLK AJSLKAJSDF KLAJSDFKLJASDL KJASLDKFJALSKSalkskfjaklsfj aljflkasj fklasjf klasjklasjf alsk",
		TaskId:     234,
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteListRequest(t *testing.T) {
	tRequest := NoteListRequest{
		TaskId: 234,
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestNoteDeleteRequest(t *testing.T) {
	tRequest := NoteDeleteRequest{
		NoteId: 234,
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderCreateRequest(t *testing.T) {
	tRequest := ReminderCreateRequest{
		Time: "23-12-2001",
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderGetRequest(t *testing.T) {
	tRequest := ReminderGetRequest{
		ReminderId: 234,
	}
	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderListRequest(t *testing.T) {
	tRequest := ReminderListRequest{
		DueTime: "12-12-2004",
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderUpdateRequest(t *testing.T) {
	tRequest := ReminderUpdateRequest{
		DueTime:    "12-12-2004",
		ReminderId: 2134,
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}

func TestReminderDeleteRequest(t *testing.T) {
	tRequest := ReminderDeleteRequest{
		ReminderId: 2134,
	}

	if err := decodeAndComp(&tRequest); err != nil {
		t.Fatal(err)
	}
}
