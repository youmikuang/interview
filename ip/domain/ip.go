package domain

import "context"

type IPAddress string
type IPChannels []string

type Location struct {
	Province string
	City     string
}

type ChannelResult struct {
	ChannelID string
	Location  *Location
	Cost      float32
	Err       error
}

type IPQueryRepository interface {
	Query(ctx context.Context, ip IPAddress, channel string) (*Location, error)
}
