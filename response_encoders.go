package kbp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (t *TaskCreateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.TaskId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *TaskGetResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.TaskId); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, t.TaskStatus); err != nil {
		return nil, err
	}
	if err := encodeString(
		buf,
		stringEncodePair{t.TaskName, Byte},
		stringEncodePair{t.Description, Uint64},
		stringEncodePair{t.ProjectName, Byte},
		stringEncodePair{t.DueDate, Byte},
	); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *TaskListResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	t.length = uint64(len(t.Tasks))
	if err := binary.Write(buf, binary.BigEndian, t.length); err != nil {
		return nil, err
	}
	for _, task := range t.Tasks {
		taskBytes, err := task.Encode()
		if err != nil {
			return nil, fmt.Errorf("error encoding one of the tasks : %s", err)
		}
		if err := binary.Write(buf, binary.BigEndian, uint64(len(taskBytes))); err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, taskBytes); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (t *TaskUpdateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.TaskId); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, t.TaskStatus); err != nil {
		return nil, err
	}
	if err := encodeString(
		buf,
		stringEncodePair{t.TaskName, Byte},
		stringEncodePair{t.Description, Uint64},
		stringEncodePair{t.ProjectName, Byte},
		stringEncodePair{t.DueDate, Byte},
	); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *TaskDeleteResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.TaskId); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, t.TaskStatus); err != nil {
		return nil, err
	}
	if err := encodeString(
		buf,
		stringEncodePair{t.TaskName, Byte},
		stringEncodePair{t.Description, Uint64},
		stringEncodePair{t.ProjectName, Byte},
		stringEncodePair{t.DueDate, Byte},
	); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *ProjectCreateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, p.ProjectId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *ProjectGetResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, p.ProjectId); err != nil {
		return nil, err
	}
	if err := encodeString(buf,
		stringEncodePair{p.ProjectName, Byte},
		stringEncodePair{p.Description, Uint64},
	); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *ProjectListResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	p.length = uint64(len(p.Projects))
	if err := binary.Write(buf, binary.BigEndian, p.length); err != nil {
		return nil, err
	}
	for _, project := range p.Projects {
		projectBytes, err := project.Encode()
		if err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, uint64(len(projectBytes))); err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, projectBytes); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (p *ProjectUpdateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, p.ProjectId); err != nil {
		return nil, err
	}
	if err := encodeString(
		buf,
		stringEncodePair{p.ProjectName, Byte},
		stringEncodePair{p.Description, Uint64},
	); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *ProjectDeleteResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, p.ProjectId); err != nil {
		return nil, err
	}
	if err := encodeString(
		buf,
		stringEncodePair{p.ProjectName, Byte},
		stringEncodePair{p.Description, Uint64},
	); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (n *NoteCreateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, n.NoteId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (n *NoteGetResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, n.NoteId); err != nil {
		return nil, err
	}
	if err := encodeString(
		buf,
		stringEncodePair{n.NoteDetail, Uint64},
	); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (n *NoteListResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	n.length = uint64(len(n.Notes))
	if err := binary.Write(buf, binary.BigEndian, n.length); err != nil {
		return nil, err
	}
	for _, note := range n.Notes {
		noteBytes, err := note.Encode()
		if err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, uint64(len(noteBytes))); err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, noteBytes); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (n *NoteUpdateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, n.NoteId); err != nil {
		return nil, err
	}
	if err := encodeString(buf, stringEncodePair{n.NoteDetail, Uint64}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (n *NoteDeleteResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, n.NoteId); err != nil {
		return nil, err
	}
	if err := encodeString(buf, stringEncodePair{n.NoteDetail, Uint64}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (r *ReminderCreateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, r.ReminderId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (r *ReminderGetResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, r.ReminderId); err != nil {
		return nil, err
	}
	if err := encodeString(buf, stringEncodePair{r.DueTime, Byte}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (r *ReminderListResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	r.length = uint64(len(r.Reminders))
	if err := binary.Write(buf, binary.BigEndian, r.length); err != nil {
		return nil, err
	}
	for _, reminder := range r.Reminders {
		reminderBytes, err := reminder.Encode()
		if err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, uint64(len(reminderBytes))); err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, reminderBytes); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (r *ReminderUpdateResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, r.ReminderId); err != nil {
		return nil, err
	}
	if err := encodeString(buf, stringEncodePair{r.DueTime, Byte}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (r *ReminderDeleteResponse) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, r.ReminderId); err != nil {
		return nil, err
	}
	if err := encodeString(buf, stringEncodePair{r.DueTime, Byte}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
