package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hafizarr/grpc-gateway/common/config"
	"github.com/hafizarr/grpc-gateway/common/models"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoServer struct {
	models.UnimplementedTodoServiceServer
}

var Todos = make(map[string]*models.Todo)

func main() {
	srv := grpc.NewServer()
	userSrv := new(TodoServer)
	models.RegisterTodoServiceServer(srv, userSrv)

	log.Println("Starting Todo Server at ", config.SERVICE_USERS_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_USERS_PORT)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}

	// setup http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost"+config.SERVICE_USERS_PORT, "gRPC server endpoint")
		_ = models.RegisterTodoServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
		log.Println("Starting Todo Server HTTP at 9001 ")
		http.ListenAndServe(":9001", mux)
	}()

	log.Fatalln(srv.Serve(listener))
}

func (t *TodoServer) GetAll(ctx context.Context, req *emptypb.Empty) (*models.GetAllResponse, error) {
	var todos []*models.Todo
	for _, v := range Todos {
		todos = append(todos, &models.Todo{
			Id:   v.GetId(),
			Name: v.GetName(),
		})
	}
	return &models.GetAllResponse{Data: todos}, nil
}
