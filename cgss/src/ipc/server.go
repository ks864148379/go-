package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string 'json/"method"'
	Params string 'json:"params"'
}

type Response struct {
	Code string 'json:"code"'
	Body string 'json:"body"'
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server)  *IpcServer{
	return &IpcServer{Server}
}

func (server *IpcServer)Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c

			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format")
				return
			}

			resp := server.Handle(req.Method, req.Params)

			b, err := json.Marchal(resp)

			c <- string(b)
		}

		fmt.Println("session closed")
	}(session)

	fmt.Println("A new session has been created successfully!")

	return session
}