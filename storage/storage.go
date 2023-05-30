package storage

import (
	"article/article_go_user_service/genproto/user_service"
	"article/article_go_user_service/models"
	"context"
)

type StoragI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.PrimaryKey, err error)
	GetById(ctx context.Context, req *user_service.PrimaryKey) (resp *user_service.User, err error)
	GetList(ctx context.Context, req *user_service.GetAllUserRequest) (resp *user_service.GetAllUserResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateUserRequest) (rowsAffected int64, err error)
	PatchUpdate(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.PrimaryKey) (err error)
}
