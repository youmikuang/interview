package repository

import (
	"errors"
	"product-matching/api/repository"
	"product-matching/domain/model"
)

// MockProductRepo 产品仓储适配器（实现顶级api/repository接口）
type MockProductRepo struct{}

// 编译期校验：确保实现了顶级接口
var _ repository.ProductRepository = (*MockProductRepo)(nil)

// ListAllProducts 实现顶级接口方法
func (m *MockProductRepo) ListAllProducts() ([]*model.Product, error) {
	hasCarTrue := true
	hasSocialTrue := true
	return []*model.Product{
		{
			ID:   "P001",
			Name: "汽车贷",
			FilterRules: struct {
				AgeMin          int
				AgeMax          int
				AllowedRegions  []string
				HasCar          *bool
				HasSocial       *bool
				NeedRemoteCheck bool
			}{
				AgeMin:          20,
				AgeMax:          50,
				AllowedRegions:  []string{"北京", "天津"},
				HasCar:          &hasCarTrue,
				HasSocial:       nil,
				NeedRemoteCheck: true,
			},
		},
		{
			ID:   "P002",
			Name: "社保贷",
			FilterRules: struct {
				AgeMin          int
				AgeMax          int
				AllowedRegions  []string
				HasCar          *bool
				HasSocial       *bool
				NeedRemoteCheck bool
			}{
				AgeMin:          22,
				AgeMax:          60,
				AllowedRegions:  []string{"北京", "上海"},
				HasCar:          nil,
				HasSocial:       &hasSocialTrue,
				NeedRemoteCheck: false,
			},
		},
	}, nil
}

// GetProductByID 实现顶级接口方法
func (m *MockProductRepo) GetProductByID(id string) (*model.Product, error) {
	products, _ := m.ListAllProducts()
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, errors.New("product not found")
}
