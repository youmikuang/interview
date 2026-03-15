package repository

import "product-matching/domain/model"

// ChannelRepository 渠道仓储端口
type ChannelRepository interface {
	GetChannelByID(id string) (*model.Channel, error)
}
