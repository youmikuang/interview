package model

import (
	"crypto/md5"
	"encoding/hex"
)

// User 用户实体
type User struct {
	Phone     string // 手机号（核心标识）
	Name      string // 姓名
	Age       int    // 年龄
	Gender    string // 性别
	Region    string // 地区（如北京、天津）
	HasHouse  bool   // 有无房产
	HasCar    bool   // 有无车辆
	HasSocial bool   // 有无社保
}

// GetPhoneMD5 获取手机号的MD5值
func (u *User) GetPhoneMD5() string {
	h := md5.New()
	h.Write([]byte(u.Phone))
	return hex.EncodeToString(h.Sum(nil))
}
