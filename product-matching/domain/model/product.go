package model

// Product 产品实体：存储产品的筛选配置
type Product struct {
	ID          string   // 产品ID
	Name        string   // 产品名称
	FilterRules struct { // 产品自身筛选规则
		AgeMin          int      // 最小年龄
		AgeMax          int      // 最大年龄
		AllowedRegions  []string // 允许的地区
		HasCar          *bool    // 是否要求有车（nil不限制）
		HasSocial       *bool    // 是否要求有社保（nil不限制）
		NeedRemoteCheck bool     // 是否需要远程API（手机号MD5）校验
	}
}
