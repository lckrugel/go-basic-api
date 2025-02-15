package controller

import (
	"encoding/json"
	"net/http"

	"github.com/lckrugel/go-basic-api/internal/app/model"
	"github.com/lckrugel/go-basic-api/internal/app/service"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.userService.Create(r.Context(), service.CreateUserDTO{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {}

func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {}

func (u *UserController) Show(w http.ResponseWriter, r *http.Request) {}
