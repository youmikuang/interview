package repository

import (
	"errors"
	"fmt"
	"product-matching/api/repository"
	"product-matching/config"
	"product-matching/domain/model"
)

// MockProductRepo 产品仓储适配器（实现顶级api/repository接口）
type MockProductRepo struct{}

// 编译期校验：确保实现了顶级接口
var _ repository.ProductRepository = (*MockProductRepo)(nil)

// ListAllProducts 实现顶级接口方法
func (m *MockProductRepo) ListAllProducts() ([]*model.Product, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	var products []*model.Product
	for _, p := range cfg.Products {
		product := &model.Product{
			ID:   p.ID,
			Name: p.Name,
		}
		product.FilterRules.AgeMin = p.FilterRules.AgeMin
		product.FilterRules.AgeMax = p.FilterRules.AgeMax
		product.FilterRules.AllowedRegions = p.FilterRules.AllowedRegions
		product.FilterRules.HasCar = p.FilterRules.HasCar
		product.FilterRules.HasSocial = p.FilterRules.HasSocial
		product.FilterRules.NeedRemoteCheck = p.FilterRules.NeedRemoteCheck
		products = append(products, product)
	}
	return products, nil
}

// GetProductByID 实现顶级接口方法
func (m *MockProductRepo) GetProductByID(id string) (*model.Product, error) {
	products, err := m.ListAllProducts()
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, errors.New("product not found")
}