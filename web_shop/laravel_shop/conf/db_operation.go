package conf

import (
	"github.com/go-xorm/xorm"
	"../models"
)


type DBOperate struct {
	engine *xorm.Engine
}

func NewShopDB(engine *xorm.Engine) *DBOperate {
	return &DBOperate{
		engine:engine,
	}
}

//创建
func (op *DBOperate) Create(data *models.User) error {
	_, err := op.engine.Insert(data)
	return err
}
//更新
func (op *DBOperate) Update(data *models.User ,columns []string) error {
	_, err := op.engine.ID(data.ID).MustCols(columns...).Update(data)
	return err
}
