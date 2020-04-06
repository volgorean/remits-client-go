package remits

import (
	"github.com/fxamacker/cbor/v2"
	"time"
)

type Log struct {
	Name      string    `cbor:"log_name"`
	CreatedAt time.Time `cbor:"created_at",omitempty`
}

type ShowLogResp struct {
	Name      string    `cbor:"name"`
	CreatedAt time.Time `cbor:"created_at",omitempty`
}

func (c *Client) ListLogs() ([]string, error) {
	resp, err := c.Send(RequestFrame, ListLogCode, nil)
	if err != nil {
		return []string{}, err
	}

	data, err := resp.DataBytes()
	if err != nil {
		return []string{}, err
	}

	var logs []string
	err = cbor.Unmarshal(data, &logs)
	if err != nil {
		return []string{}, err
	}
	return logs, nil
}

func (c *Client) AddLog(name string) (string, error) {
	payload, err := cbor.Marshal(Log{Name: name})
	if err != nil {
		return "", err
	}

	resp, err := c.Send(RequestFrame, AddLogCode, payload)
	if err != nil {
		return "", err
	}

	info, err := resp.InfoString()
	if err != nil {
		return "", err
	}
	return info, nil
}

func (c *Client) ShowLog(name string) (Log, error) {
	payload, err := cbor.Marshal(Log{Name: name})
	if err != nil {
		return Log{}, err
	}

	resp, err := c.Send(RequestFrame, ShowLogCode, payload)
	if err != nil {
		return Log{}, err
	}

	data, err := resp.DataBytes()
	if err != nil {
		return Log{}, err
	}

	var log ShowLogResp
	err = cbor.Unmarshal(data, &log)
	if err != nil {
		return Log{}, err
	}
	return Log{Name: log.Name, CreatedAt: log.CreatedAt}, nil
}

func (c *Client) DeleteLog(name string) (string, error) {
	payload, err := cbor.Marshal(Log{Name: name})
	if err != nil {
		return "", err
	}

	resp, err := c.Send(RequestFrame, DeleteLogCode, payload)
	if err != nil {
		return "", err
	}

	info, err := resp.InfoString()
	if err != nil {
		return "", err
	}
	return info, nil
}
