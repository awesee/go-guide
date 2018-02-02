package main

import "fmt"

type StuffClient interface {
	DoStuff() error
}

type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}

var defaultStuffClient = stuffClient{
	timeout: 2,
	retries: 3,
}

type StuffClientOption func(*stuffClient)

func WithRetries(r int) StuffClientOption {
	return func(o *stuffClient) {
		o.retries = r
	}
}

func WithTimeout(t int) StuffClientOption {
	return func(o *stuffClient) {
		o.timeout = t
	}
}

type Connection struct{}

func NewStuffClient(conn Connection, opts ...StuffClientOption) StuffClient {
	client := defaultStuffClient
	for _, o := range opts {
		o(&client)
	}

	client.conn = conn
	return client
}
func (c stuffClient) DoStuff() error {
	return nil
}

func main() {

	client := NewStuffClient(struct{}{})
	fmt.Println(client)
}
