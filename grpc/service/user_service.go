package service

import (
	"article/article_go_user_service/config"
	"article/article_go_user_service/genproto/user_service"
	"article/article_go_user_service/grpc/client"
	"article/article_go_user_service/pkg/logger"
	"article/article_go_user_service/storage"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StoragI
	services client.ServiceManagerI
	user_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StoragI, svcs client.ServiceManagerI) *userService {
	return &userService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (u userService) Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.User, err error) {
	u.log.Info("---CreateUser--->", logger.Any("req", req))

	pKey, err := u.strg.User().Create(context.Background(), req)
	if err != nil {
		u.log.Error("!!!CreateUser!!!", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = u.strg.User().GetById(ctx, pKey)
	return nil, nil
}
