// Code generated by Kitex v0.6.1. DO NOT EDIT.

package studentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	demo "rpc_server/kitex_gen/demo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, student *demo.Student, callOptions ...callopt.Option) (r *demo.RegisterResp, err error)
	Query(ctx context.Context, req *demo.QueryReq, callOptions ...callopt.Option) (r *demo.Student, err error)
	GetPort(ctx context.Context, req *demo.GetPortReq, callOptions ...callopt.Option) (r *demo.GetPortResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kStudentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kStudentServiceClient struct {
	*kClient
}

func (p *kStudentServiceClient) Register(ctx context.Context, student *demo.Student, callOptions ...callopt.Option) (r *demo.RegisterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, student)
}

func (p *kStudentServiceClient) Query(ctx context.Context, req *demo.QueryReq, callOptions ...callopt.Option) (r *demo.Student, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Query(ctx, req)
}

func (p *kStudentServiceClient) GetPort(ctx context.Context, req *demo.GetPortReq, callOptions ...callopt.Option) (r *demo.GetPortResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPort(ctx, req)
}
