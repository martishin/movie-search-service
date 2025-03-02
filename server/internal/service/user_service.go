package service

import (
	"context"
	"errors"

	db "github.com/martishin/movie-search-service/internal/db/generated"
	"github.com/martishin/movie-search-service/internal/model/domain"
	"github.com/martishin/movie-search-service/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, firstName, lastName, email, pictureURL string, password string) (domain.User, error) {
	if email == "" {
		return domain.User{}, errors.New("email cannot be empty")
	}

	dbUser, err := s.userRepo.CreateUser(ctx, firstName, lastName, email, pictureURL, password)
	if err != nil {
		return domain.User{}, err
	}

	return mapDBUserToDomainUser(dbUser), nil
}

func (s *UserService) FindOrCreateUser(ctx context.Context, firstName, lastName, email, pictureURL string, password string) (domain.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return mapDBUserToDomainUser(dbUser), nil
	}

	createdUser, err := s.userRepo.CreateUser(ctx, firstName, lastName, email, pictureURL, password)
	if err != nil {
		return domain.User{}, err
	}

	return mapDBUserToDomainUser(createdUser), nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	dbUser, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return mapDBUserToDomainUser(dbUser), nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return mapDBUserToDomainUser(dbUser), nil
}

func (s *UserService) GetUserIDAndPasswordByEmail(ctx context.Context, email string) (int, string, error) {
	dbUser, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return 0, "", err
	}
	return int(dbUser.ID), dbUser.Password.String, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]domain.User, error) {
	dbUsers, err := s.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = mapDBUserToDomainUser(dbUser)
	}

	return users, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.userRepo.DeleteUser(ctx, id)
}

func mapDBUserToDomainUser(dbUser db.User) domain.User {
	return domain.User{
		ID:         int(dbUser.ID),
		FirstName:  dbUser.FirstName,
		LastName:   dbUser.LastName,
		Email:      dbUser.Email,
		PictureURL: dbUser.PictureUrl.String,
	}
}
