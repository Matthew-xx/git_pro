package repositories


import (
	"github.com/go-xorm/xorm"
	"../models"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type UserRepository struct{
	engine *xorm.Engine
	Session *sessions.Session
}

func NewUserRepository(engine *xorm.Engine) *UserRepository {
	return &UserRepository{
		engine:engine,
	}
}

func (d *UserRepository) Get(id int64) []models.User {
	data := []models.User{}
	err := d.engine.Where("id", id).Find(data)
	if err == nil {
		return data
	}else {
		return data
	}
}

func (d *UserRepository) GetAll() []models.User {
	datalist := make([]models.User, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *UserRepository) Logout() mvc.Result{
	return mvc.Response{
		Path: "/index",
	}
}


func (d *UserRepository) Search(P_name string) []models.User {
	datalist := make([]models.User, 0)
	err := d.engine.Where("productname=?", P_name).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

//删除
func (d *UserRepository) Delete(id int64) error {
	data := &models.User{ID:id}
	_, err := d.engine.Id(data.ID).Update(data)
	return err
}
//更新信息
func (d *UserRepository) Update(data *models.User, columns []string) error {
	_, err := d.engine.Id(data.ID).MustCols(columns...).Update(data)
	return err
}
//创建user
func (d *UserRepository) Create(data *models.User) error {
	_, err := d.engine.Insert(data)
	return err
}

//验证用户登录
func (d *UserRepository) GetByUsernameAndPassword(username, userPassword string) bool{
	var isexist =false
	if username == "" || userPassword == "" {
		isexist = false
		return isexist
	}else {
		if d.engine.Select(username).Where(userPassword) != nil {
			isexist = true
			return isexist
		}
	}
	return isexist
}

//用户中心
func (d *UserRepository) GetByID(id int64) ([]models.User,bool) {
	data := []models.User{}
	err := d.engine.Where(id).Find(data)
	if err != nil {
		return data,false
	} else {
		return data,true
	}
}
