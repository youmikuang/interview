package application

import (
	"errors"
	_ "product-matching/api/repository"
	_ "product-matching/api/service"
	"product-matching/domain/model"
	domain_service "product-matching/domain/service"
)

// ProductMatchingApp 应用服务
type ProductMatchingApp struct {
	filterService *domain_service.ProductFilter // 注意：区分domain/service和api/service
}

// NewProductMatchingApp 构造函数
func NewProductMatchingApp(filter *domain_service.ProductFilter) *ProductMatchingApp {
	return &ProductMatchingApp{filterService: filter}
}

// MatchProducts 编排匹配流程
func (app *ProductMatchingApp) MatchProducts(user *model.User, channelID string) ([]*model.Product, error) {
	// 前置校验
	if user.Phone == "" {
		return nil, errors.New("user phone is empty")
	}
	if channelID == "" {
		return nil, errors.New("channel ID is empty")
	}

	// 调用领域服务
	return app.filterService.Filter(user, channelID)
}
