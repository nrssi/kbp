package kbp

import (
	"bytes"
	"encoding/binary"
)

func decodeString(buf *bytes.Reader, lenType stringLengthType) (string, error) {
	var val string
	switch lenType {
	case Byte:
		var strLen byte
		if err := binary.Read(buf, binary.BigEndian, &strLen); err != nil {
			return "", err
		}
		strBytes := make([]byte, strLen)
		if err := binary.Read(buf, binary.BigEndian, &strBytes); err != nil {
			return "", nil
		}
		val = string(strBytes)
	case Uint64:
		var strLen uint64
		if err := binary.Read(buf, binary.BigEndian, &strLen); err != nil {
			return "", err
		}
		strBytes := make([]byte, strLen)
		if err := binary.Read(buf, binary.BigEndian, &strBytes); err != nil {
			return "", nil
		}
		val = string(strBytes)
	default:
		panic("Not a valid string length type")
	}
	return val, nil
}

func (t *TaskCreateRequest) Decode(data []byte) error {
	var err error
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &t.Status); err != nil {
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

func (t *TaskGetRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &t.TaskId); err != nil {
		return err
	}
	return nil
}

func (t *TaskListRequest) Decode(data []byte) error {
	var err error
	buf := bytes.NewReader(data)
	if err = binary.Read(buf, binary.BigEndian, &t.Status); err != nil {
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

func (t *TaskUpdateRequest) Decode(data []byte) error {
	var err error
	buf := bytes.NewReader(data)
	if err = binary.Read(buf, binary.BigEndian, &t.TaskId); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, &t.Status); err != nil {
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

func (t *TaskDeleteRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &t.TaskId); err != nil {
		return err
	}
	return nil
}

func (p *ProjectCreateRequest) Decode(data []byte) error {
	var err error
	buf := bytes.NewReader(data)
	if p.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if p.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (p *ProjectGetRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &p.ProjectId); err != nil {
		return err
	}
	return nil
}

func (p *ProjectListRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if p.ProjectName, err = decodeString(buf, Byte); err != nil {
		return err
	}
	if p.Description, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (p *ProjectUpdateRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err = binary.Read(buf, binary.BigEndian, &p.ProjectId); err != nil {
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

func (p *ProjectDeleteRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &p.ProjectId); err != nil {
		return err
	}
	return nil
}

func (n *NoteCreateRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err = binary.Read(buf, binary.BigEndian, &n.TaskId); err != nil {
		return err
	}
	if n.NoteDetail, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (n *NoteListRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &n.TaskId); err != nil {
		return err
	}
	return nil
}

func (n *NoteUpdateRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err = binary.Read(buf, binary.BigEndian, &n.NoteId); err != nil {
		return err
	}
	if n.NoteDetail, err = decodeString(buf, Uint64); err != nil {
		return err
	}
	return nil
}

func (n *NoteDeleteRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &n.NoteId); err != nil {
		return err
	}
	return nil
}

func (r *ReminderCreateRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if r.Time, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (r *ReminderGetRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &r.ReminderId); err != nil {
		return err
	}
	return nil
}

func (r *ReminderListRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if r.DueTime, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (r *ReminderUpdateRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	var err error
	if err = binary.Read(buf, binary.BigEndian, &r.ReminderId); err != nil {
		return err
	}
	if r.DueTime, err = decodeString(buf, Byte); err != nil {
		return err
	}
	return nil
}

func (r *ReminderDeleteRequest) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.BigEndian, &r.ReminderId); err != nil {
		return err
	}
	return nil
}
