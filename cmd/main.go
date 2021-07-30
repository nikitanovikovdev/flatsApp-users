package main

import (
	"github.com/nikitanovikovdev/flatsApp-users/internal"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/platform/repository"
	"github.com/nikitanovikovdev/flatsApp-users/pkg/users"
	"github.com/nikitanovikovdev/flatsApp-users/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"

	"log"
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


	s := grpc.NewServer()
	srv  := internal.NewGRPCServer(handler)
	auth.RegisterAuthorizationServer(s, srv)

	l, err := net.Listen("tcp", ":8040")
	if err != nil {
		log.Fatalf("failed to connection: %v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to listn: %v", err)
	}











	//
	//go func() {
	//	if err = server.Run(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<- quit
	//
	//if err := server.Shutdown(context.Background()); err != nil {
	//	log.Fatalf("error occured on server shutting down")
	//}
	//
	//if err := db.Disconnect(context.Background()); err != nil {
	//	log.Fatalf("error occured on db connection close")
	//}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
