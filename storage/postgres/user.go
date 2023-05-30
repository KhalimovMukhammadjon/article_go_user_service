package postgres

import (
	"article/article_go_user_service/genproto/user_service"
	"article/article_go_user_service/models"
	"article/article_go_user_service/pkg/helper"
	"article/article_go_user_service/storage"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.UserRepoI {
	return &userRepo{
		db: db,
	}
}

// Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.PrimaryKey, err error)
// GetById(ctx context.Context, req *user_service.PrimaryKey) (resp *user_service.User, err error)
// GetList(ctx context.Context, req *user_service.GetAllUserRequest) (resp *user_service.GetAllUserResponse, err error)
// Update(ctx context.Context, req *user_service.UpdateUserRequest) error
// PatchUpdate(ctx context.Context, req *user_service.UpdatePatchUser) (rowsAffected int64, err error)
// Delete(ctx context.Context, req *user_service.PrimaryKey) (err error)

func (u *userRepo) Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.PrimaryKey, err error) {
	query := `
		INSERT INTO users
			(id,first_name,last_name,phone_number)
		VALUES(
			$1,
			$2,
			$3,
			$4,
		) 
	`
	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = u.db.Exec(ctx, query,
		uuid.String(),
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
	)
	if err != nil {
		return resp, err
	}

	resp = &user_service.PrimaryKey{
		Id: uuid.String(),
	}

	return resp, nil
}

func (u *userRepo) GetById(ctx context.Context, req *user_service.PrimaryKey) (resp *user_service.User, err error) {
	resp = &user_service.User{}

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			phone_number,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`
	err = u.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.FirstName,
		&resp.LastName,
		&resp.PhoneNumber,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *userRepo) GetList(ctx context.Context, req *user_service.GetAllUserRequest) (resp *user_service.GetAllUserResponse, err error) {
	resp = &user_service.GetAllUserResponse{}

	var (
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*),OVER(),
			id,
			first_name,
			last_name,
			phone_number,
			created_at,
			updated_at
		FROM users
	`

	if len(req.Search) > 0 {
		filter += " AND first_name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user user_service.User
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.User = append(resp.User, &user)
	}

	return resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *user_service.UpdateUserRequest) (rowsAffected int64, err error) {
	query := `
		UPDATE users SET
			first_name = :first_name,
			last_name = :last_name,
			phone_number = :phone_number
			update_at = NOW()
		WHERE id = :id
	`

	params := map[string]interface{}{
		"id":           req.User.Id,
		"first_name":   req.User.FirstName,
		"last_name":    req.User.LastName,
		"phone_number": req.User.PhoneNumber,
		"updated_at":   req.User.UpdatedAt,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return result.RowsAffected(), err
	}

	return result.RowsAffected(), err
}

func (u *userRepo) PatchUpdate(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {
	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Fields)-1 {
			set += ", "
		}
		ind++
	}

	query = `
	  UPDATE
		"users"
	 + set +  , updated_at = now()
	  WHERE
		id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (u *userRepo) Delete(ctx context.Context, req *user_service.PrimaryKey) (err error) {
	query := `DELETE FROM users WHERE id = $1`

	_, err = u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
