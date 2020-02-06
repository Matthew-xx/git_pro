package services

import (
	"../models"
	"../repositories"
	"../datasource"
)
type ProductService interface {
	Get(id int64,P_name string) []models.Product
	GetAll() []models.Product
	Search(P_name string) []models.Product
}

type productService struct {
	repo *repositories.ProductRepository
}

func NewProductService() ProductService {
	return &productService{
		repo : repositories.NewProductRepository(datasource.InstanceMaster()),
	}
}


func (p *productService) GetAll() []models.Product{
	return p.repo.GetAll()
}

func (p *productService) Get(id int64,P_name string) []models.Product{
	return p.repo.Get(id,P_name)
}

func (p *productService) Search(P_name string) []models.Product{
	return p.repo.Search(P_name)
}