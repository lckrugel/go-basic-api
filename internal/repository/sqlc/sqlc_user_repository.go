package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lckrugel/go-basic-api/internal/app/model"
)

type SQLCUserRepository struct {
	queries *Queries
}

func NewSQLCUserRepository(q *Queries) *SQLCUserRepository {
	return &SQLCUserRepository{
		queries: q,
	}
}

func (s *SQLCUserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	sqlcUsers, err := s.queries.ListUsers(ctx)
	modelUsers := convertToModelSlice(sqlcUsers)
	return modelUsers, err
}

func (s *SQLCUserRepository) FindById(ctx context.Context, id int32) (model.User, error) {
	sqlcUser, err := s.queries.GetUser(ctx, id)
	modelUser := convertoToModel(sqlcUser)
	return modelUser, err
}

func (s *SQLCUserRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	argUser := converToSQLC(user)
	sqlcUser, err := s.queries.CreateUser(ctx, CreateUserParams{
		Username: argUser.Username,
		Email:    argUser.Email,
		Password: argUser.Password,
	})
	modelUser := convertoToModel(sqlcUser)
	return modelUser, err
}

func (s *SQLCUserRepository) Update(ctx context.Context, user model.User) (model.User, error) {
	argUser := converToSQLC(user)
	sqlcUser, err := s.queries.UpdateUser(ctx, UpdateUserParams(argUser))
	modelUser := convertoToModel(sqlcUser)
	return modelUser, err
}

func (s *SQLCUserRepository) Delete(ctx context.Context, id int32) error {
	return s.queries.DeleteUser(ctx, id)
}

func (s *SQLCUserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	pgEmail := pgtype.Text{
		String: email,
		Valid:  true,
	}
	sqlcUser, err := s.queries.GetUserByEmail(ctx, pgEmail)
	modelUser := convertoToModel(sqlcUser)
	return modelUser, err
}

func (s *SQLCUserRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	pgUsername := pgtype.Text{
		String: username,
		Valid:  true,
	}
	sqlcUser, err := s.queries.GetUserByUsername(ctx, pgUsername)
	modelUser := convertoToModel(sqlcUser)
	return modelUser, err
}

func convertoToModel(user User) model.User {
	return model.User{
		ID:       user.ID,
		Username: user.Username.String,
		Email:    user.Email.String,
		Password: user.Password.String,
	}
}

func converToSQLC(user model.User) User {
	return User{
		ID:       user.ID,
		Username: pgtype.Text{String: user.Username, Valid: true},
		Email:    pgtype.Text{String: user.Email, Valid: true},
		Password: pgtype.Text{String: user.Password, Valid: true},
	}
}

func convertToModelSlice(users []User) []model.User {
	var modelUsers []model.User
	for _, user := range users {
		modelUsers = append(modelUsers, convertoToModel(user))
	}
	return modelUsers
}
