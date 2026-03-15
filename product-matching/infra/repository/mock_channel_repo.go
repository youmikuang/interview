package repository

import (
	"errors"
	"fmt"
	"product-matching/api/repository"
	"product-matching/config"
	"product-matching/domain/model"
)

// MockChannelRepo 渠道仓储适配器（实现顶级api/repository接口）
type MockChannelRepo struct{}

// 编译期校验
var _ repository.ChannelRepository = (*MockChannelRepo)(nil)

// GetChannelByID 实现顶级接口方法
func (m *MockChannelRepo) GetChannelByID(id string) (*model.Channel, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	for _, c := range cfg.Channels {
		if c.ID == id {
			channel := &model.Channel{
				ID:   c.ID,
				Name: c.Name,
			}
			channel.FilterRules.AllowedProductIDs = c.FilterRules.AllowedProductIDs
			channel.FilterRules.UserAgeMin = c.FilterRules.UserAgeMin
			channel.FilterRules.UserAgeMax = c.FilterRules.UserAgeMax
			channel.FilterRules.AllowedRegions = c.FilterRules.AllowedRegions
			channel.FilterRules.HasCarRequired = c.FilterRules.HasCarRequired
			channel.FilterRules.HasHouseRequired = c.FilterRules.HasHouseRequired
			return channel, nil
		}
	}
	return nil, errors.New("channel not found")
}