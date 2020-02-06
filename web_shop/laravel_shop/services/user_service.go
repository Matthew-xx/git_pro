package services

import (
	"../models"
	"../repositories"
	"../datasource"
	"github.com/kataras/iris/mvc"
)
type UserService interface {
	Get(id int64) []models.User
	GetAll() []models.User
	Search(P_name string) []models.User
	Logout()  mvc.Result
	Create(data *models.User) error
	GetByUsernameAndPassword(username ,password string)  bool
	GetByID(id int64)  ([]models.User,bool)
}

type userService struct {
	repo *repositories.UserRepository
}

func NewUserService() UserService {
	return &userService{
		repo : repositories.NewUserRepository(datasource.InstanceMaster()),
	}
}


func (u *userService) GetAll() []models.User{
	return u.repo.GetAll()
}

func (u *userService) Get(id int64) []models.User{
	return u.repo.Get(id)
}

func (u *userService) Search(P_name string) []models.User{
	return u.repo.Search(P_name)
}

func (u *userService) Logout() mvc.Result {
	return u.repo.Logout()
}

//
func (u *userService) Create(data *models.User) error {
	return u.repo.Create(data)
}

func (u *userService) GetByUsernameAndPassword(username,password string) bool {
	return u.repo.GetByUsernameAndPassword(username,password)
}

func (u *userService) GetByID(id int64) ([]models.User,bool){
	return u.repo.GetByID(id)
}