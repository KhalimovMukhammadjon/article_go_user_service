package client

import "article/article_go_user_service/config"

type ServiceManagerI interface {
}

type grpcClients struct {
}

func NewGrpcClient(cfg config.Config) (ServiceManagerI, error) {
	return nil, nil
}
