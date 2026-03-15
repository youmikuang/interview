package application

import (
	"context"
	"sync"
	"time"

	"ipquery/domain"
)

type IPService struct {
	repo     domain.IPQueryRepository
	timeout  time.Duration
	channels []string
}

func NewIPService(repo domain.IPQueryRepository, timeout time.Duration) *IPService {
	return &IPService{
		repo:    repo,
		timeout: timeout,
		channels: []string{
			"A", "B", "C", "D",
		},
	}
}

// BatchQueryIP 并发查询所有渠道
func (s *IPService) BatchQueryIP(ctx context.Context, ip domain.IPAddress) map[string]domain.ChannelResult {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	var wg sync.WaitGroup
	ch := make(chan domain.ChannelResult, len(s.channels))

	for _, channel := range s.channels {
		wg.Add(1)
		go func(c string) {
			defer wg.Done()
			loc, err := s.repo.Query(ctx, ip, c)
			ch <- domain.ChannelResult{
				ChannelID: c,
				Location:  loc,
				Err:       err,
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// 聚合结果
	result := make(map[string]domain.ChannelResult)
	for item := range ch {
		result[item.ChannelID] = item
	}

	return result
}
