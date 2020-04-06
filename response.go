package remits

import (
	"encoding/binary"
	"errors"
	"github.com/fxamacker/cbor/v2"
)

const (
	// kind
	RequestFrame = uint8(0)
	InfoFrame    = uint8(1)
	DataFrame    = uint8(2)
	ErrorFrame   = uint8(3)

	// code
	ShowLogCode        = uint8(0)
	AddLogCode         = uint8(1)
	DeleteLogCode      = uint8(2)
	ListLogCode        = uint8(3)
	AddMessageCode     = uint8(4)
	AddIteratorCode    = uint8(5)
	ListIteratorCode   = uint8(6)
	NextIteratorCode   = uint8(7)
	DeleteIteratorCode = uint8(8)
)

type Response struct {
	Kind    uint8
	Code    uint8
	Payload []byte
}

func (r *Response) Error() error {
	if r.Kind != ErrorFrame {
		return errors.New("Not an Error Frame")
	}

	var errorString string
	err := cbor.Unmarshal(r.Payload, &errorString)
	if err != nil {
		return err
	}
	return errors.New(errorString)
}

func (r *Response) InfoString() (string, error) {
	if r.Kind == ErrorFrame {
		return "", r.Error()
	}
	if r.Kind != InfoFrame {
		return "", errors.New("Not a Info Frame")
	}

	var infoString string
	err := cbor.Unmarshal(r.Payload, &infoString)
	if err != nil {
		return "", err
	}
	return infoString, nil
}

// Currently only handles first message...
func (r *Response) DataBytes() ([]byte, error) {
	if r.Kind == ErrorFrame {
		return nil, r.Error()
	}
	if r.Kind != DataFrame {
		return nil, errors.New("Not a Data Frame")
	}

	// Length of first message
	msgLen := binary.BigEndian.Uint32(r.Payload[:4])
	return r.Payload[4 : 4+msgLen], nil
}
