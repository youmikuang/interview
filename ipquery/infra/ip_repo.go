package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"ipquery/domain"
	"net/http"
	"time"
)

type HTTPRepo struct {
	client *http.Client
	urls   map[string]string
}

func NewHTTPRepo() *HTTPRepo {
	return &HTTPRepo{
		client: &http.Client{},
		urls: map[string]string{
			"A": "https://ip.useragentinfo.com/json?ip=%s",
			"B": "http://ip-b.com/%s",
			"C": "http://ip-api.com/json/%s?lang=zh-CN",
			"D": "https://api.ipinfo.io/lite/%s?token=ce146ea4171769",
		},
	}
}

func (r *HTTPRepo) Query(ctx context.Context, ip domain.IPAddress, channel string) (*domain.Location, error) {
	url, ok := r.urls[channel]
	if !ok {
		return nil, fmt.Errorf("invalid channel: %s", channel)
	}

	finalURL := fmt.Sprintf(url, ip)
	req, _ := http.NewRequestWithContext(ctx, "GET", finalURL, nil)

	// 测试超时
	if channel == "A" {
		select {
		case <-time.After(6 * time.Second):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	var loc domain.Location
	var result map[string]interface{}
	e := json.NewDecoder(resp.Body).Decode(&result)
	if e != nil {
		return nil, e
	}

	loc = HandleTransData(channel, result, loc)
	return &loc, nil
}

func HandleTransData(channel string, result map[string]interface{}, loc domain.Location) domain.Location {
	loc.Province = "未知"
	loc.City = "未知"

	switch channel {
	case "C":
		if v, ok := result["regionName"].(string); ok && v != "" {
			loc.Province = v
		}
		if v, ok := result["city"].(string); ok && v != "" {
			loc.City = v
		}
	case "D":
		if v, ok := result["country"].(string); ok && v != "" {
			loc.Province = v
			loc.City = v
		}
	case "A":
	case "B":
	default:
	}

	return loc
}
