package remits

import (
	"github.com/fxamacker/cbor/v2"
)

type Iterator struct {
	LogName string `cbor:"log_name"`
	Name    string `cbor:"iterator_name"`
	Kind    string `cbor:"iterator_kind"`
	Func    string `cbor:"iterator_func"`
}

type IteratorNext struct {
	Name      string `cbor:"iterator_name"`
	MessageID int    `cbor:"message_id"`
	Count     int    `cbor:"count"`
}

type ItoratorListReq struct {
	LogName string `cbor:"log_name"`
}

type ItoratorDelReq struct {
	LogName string `cbor:"log_name"`
	Name    string `cbor:"iterator_name"`
}

func (c *Client) ListIterators(log string) ([]string, error) {
	payload, err := cbor.Marshal(ItoratorListReq{log})
	if err != nil {
		return []string{}, err
	}

	resp, err := c.Send(RequestFrame, ListIteratorCode, payload)
	if err != nil {
		return []string{}, err
	}

	data, err := resp.DataBytes()
	if err != nil {
		return []string{}, err
	}

	var iterator string
	err = cbor.Unmarshal(data, &iterator)
	if err != nil {
		return []string{}, err
	}
	return []string{iterator}, nil
}

func (c *Client) AddIterator(log, name, k, f string) (string, error) {
	payload, err := cbor.Marshal(Iterator{log, name, k, f})
	if err != nil {
		return "", err
	}

	resp, err := c.Send(RequestFrame, AddIteratorCode, payload)
	if err != nil {
		return "", err
	}

	info, err := resp.InfoString()
	if err != nil {
		return "", err
	}
	return info, nil
}

func (c *Client) NextIterator(name string, id, count int) ([]byte, error) {
	payload, err := cbor.Marshal(IteratorNext{name, id, count})
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.Send(RequestFrame, NextIteratorCode, payload)
	if err != nil {
		return []byte{}, err
	}

	data, err := resp.DataBytes()
	if err != nil {
		return []byte{}, err
	}

	// messages are currently returned as a 1-based map
	// instead of byte slices due to Lua limitations
	var vMap map[int]byte
	err = cbor.Unmarshal(data, &vMap)
	if err != nil {
		return []byte{}, err
	}

	vSlice := make([]byte, len(vMap))
	for i, _ := range vSlice {
		vSlice[i] = vMap[i+1]
	}
	return vSlice, nil
}

func (c *Client) DeleteIterator(log, name string) (string, error) {
	payload, err := cbor.Marshal(ItoratorDelReq{log, name})
	if err != nil {
		return "", err
	}

	resp, err := c.Send(RequestFrame, DeleteIteratorCode, payload)
	if err != nil {
		return "", err
	}

	info, err := resp.InfoString()
	if err != nil {
		return "", err
	}
	return info, nil
}
