package grpc

import (
	"article/article_go_user_service/config"
	"article/article_go_user_service/genproto/user_service"
	"article/article_go_user_service/grpc/client"
	"article/article_go_user_service/grpc/service"
	"article/article_go_user_service/pkg/logger"
	"article/article_go_user_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StoragI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	user_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)

	return
}
