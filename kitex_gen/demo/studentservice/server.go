// Code generated by Kitex v0.6.1. DO NOT EDIT.
package studentservice

import (
	server "github.com/cloudwego/kitex/server"
	demo "rpc_server/kitex_gen/demo"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler demo.StudentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
