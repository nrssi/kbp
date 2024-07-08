package kbp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type stringLengthType byte

type stringEncodePair struct {
	val        string
	lengthType stringLengthType
}

const (
	Byte   stringLengthType = 0x01
	Uint64 stringLengthType = 0x02
)

func encodeString(buf *bytes.Buffer, vals ...stringEncodePair) error {
	for _, pair := range vals {
		switch pair.lengthType {
		case Byte:
			strLen := byte(len(pair.val))
			if err := binary.Write(buf, binary.BigEndian, strLen); err != nil {
				return fmt.Errorf("error encoding length of string %s with length %d", pair.val, len(pair.val))
			}
			if err := binary.Write(buf, binary.BigEndian, []byte(pair.val)[:strLen]); err != nil {
				return fmt.Errorf("error encoding string %s with length %d", pair.val, len(pair.val))
			}
		case Uint64:
			strLen := uint64(len(pair.val))
			if err := binary.Write(buf, binary.BigEndian, strLen); err != nil {
				return fmt.Errorf("error encoding length of string %s with length %d", pair.val, len(pair.val))
			}
			if err := binary.Write(buf, binary.BigEndian, []byte(pair.val)[:strLen]); err != nil {
				return fmt.Errorf("error encoding string %s with length %d", pair.val, len(pair.val))
			}
		}
	}
	return nil
}

func (t *TaskCreateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.Status); err != nil {
		return nil, err
	}
	err := encodeString(
		buf,
		stringEncodePair{t.TaskName, Byte},
		stringEncodePair{t.Description, Uint64},
		stringEncodePair{t.ProjectName, Byte},
		stringEncodePair{t.DueDate, Byte},
	)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (tg *TaskGetRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, tg.TaskId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (tl *TaskListRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, tl.Status); err != nil {
		return nil, err
	}
	err := encodeString(
		buf,
		stringEncodePair{tl.TaskName, Byte},
		stringEncodePair{tl.Description, Uint64},
		stringEncodePair{tl.ProjectName, Byte},
		stringEncodePair{tl.DueDate, Byte},
	)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (tu *TaskUpdateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, tu.TaskId); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, tu.Status); err != nil {
		return nil, err
	}
	err := encodeString(
		buf,
		stringEncodePair{tu.TaskName, Byte},
		stringEncodePair{tu.Description, Uint64},
		stringEncodePair{tu.ProjectName, Byte},
		stringEncodePair{tu.DueDate, Byte},
	)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (td *TaskDeleteRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, td.TaskId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (pc *ProjectCreateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)

	err := encodeString(
		buf,
		stringEncodePair{pc.ProjectName, Byte},
		stringEncodePair{pc.Description, Uint64},
	)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (pg *ProjectGetRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, pg.ProjectId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (pl *ProjectListRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := encodeString(
		buf,
		stringEncodePair{pl.ProjectName, Byte},
		stringEncodePair{pl.Description, Uint64},
	)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (pu *ProjectUpdateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, pu.ProjectId); err != nil {
		return nil, err
	}
	err := encodeString(
		buf,
		stringEncodePair{pu.ProjectName, Byte},
		stringEncodePair{pu.Description, Uint64},
	)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (pd *ProjectDeleteRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, pd.ProjectId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *NoteCreateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, nc.TaskId); err != nil {
		return nil, err
	}
	err := encodeString(buf, stringEncodePair{nc.NoteDetail, Uint64})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *NoteListRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, nc.TaskId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *NoteUpdateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, nc.NoteId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *NoteDeleteRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, nc.NoteId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *ReminderCreateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := encodeString(buf, stringEncodePair{nc.Time, Byte})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *ReminderGetRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, nc.ReminderId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *ReminderListRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := encodeString(buf, stringEncodePair{nc.DueTime, Byte})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *ReminderUpdateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, nc.ReminderId); err != nil {
		return nil, err
	}
	if err := encodeString(buf, stringEncodePair{nc.DueTime, Byte}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nc *ReminderDeleteRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, nc.ReminderId); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
