package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/apperror"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type UserService struct {
	q *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{q: q}
}

func (s *UserService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*db.User, error) {
	existing, err := s.GetByEmail(ctx, req.Email)
	if err == nil {
		return existing, nil
	}

	user, err := s.q.CreateUser(
		ctx,
		db.CreateUserParams{
			Column1:  uuid.New(),
			Username: req.Name,
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	slog.Info("User created", "email", req.Email)
	return &user, nil
}

func (s *UserService) GetByEmail(
	ctx context.Context,
	email string,
) (*db.User, error) {
	user, err := s.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, apperror.NotFound("User not found")
	}
	return &user, nil
}

func (s *UserService) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*db.User, error) {
	user, err := s.q.GetUser(ctx, id)
	if err != nil {
		return nil, apperror.NotFound("User not found")
	}
	return &user, nil
}
