package remits

import (
	"encoding/binary"
	"io"
	"net"
)

type Client struct {
	conn net.Conn // should we export this?
	// timeout
	// auth?
	// options?
}

func Connect(url string) (Client, error) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return Client{}, err
	}
	client := Client{
		conn: conn,
	}
	return client, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Send(kind, code uint8, payload []byte) (Response, error) {
	// SEND
	binary.Write(c.conn, binary.BigEndian, uint32(len(payload)+2))
	binary.Write(c.conn, binary.BigEndian, kind)
	binary.Write(c.conn, binary.BigEndian, code)
	c.conn.Write(payload)

	// RESPONSE FRAME LENGTH
	respLengthBuf, err := c.readBytes(4)
	if err != nil {
		return Response{}, err
	}
	respLength := binary.BigEndian.Uint32(respLengthBuf)

	// FRAME KIND, CODE
	respInfoBuf, err := c.readBytes(2)
	if err != nil {
		return Response{}, err
	}

	// PAYLOAD
	payloadBuf, err := c.readBytes(int(respLength - 2))
	if err != nil {
		return Response{}, err
	}

	response := Response{
		Kind:    respInfoBuf[0],
		Code:    respInfoBuf[1],
		Payload: payloadBuf,
	}
	return response, nil
}

func (c *Client) readBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(c.conn, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
