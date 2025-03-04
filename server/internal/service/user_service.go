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

func mapDBUserToDomainUser(dbUser db.User) domain.User {
	return domain.User{
		ID:         int(dbUser.ID),
		FirstName:  dbUser.FirstName,
		LastName:   dbUser.LastName,
		Email:      dbUser.Email,
		PictureURL: dbUser.PictureUrl.String,
	}
}

func (s *UserService) LikeMovie(ctx context.Context, userID, movieID int) error {
	return s.userRepo.LikeMovie(ctx, userID, movieID)
}

func (s *UserService) UnlikeMovie(ctx context.Context, userID, movieID int) error {
	return s.userRepo.UnlikeMovie(ctx, userID, movieID)
}

func (s *UserService) GetLikedMovies(ctx context.Context, userID int) ([]domain.Movie, error) {
	dbMovies, err := s.userRepo.GetLikedMovies(ctx, userID)
	if err != nil {
		return nil, err
	}

	movieMap := make(map[int]domain.Movie)

	for _, row := range dbMovies {
		movieID := int(row.MovieID)
		movie, exists := movieMap[movieID]
		userRating, _ := row.UserRating.Float64Value()

		if !exists {
			movie = domain.Movie{
				ID:          movieID,
				Title:       row.Title,
				ReleaseDate: row.ReleaseDate.Time,
				RunTime:     int(row.Runtime.Int32),
				MPAARating:  row.MpaaRating.String,
				Description: row.Description.String,
				Image:       row.Image.String,
				UserRating:  userRating.Float64,
				Video:       row.Video.String,
				Genres:      []*domain.Genre{},
			}
		}

		if row.GenreID.Valid {
			movie.Genres = append(movie.Genres, &domain.Genre{
				ID:    int(row.GenreID.Int32),
				Genre: row.Genre.String,
			})
		}

		movieMap[movieID] = movie
	}

	movies := make([]domain.Movie, 0)
	for _, movie := range movieMap {
		movies = append(movies, movie)
	}

	return movies, nil
}
