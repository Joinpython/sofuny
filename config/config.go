package config

import (
	"encoding/json"
	"io/ioutil"
)

var EMAIL_FROM = "超有趣码 <tqblogs@163.com>"
var EMAIL_SUBJECT = "给自己的一封信"
var EMAIL_USER = "tqblogs"
var EMAIL_PASS = "admin123456789"
var EMAIL_HOST = "smtp.163.com"
var EMAIL_ADDR = "smtp.163.com:25"

//定义配置文件解析后的结构
// 数据库配置
type Database struct {
	Driver   string `json:"driver"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// auth 配置
type Auth struct {
	Github Github `json:"github"`
}

// auth github 配置
type Github struct {
	State  string `json:"state"`
	AppID  string `json:"app_id"`
	AppKey string `json:"app_key"`
}

// 又拍云 配置
type UpyunConfig struct {
	Bucket   string `json:"bucket"`
	Operator string `json:"operator"`
	Password string `json:"password"`
	FileUrl  string `json:"fileUrl"`
}

// submail 配置
// 邮件
type SubmailEmail struct {
	AppID  string `json:"app_id"`
	AppKey string `json:"app_key"`
	Reply  string `json:"reply"`
}

// 短信
type SubmailSms struct {
	AppID  string `json:"app_id"`
	AppKey string `json:"app_key"`
}

// 赛邮云　主要配置
type SubmailConfig struct {
	Email SubmailEmail `json:"email"`
	Sms   SubmailSms   `json:"sms"`
}

// 基本认证
type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

// 配置
type ServerConfig struct {
	Name      string        `json:"name"`
	Api       string        `json:"api"`
	StaticApi string        `json:"staticApi"`
	Author    string        `json:"author"`
	Category  []string      `json:"category"`
	Database  Database      `json:"database"`
	Auth      Auth          `json:"auth"`
	Upyun     UpyunConfig   `json:"upyun"`
	Submail   SubmailConfig `json:"submail"`
	BasicAuth BasicAuth     `json:"basicAuth"`
}

type JsonStruct struct{}

func NewJsonStruct() *JsonStruct { return &JsonStruct{} }

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

// 返回一个config 配置
func Config() ServerConfig {
	jsonParse := NewJsonStruct()
	config := ServerConfig{}
	jsonParse.Load("./config.json", &config)
	return config
}
