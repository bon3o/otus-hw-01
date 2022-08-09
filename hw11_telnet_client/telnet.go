package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *client) Connect() error {
	if c.in == nil {
		return fmt.Errorf("in is invalid")
	}
	if c.out == nil {
		return fmt.Errorf("out is invalid")
	}

	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	if conn == nil {
		return fmt.Errorf("connection not established")
	}
	c.conn = conn
	return nil
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Send() error {
	if _, err := io.Copy(c.conn, c.in); err != nil {
		return err
	}
	return nil
}

func (c *client) Receive() error {
	if _, err := io.Copy(c.out, c.conn); err != nil {
		return err
	}
	return nil
}
