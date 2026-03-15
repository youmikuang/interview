package domain

import "context"

// IPAddress IP 值对象
type IPAddress string

// Location 归属地
type Location struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

type ChannelResult struct {
	ChannelID string
	Location  *Location
	Err       error
}

type IPQueryRepository interface {
	Query(ctx context.Context, ip IPAddress, channel string) (*Location, error)
}
