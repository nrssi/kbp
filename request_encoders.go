package kbp

import (
	"bytes"
	"encoding/binary"
)

func (tc *TaskCreateRequest) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	tc.taskNameLen = byte(len(tc.TaskName))
	if err := binary.Write(buf, binary.BigEndian, tc.taskNameLen); err != nil {
		return nil, err
	}
	taskNameBytes := []byte(tc.TaskName)
	if err := binary.Write(buf, binary.BigEndian, taskNameBytes[:tc.taskNameLen]); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, tc.descriptionLen); err != nil {
		return nil, err
	}
	descriptionBytes := []byte(tc.Description)
	if err := binary.Write(buf, binary.BigEndian, descriptionBytes[:tc.descriptionLen]); err != nil {
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
