package service

import "product-matching/api/service"

// RemoteAPI 远程校验适配器（实现顶级api/service接口）
type RemoteAPI struct {
	APIURL string
}

// 编译期校验
var _ service.RemoteChecker = (*RemoteAPI)(nil)

// NewRemoteAPI 构造函数
func NewRemoteAPI(apiURL string) *RemoteAPI {
	return &RemoteAPI{APIURL: apiURL}
}

// CheckPhoneMD5 实现顶级接口方法
func (api *RemoteAPI) CheckPhoneMD5(md5Str string) (bool, error) {
	// 模拟远程API逻辑
	if len(md5Str) > 0 && md5Str[:6] == "e10adc" {
		return true, nil
	}
	return false, nil
}
