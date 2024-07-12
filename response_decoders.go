package kbp

import (
	"bytes"
	"encoding/binary"
)

func (t *TaskCreateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &t.TaskId); err != nil {
		return err
	}
	return nil
}

func (t *TaskGetResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &t.TaskId); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &t.TaskStatus); err != nil {
		return err
	}
	if t.TaskName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if t.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	if t.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if t.DueDate, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (t *TaskListResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &t.length); err != nil {
		return err
	}
	t.Tasks = make([]TaskGetResponse, 0)
	for i := 0; i < int(t.length); i++ {
		var taskLen uint64
		if err := binary.Read(buf, binary.BigEndian, &taskLen); err != nil {
			return err
		}
		taskBytes := make([]byte, taskLen)
		if err := binary.Read(buf, binary.BigEndian, &taskBytes); err != nil {
			return err
		}
		var task TaskGetResponse
		if err := task.Decode(taskBytes); err != nil {
			return err
		}
		t.Tasks = append(t.Tasks, task)
	}
	return nil
}

func (t *TaskUpdateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &t.TaskId); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &t.TaskStatus); err != nil {
		return err
	}
	if t.TaskName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if t.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	if t.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if t.DueDate, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (t *TaskDeleteResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &t.TaskId); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &t.TaskStatus); err != nil {
		return err
	}
	if t.TaskName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if t.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	if t.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if t.DueDate, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (p *ProjectCreateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &p.ProjectId); err != nil {
		return err
	}
	return nil
}

func (p *ProjectGetResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &p.ProjectId); err != nil {
		return err
	}
	if p.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if p.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (p *ProjectListResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)

	if err := binary.Read(buf, binary.BigEndian, &p.length); err != nil {
		return err
	}
	p.Projects = make([]ProjectGetResponse, 0)
	for i := 0; i < int(p.length); i++ {
		var projectLen uint64
		if err := binary.Read(buf, binary.BigEndian, &projectLen); err != nil {
			return err
		}
		projectBytes := make([]byte, projectLen)
		if err := binary.Read(buf, binary.BigEndian, &projectBytes); err != nil {
			return err
		}
		var project ProjectGetResponse
		if err := project.Decode(projectBytes); err != nil {
			return err
		}
		p.Projects = append(p.Projects, project)
	}
	return nil
}

func (p *ProjectUpdateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &p.ProjectId); err != nil {
		return err
	}
	if p.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if p.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (p *ProjectDeleteResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &p.ProjectId); err != nil {
		return err
	}
	if p.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if p.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (n *NoteCreateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &n.NoteId); err != nil {
		return err
	}
	return nil
}

func (n *NoteGetResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &n.NoteId); err != nil {
		return err
	}
	if n.NoteDetail, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (n *NoteListResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &n.length); err != nil {
		return err
	}
	n.Notes = make([]NoteGetResponse, 0)
	for i := 0; i < int(n.length); i++ {
		var noteLen uint64
		if err := binary.Read(buf, binary.BigEndian, &noteLen); err != nil {
			return err
		}
		noteBytes := make([]byte, noteLen)
		if err := binary.Read(buf, binary.BigEndian, &noteBytes); err != nil {
			return err
		}
		var note NoteGetResponse
		if err := note.Decode(noteBytes); err != nil {
			return err
		}
		n.Notes = append(n.Notes, note)
	}
	return nil
}

func (n *NoteUpdateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &n.NoteId); err != nil {
		return err
	}
	if n.NoteDetail, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (n *NoteDeleteResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &n.NoteId); err != nil {
		return err
	}
	if n.NoteDetail, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (r *ReminderCreateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &r.ReminderId); err != nil {
		return err
	}
	return nil
}

func (r *ReminderGetResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &r.ReminderId); err != nil {
		return err
	}
	if r.DueTime, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (r *ReminderListResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &r.length); err != nil {
		return err
	}
	r.Reminders = make([]ReminderGetResponse, 0)
	for i := 0; i < int(r.length); i++ {
		var reminderLen uint64
		if err := binary.Read(buf, binary.BigEndian, &reminderLen); err != nil {
			return err
		}
		reminderBytes := make([]byte, reminderLen)
		if err := binary.Read(buf, binary.BigEndian, &reminderBytes); err != nil {
			return err
		}
		var reminder ReminderGetResponse
		if err := reminder.Decode(reminderBytes); err != nil {
			return err
		}
		r.Reminders = append(r.Reminders, reminder)
	}
	return nil
}

func (r *ReminderUpdateResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &r.ReminderId); err != nil {
		return err
	}
	if r.DueTime, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (r *ReminderDeleteResponse) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err := binary.Read(buf, binary.BigEndian, &r.ReminderId); err != nil {
		return err
	}
	if r.DueTime, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}
