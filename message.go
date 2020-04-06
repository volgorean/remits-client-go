package remits

import (
	"github.com/fxamacker/cbor/v2"
)

type Message struct {
	LogName string `cbor:"log_name"`
	Message []byte `cbor:"message"`
}

func (c *Client) AddMessage(log string, message []byte) (string, error) {
	// Remits expects the message itself to be CBOR encoded
	messageCbor, err := cbor.Marshal(message)
	if err != nil {
		return "", err
	}

	payload, err := cbor.Marshal(Message{log, messageCbor})
	if err != nil {
		return "", err
	}

	resp, err := c.Send(RequestFrame, AddMessageCode, payload)
	if err != nil {
		return "", err
	}

	info, err := resp.InfoString()
	if err != nil {
		return "", err
	}
	return info, nil
}
