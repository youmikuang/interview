package service

// RemoteChecker 远程校验端口（定义手机号MD5校验能力）
type RemoteChecker interface {
	CheckPhoneMD5(md5Str string) (bool, error)
}
