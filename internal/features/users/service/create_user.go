package users_service

import (
	"context"
	"fmt"

	"github.com/Mikton22/golang-todo/internal/core/domain"
	core_errors "github.com/Mikton22/golang-todo/internal/core/errors"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w, err", core_errors.ErrInvalidArgument)
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
