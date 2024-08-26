package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	desc "github.com/dogmego/MicroservicesPractice/Auth/grpc/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	log.Printf("user created on email: %v, name: %v, role: %v",
		req.Email, req.Name, req.Role)

	return &desc.CreateUserResponse{Id: gofakeit.Int64()}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	log.Printf("get user with id: %d",
		req.Id)

	return &desc.GetUserResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      0,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateUserInfo) (*emptypb.Empty, error) {
	log.Printf("update user with email: %v; name: %v; ID: %v",
		req.Email, req.Name, req.Id)

	return nil, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Printf("delete user with id: %d",
		req.Id)

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("grpc_server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
