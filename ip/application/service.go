package application

import (
	"context"
	"ip/domain"
	"math"
	"sync"
	"time"
)

type IPService struct {
	repo     domain.IPQueryRepository
	timeout  time.Duration
	channels []string
}

func NewIpService(repo domain.IPQueryRepository, timeout time.Duration) *IPService {
	return &IPService{
		repo:    repo,
		timeout: timeout,
		channels: []string{
			"A", "B", "C", "D",
		},
	}
}

func (s *IPService) BatchQueryIP(ctx context.Context, ip domain.IPAddress, channels domain.IPChannels) map[string]domain.ChannelResult {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	var wg sync.WaitGroup
	ch := make(chan domain.ChannelResult, len(channels))
	for _, c := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			loc, err := s.repo.Query(ctx, ip, c)
			ch <- domain.ChannelResult{
				Location:  loc,
				ChannelID: c,
				Cost:      float32(math.Round(time.Since(start).Seconds()*1000) / 1000),
				Err:       err,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := make(map[string]domain.ChannelResult)
	for item := range ch {
		result[item.ChannelID] = item
	}
	return result
}
