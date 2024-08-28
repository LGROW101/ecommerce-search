package service

import (
	"github.com/LGROW101/ecommerce-search/model"
	"github.com/LGROW101/ecommerce-search/repository"
)

type SearchService struct {
	productRepo *repository.ProductRepository
}

func NewSearchService(productRepo *repository.ProductRepository) *SearchService {
	return &SearchService{productRepo: productRepo}
}

func (s *SearchService) Search(query string) ([]model.Product, error) {
	return s.productRepo.Search(query)
}

func (s *SearchService) GetProductDetails(id uint) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}
