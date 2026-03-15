package service

import (
	"product-matching/api/repository" // 依赖顶级接口层
	"product-matching/api/service"    // 依赖顶级接口层
	"product-matching/domain/model"
)

// ProductFilter 产品筛选领域服务（仅依赖顶级接口层）
type ProductFilter struct {
	remoteChecker service.RemoteChecker        // 依赖顶级远程校验接口
	productRepo   repository.ProductRepository // 依赖顶级产品仓储接口
	channelRepo   repository.ChannelRepository // 依赖顶级渠道仓储接口
}

// NewProductFilter 构造函数：注入接口实现（依赖倒置）
func NewProductFilter(
	checker service.RemoteChecker,
	productRepo repository.ProductRepository,
	channelRepo repository.ChannelRepository,
) *ProductFilter {
	return &ProductFilter{
		remoteChecker: checker,
		productRepo:   productRepo,
		channelRepo:   channelRepo,
	}
}

// Filter 核心筛选逻辑（无变化，仅依赖的接口路径改变）
func (f *ProductFilter) Filter(user *model.User, channelID string) ([]*model.Product, error) {
	// 步骤1：获取渠道信息
	channel, err := f.channelRepo.GetChannelByID(channelID)
	if err != nil {
		return nil, err
	}

	// 步骤2：获取所有产品
	allProducts, err := f.productRepo.ListAllProducts()
	if err != nil {
		return nil, err
	}

	// 步骤3：多层筛选
	var matchedProducts []*model.Product
	for _, product := range allProducts {
		if !f.matchChannelProductRule(product, channel) {
			continue
		}
		if !f.matchChannelUserRule(user, channel) {
			continue
		}
		if !f.matchProductRule(user, product) {
			continue
		}
		if product.FilterRules.NeedRemoteCheck {
			ok, err := f.remoteChecker.CheckPhoneMD5(user.GetPhoneMD5())
			if err != nil || !ok {
				continue
			}
		}
		matchedProducts = append(matchedProducts, product)
	}

	return matchedProducts, nil
}

// 以下辅助方法（matchChannelProductRule/matchChannelUserRule/matchProductRule）
// 完全复用之前的逻辑，无修改
func (f *ProductFilter) matchChannelProductRule(product *model.Product, channel *model.Channel) bool {
	for _, pid := range channel.FilterRules.AllowedProductIDs {
		if pid == product.ID {
			return true
		}
	}
	return false
}

func (f *ProductFilter) matchChannelUserRule(user *model.User, channel *model.Channel) bool {
	rules := channel.FilterRules
	if user.Age < rules.UserAgeMin || user.Age > rules.UserAgeMax {
		return false
	}
	regionAllowed := false
	for _, r := range rules.AllowedRegions {
		if r == user.Region {
			regionAllowed = true
			break
		}
	}
	if !regionAllowed {
		return false
	}
	if rules.HasCarRequired != nil && *rules.HasCarRequired != user.HasCar {
		return false
	}
	if rules.HasHouseRequired != nil && *rules.HasHouseRequired != user.HasHouse {
		return false
	}
	return true
}

func (f *ProductFilter) matchProductRule(user *model.User, product *model.Product) bool {
	rules := product.FilterRules
	if user.Age < rules.AgeMin || user.Age > rules.AgeMax {
		return false
	}
	regionAllowed := false
	for _, r := range rules.AllowedRegions {
		if r == user.Region {
			regionAllowed = true
			break
		}
	}
	if !regionAllowed {
		return false
	}
	if rules.HasCar != nil && *rules.HasCar != user.HasCar {
		return false
	}
	if rules.HasSocial != nil && *rules.HasSocial != user.HasSocial {
		return false
	}
	return true
}
