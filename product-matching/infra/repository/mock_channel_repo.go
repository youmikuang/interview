package repository

import (
	"errors"
	"product-matching/api/repository"
	"product-matching/domain/model"
)

// MockChannelRepo 渠道仓储适配器（实现顶级api/repository接口）
type MockChannelRepo struct{}

// 编译期校验
var _ repository.ChannelRepository = (*MockChannelRepo)(nil)

// GetChannelByID 实现顶级接口方法
func (m *MockChannelRepo) GetChannelByID(id string) (*model.Channel, error) {
	hasHouseTrue := true
	if id == "C001" {
		return &model.Channel{
			ID:   "C001",
			Name: "微信渠道",
			FilterRules: struct {
				AllowedProductIDs []string
				UserAgeMin        int
				UserAgeMax        int
				AllowedRegions    []string
				HasCarRequired    *bool
				HasHouseRequired  *bool
			}{
				AllowedProductIDs: []string{"P001", "P002"},
				UserAgeMin:        20,
				UserAgeMax:        50,
				AllowedRegions:    []string{"北京", "天津"},
				HasCarRequired:    nil,
				HasHouseRequired:  &hasHouseTrue,
			},
		}, nil
	}
	return nil, errors.New("channel not found")
}
