package model

// Channel 渠道实体：存储渠道的筛选配置
type Channel struct {
	ID          string   // 渠道ID
	Name        string   // 渠道名称
	FilterRules struct { // 渠道筛选规则：既筛产品也筛用户
		AllowedProductIDs []string // 渠道允许的产品ID列表
		UserAgeMin        int      // 用户最小年龄
		UserAgeMax        int      // 用户最大年龄
		AllowedRegions    []string // 允许的用户地区列表
		HasCarRequired    *bool    // 是否要求有车辆（nil表示不限制）
		HasHouseRequired  *bool    // 是否要求有房产（nil表示不限制）
	}
}
