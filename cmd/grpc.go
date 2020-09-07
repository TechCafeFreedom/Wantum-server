package main

import (
	"fmt"
	"log"
	"net"
	"time"
	grpcWire "wantum/cmd/grpc"
	grpcMiddleware "wantum/pkg/api/middleware/grpc"
	middleware "wantum/pkg/api/middleware/grpc"
	"wantum/pkg/domain/repository"
	"wantum/pkg/pb"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(firebaseClient grpcMiddleware.FirebaseAuth, masterTxManager repository.MasterTxManager) {
	// APIインスタンスの作成
	userAPI := grpcWire.InitUserAPI(masterTxManager)

	// gRPC: 8011
	port := 8011
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Run server port: %d", port)

	// gRPC Server Option Set
	ops := make([]grpc.ServerOption, 0)
	ops = append(ops,
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_auth.UnaryServerInterceptor(firebaseClient.MiddlewareFunc()),
				middleware.UnaryErrorHandling(),
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_auth.StreamServerInterceptor(firebaseClient.MiddlewareFunc()),
				middleware.StreamErrorHandling(),
			),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    5 * time.Second,
			Timeout: 5 * time.Hour,
		}),
	)
	grpcServer := grpc.NewServer(
		ops...,
	)

	// User Service
	pb.RegisterUserServiceServer(grpcServer, &userAPI)

	// Serve
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
