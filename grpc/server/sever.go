package main

import (
	"context"
	"ex5/grpc/exam1"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

// Them userpartner (response)
func (server) Insert(ctx context.Context, req *exam1.InsertRequest) (*exam1.InsertResponse, error) {
	log.Printf("calling insert %+v\n", req.GetUser())
	ci := ConvertUser(req.GetUser())

	err := ci.InsertUser()

	if err != nil {
		resp := &exam1.InsertResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("insert err %v", err),
		}
		return resp, nil
	}

	resp := &exam1.InsertResponse{
		StatusCode: 1,
		Message:    "OK",
	}
	return resp, nil
}

// Doc userpartner (response)
func (server) Read(ctx context.Context, req *exam1.ReadRequest) (*exam1.ReadResponse, error) {
	log.Printf("calling read %s\n", req.GetUserId())
	ci, err := ReadUser(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Read user %s err %v", req.GetUserId(), err)
	}
	return &exam1.ReadResponse{
		User : ConvertUser1(ci),
	}, nil
}

// Sua userpartner (response)
func (server) Update(ctx context.Context, req *exam1.UpdateRequest) (*exam1.UpdateResponse, error) {
	if req.GetNewUser() == nil || req.GetUserID() == "" {
		return nil, status.Error(codes.InvalidArgument, "update req with nil user")
	}
	log.Printf("calling update with data %v,  id : %v\n", req.GetNewUser(), req.GetUserID())
	ci := ConvertUser(req.GetNewUser())
	err := ci.UpdateUser(req.GetUserID())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "update err %v", err)
	}

	update, err := ReadUser(req.GetUserID())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "try to read update user %+v err %v", req.GetNewUser(), err)
	}

	return &exam1.UpdateResponse{
		UpdateUser: ConvertUser1(update),
	}, nil
}

// Danh sach userpantner (response)
func (server) List(req *exam1.ListRequest, srv exam1.UserServer_ListServer) (error) {
	if req.GetCount() <= 0 {
		return status.Error(codes.InvalidArgument, "list req with nil user")
	}
	log.Println("Calling list with data: ")
	users,err := ListUser(req.GetCount())
	if err != nil {
		return status.Errorf(codes.Unknown, "err %v", err)
	}
	if req.GetCount() == 0 {
		log.Println("User not found")
		return nil
	}
	for i := 0; i < int(req.GetCount()); i++ {
		srv.Send(&exam1.ListResponse{
			Users: &users[i],
		})
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	exam1.RegisterUserServerServer(s, &server{})

	fmt.Println("contact service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}

}
