package main

import (
	"context"
	demo "rpc_server/kitex_gen/demo"
	"strconv"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

var studentMap = make(map[int32]*demo.Student)

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	resp = demo.NewRegisterResp()
	if _, exist := studentMap[student.Id]; exist {
		resp.Success = false
		resp.Message = "Student Already Exists"
	} else {
		studentMap[student.Id] = student
		resp.Success = true
		resp.Message = "Register Success"
	}
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	resp = demo.NewStudent()
	if _, exist := studentMap[req.Id]; !exist {
		resp.Id = -1
		resp.Name = "Student Not Exist"
		resp.College = &demo.College{Name: "Unknown", Address: "Unknown"}
		resp.Email = nil
		resp.Sex = nil
	} else {
		resp = studentMap[req.Id]
	}
	return
}

// GetPort implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) GetPort(ctx context.Context, req *demo.GetPortReq) (resp *demo.GetPortResp, err error) {
	resp = new(demo.GetPortResp)
	resp.Port = strconv.Itoa(bindPort)
	return
}
