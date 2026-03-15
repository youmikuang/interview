package config

import (
	"encoding/json"
	"os"
	"sync"
)

type ProductFilterRulesJSON struct {
	AgeMin          int      `json:"ageMin"`
	AgeMax          int      `json:"ageMax"`
	AllowedRegions  []string `json:"allowedRegions"`
	HasCar          *bool    `json:"hasCar"`
	HasSocial       *bool    `json:"hasSocial"`
	NeedRemoteCheck bool     `json:"needRemoteCheck"`
}

type ProductJSON struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	FilterRules ProductFilterRulesJSON `json:"filterRules"`
}

type ChannelFilterRulesJSON struct {
	AllowedProductIDs []string `json:"allowedProductIDs"`
	UserAgeMin        int      `json:"userAgeMin"`
	UserAgeMax        int      `json:"userAgeMax"`
	AllowedRegions    []string `json:"allowedRegions"`
	HasCarRequired    *bool    `json:"hasCarRequired"`
	HasHouseRequired  *bool    `json:"hasHouseRequired"`
}

type ChannelJSON struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	FilterRules ChannelFilterRulesJSON `json:"filterRules"`
}

type ConfigData struct {
	Products []ProductJSON `json:"products"`
	Channels []ChannelJSON `json:"channels"`
}

var (
	cfg     *ConfigData
	once    sync.Once
	loadErr error
)

func LoadConfig() (*ConfigData, error) {
	once.Do(func() {
		data, err := os.ReadFile("config/config.json")
		if err != nil {
			loadErr = err
			return
		}
		var c ConfigData
		if err := json.Unmarshal(data, &c); err != nil {
			loadErr = err
			return
		}
		cfg = &c
	})
	return cfg, loadErr
}
