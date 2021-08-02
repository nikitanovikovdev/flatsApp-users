package main

import (
	"github.com/nikitanovikovdev/flatsApp-users/internal"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/platform/repository"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/users"
	authorization "github.com/nikitanovikovdev/flatsApp-users/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewMongoDB(&repository.MongoConfig{
		Host: viper.GetString("mongodb.host"),
		Port: viper.GetString("mongodb.port"),
	})
	if err != nil {
		log.Fatalf("failed to create new mongo repository: %v", err)
	}

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	handler := users.NewHandler(service)

	server := internal.NewServer(viper.GetString("server.host"), viper.GetString("server.port"), users.NewRouter(handler))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}


	s := grpc.NewServer()
	srv := &internal.GRPCServer{}
	authorization.RegisterAuthServer(s, srv)

	l, err := net.Listen("tcp", ":8040")
	if err != nil {
		log.Fatalf("failed to connection: %v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//s := grpc.NewServer()
	//srv  := internal.NewGRPCServer(handler)
	//authorization.RegisterAuthServer(s, srv)
	//
	//l, err := net.Listen("tcp", ":8040")
	//if err != nil {
	//	log.Fatalf("failed to connection: %v", err)
	//}
	//
	//if err := s.Serve(l); err != nil {
	//	log.Fatalf("failed to listn: %v", err)
	//}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
