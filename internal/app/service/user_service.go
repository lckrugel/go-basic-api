package service

import (
	"context"
	"crypto"

	"github.com/lckrugel/go-basic-api/internal/app/model"
	"github.com/lckrugel/go-basic-api/internal/repository"
)

type IUserService interface {
	List(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, data CreateUserDTO) (model.User, error)
	Update(ctx context.Context, id int32, data UpdateUserDTO) (model.User, error)
	Delete(ctx context.Context, id int32) error
	GetById(ctx context.Context, id int32) (model.User, error)
}

type UserService struct {
	userRepository repository.UserRepository
}

type CreateUserDTO struct {
	Username string
	Email    string
	Password string
}

type UpdateUserDTO struct {
	Username *string
	Email    *string
	Password *string
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) List(ctx context.Context) ([]model.User, error) {
	return u.userRepository.FindAll(ctx)
}

func (u *UserService) Create(ctx context.Context, data CreateUserDTO) (model.User, error) {
	user := model.User{
		Username: data.Username,
		Email:    data.Email,
		// Password: hashPassword(data.Password)
		Password: data.Password,
	}
	return u.userRepository.Create(ctx, user)
}

func (u *UserService) Update(ctx context.Context, id int32, data UpdateUserDTO) (model.User, error) {
	user, err := u.userRepository.FindById(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	if data.Username != nil {
		user.Username = *data.Username
	}
	if data.Email != nil {
		user.Email = *data.Email
	}
	if data.Password != nil {
		user.Password = hashPassword(*data.Password)
	}

	return u.userRepository.Update(ctx, user)
}

func (u *UserService) Delete(ctx context.Context, id int32) error {
	return u.userRepository.Delete(ctx, id)
}

func (u *UserService) GetById(ctx context.Context, id int32) (model.User, error) {
	return u.userRepository.FindById(ctx, id)
}

func hashPassword(password string) string {
	hasher := crypto.SHA256.New()
	hasher.Write([]byte(password))
	return string(hasher.Sum(nil))
}
