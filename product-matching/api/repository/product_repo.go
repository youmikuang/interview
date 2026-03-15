package repository

import "product-matching/domain/model"

// ProductRepository 产品仓储端口（顶级接口，定义数据访问能力）
type ProductRepository interface {
	ListAllProducts() ([]*model.Product, error)
	GetProductByID(id string) (*model.Product, error)
}
