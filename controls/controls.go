package controls

import (
	"github.com/upyun/go-sdk/upyun"
	"sofuny/config"
	"sofuny/models"
)

// 数据库
var db = models.Connection()

// 又拍云
var up = upyun.NewUpYun(&upyun.UpYunConfig{
	Bucket:   config.Config().Upyun.Bucket,
	Operator: config.Config().Upyun.Operator,
	Password: config.Config().Upyun.Password,
})
