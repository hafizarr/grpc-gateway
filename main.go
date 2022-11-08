package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hafizarr/grpc-gateway/config"
	"github.com/hafizarr/grpc-gateway/proto/helloworld"
	"github.com/hafizarr/grpc-gateway/proto/users"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GreeterServer struct {
	helloworld.GreeterServer
}

type UserServer struct {
	users.UnimplementedUsersServiceServer
}

var Users = make(map[int64]*users.Users)

func main() {
	srv := grpc.NewServer()
	userSrv := new(UserServer)
	users.RegisterUsersServiceServer(srv, userSrv)

	log.Println("Starting User Server at ", config.SERVICE_USERS_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_USERS_PORT)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}

	// setup http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost"+config.SERVICE_USERS_PORT, "gRPC server endpoint")
		_ = users.RegisterUsersServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
		log.Println("Starting User Server HTTP at 9001 ")
		http.ListenAndServe(":9001", mux)
	}()

	log.Fatalln(srv.Serve(listener))
}

func (s *GreeterServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: in.Name + " world"}, nil
}

func (t *UserServer) GetAll(ctx context.Context, req *emptypb.Empty) (*users.GetAllResponse, error) {
	var listUsers []*users.Users
	for _, v := range Users {
		listUsers = append(listUsers, &users.Users{
			Id:       v.GetId(),
			Name:     v.GetName(),
			Password: v.GetPassword(),
		})
	}
	return &users.GetAllResponse{Data: listUsers}, nil
}

func (t *UserServer) GetByID(ctx context.Context, req *users.GetByIDRequest) (*users.GetByIDResponse, error) {
	todo, ok := Users[req.GetId()]
	if !ok {
		return nil, errors.New("not found")
	}
	return &users.GetByIDResponse{Data: todo}, nil
}

func (t *UserServer) Create(ctx context.Context, req *users.Users) (*users.MutationResponse, error) {
	Users[req.GetId()] = &users.Users{
		Id:       req.GetId(),
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}
	msg := fmt.Sprintf("%d successfully saved", req.GetId())
	return &users.MutationResponse{Success: msg}, nil
}

func (t *UserServer) Update(ctx context.Context, req *users.UpdateRequest) (*users.MutationResponse, error) {
	Users[req.GetId()] = &users.Users{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}
	msg := fmt.Sprintf("%d successfully updated", req.GetId())
	return &users.MutationResponse{Success: msg}, nil
}

func (t *UserServer) Delete(ctx context.Context, req *users.DeleteRequest) (*users.MutationResponse, error) {
	delete(Users, req.GetId())
	msg := fmt.Sprintf("%d successfully deleted", req.GetId())
	return &users.MutationResponse{Success: msg}, nil
}
