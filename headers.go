package kbp

type RequestOpType byte

const Version byte = 1

type StatusCode byte

type ContentType byte

type TaskStatus byte

const (
	Todo       TaskStatus = 0x01
	InProgress TaskStatus = 0x02
	Done       TaskStatus = 0x03
	Blocked    TaskStatus = 0x04
)

const (
	CreateTask     RequestOpType = 0x01
	UpdateTask     RequestOpType = 0x02
	GetTask        RequestOpType = 0x03
	ListTask       RequestOpType = 0x04
	DeleteTask     RequestOpType = 0x05
	CreateProject  RequestOpType = 0x06
	UpdateProject  RequestOpType = 0x07
	GetProject     RequestOpType = 0x08
	ListProject    RequestOpType = 0x09
	DeleteProject  RequestOpType = 0x0A
	CreatNote      RequestOpType = 0x0B
	ListNote       RequestOpType = 0x0C
	DeleteNote     RequestOpType = 0x0D
	CreateReminder RequestOpType = 0x0E
	GetReminder    RequestOpType = 0x0F
	ListReminder   RequestOpType = 0x10
	DeleteReminder RequestOpType = 0x11
)

const (
	Success StatusCode = 0x01
	Error   StatusCode = 0x02
)

const (
	Empty  ContentType = 0x01
	Single ContentType = 0x02
	List   ContentType = 0x03
)

type KbpPayload interface {
	Encode() ([]byte, error)
	Decode(data []byte) error
}

type KbpRequestHeader struct {
	Version    byte
	RequestOp  RequestOpType
	bodyLength uint64
}

type KbpResponseHeader struct {
	Version     byte
	RequestOp   RequestOpType
	Status      StatusCode
	ContentType ContentType
	bodyLength  uint64
}
