package repositories

import (
	"github.com/go-xorm/xorm"
	"../models"
)
type ProductRepository struct{
	engine *xorm.Engine
}

func NewProductRepository(engine *xorm.Engine) *ProductRepository {
	return &ProductRepository{
		engine:engine,
	}
}

func (d *ProductRepository) Get(id int64,P_name string) []models.Product {
	datalist := []models.Product{}
	err := d.engine.Where("productname=?", P_name).And("ID",id).Find(&datalist)
	if err == nil {
		return datalist
	}else {
		return datalist
	}
}

func (d *ProductRepository) GetAll() []models.Product {
	datalist := make([]models.Product, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *ProductRepository) Search(P_name string) []models.Product {
	datalist := make([]models.Product, 0)
	err := d.engine.Where("productname=?", P_name).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *ProductRepository) Delete(id int64) error {
	data := &models.Product{ID:id}
	_, err := d.engine.Id(data.ID).Update(data)
	return err
}

func (d *ProductRepository) Update(data *models.Product, columns []string) error {
	_, err := d.engine.Id(data.ID).MustCols(columns...).Update(data)
	return err
}

func (d *ProductRepository) Create(data *models.Product) error {
	_, err := d.engine.Insert(data)
	return err
}
