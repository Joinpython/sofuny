package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"net/smtp"
	"sofuny/config"
	"strconv"
	"time"
)

// 回应
type Response struct {
	StatusCode int         `json:"statusCode"`
	Msg        interface{} `json:"msg"`
	Data       interface{} `json:"data"`
	Total      int         `json:"total"`
	Time       time.Time   `json:"time"`
}

// 发送邮件
func SendMailBy163(to []string, content []byte, sub string) bool {
	e := email.NewEmail()
	e.To = to
	e.From = config.EMAIL_FROM
	e.Subject = sub
	//e.Text = content
	e.HTML = content
	err := e.Send(config.EMAIL_ADDR, smtp.PlainAuth("", config.EMAIL_USER, config.EMAIL_PASS, config.EMAIL_HOST))
	if err != nil {
		return false
	} else {
		return true
	}
}

// 订阅 邮件
func SendSubscribEmail(to []string, title string) error {
	e := email.NewEmail()
	e.To = to
	e.From = config.EMAIL_FROM
	e.Subject = "thanks for your subscripation " + config.Config().Name
	str := fmt.Sprintf("<h2>hello</h2> <p>最新文章:%s</p>", title)
	e.HTML = []byte(str)
	err := e.Send(config.EMAIL_ADDR, smtp.PlainAuth("", config.EMAIL_USER, config.EMAIL_PASS, config.EMAIL_HOST))
	return err
}

// 赛邮云　配置
var emailAppID = config.Config().Submail.Email.AppID
var emailAppkey = config.Config().Submail.Email.AppKey
var smsAppID = config.Config().Submail.Sms.AppID
var smsAppkey = config.Config().Submail.Sms.AppKey

// 获取免费邮件余额
type EmailBalance struct {
	Status      string `json:"status"`
	Balance     int    `json:"balance"`
	FreeBalance int    `json:"free_balance"`
}

func GetSubmailEmailBalance() (error error, freeBalance int) {
	// 请求url
	balance_url := "https://api.submail.cn/balance/mail"
	// 表单填充
	data := &fasthttp.Args{}
	data.Add("appid", emailAppID)
	data.Add("signature", emailAppkey)
	status, resp, err := fasthttp.Post(nil, balance_url, data)
	if err != nil {
		return err, 0
	}
	if status != fasthttp.StatusOK {
		return err, 0
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var balance EmailBalance
	if err := json.Unmarshal(resp, &balance); err != nil {
		return err, 0
	}
	if balance.Status != "success" {
		return error, 0
	}
	return nil, balance.FreeBalance
}

// 赛邮　发送邮件
type SendEmailResponse struct {
	Status string `json:"status"`
	Return []struct {
		SendID string `json:"send_id"`
		To     string `json:"to"`
	} `json:"return"`
}

// TODO 主题和内容　根据需要发送的内容自定义使用
func SendEmailBySubmail(to string, subject string, context string) bool {
	// 发送邮件url
	// 多个联系人用半角“,”隔开：
	// e.g. "leo <leo@submail.cn>, <retro@submail.cn>, service@submail.cn";
	sendEmailUrl := "https://api.submail.cn/mail/send"
	err, freeBalance := GetSubmailEmailBalance()
	if err != nil {
		return false
	}
	if freeBalance <= 0 {
		return false
	}
	data := &fasthttp.Args{}
	data.Add("appid", emailAppID)
	data.Add("signature", emailAppkey)
	data.Add("to", to)
	data.Add("from", "sofuny@sofuny.cn")
	data.Add("from_name", config.Config().Name)
	data.Add("reply", config.Config().Submail.Email.Reply)
	// subject "激活您的"+config.Config().Name+"账号"
	data.Add("subject", subject)
	// context html 文本 "您好，欢迎使用"+config.Config().Name+"账号。为了确保注册信息的真实性。"
	data.Add("html", context)
	data.Add("asynchronous", "true")
	status, resp, err := fasthttp.Post(nil, sendEmailUrl, data)
	if err != nil {
		return false
	}
	if status != fasthttp.StatusOK {
		return false
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var send SendEmailResponse
	if err := json.Unmarshal(resp, &send); err != nil {
		return false
	}
	if send.Status != "success" {
		return false
	}
	return true
}

// 短信

// 查询短信余额
type SmsBalance struct {
	Status               string `json:"status"`
	Balance              string `json:"balance"`
	TransactionalBalance string `json:"transactional_balance"`
}

func GetSubmailSmsBalance() (error error, total int) {
	// 请求url
	balance_url := "https://api.submail.cn/balance/sms"
	// 表单填充
	data := &fasthttp.Args{}
	data.Add("appid", smsAppID)
	data.Add("signature", smsAppkey)
	status, resp, err := fasthttp.Post(nil, balance_url, data)
	if err != nil {
		return err, 0
	}
	if status != fasthttp.StatusOK {
		return err, 0
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var sms SmsBalance
	if err := json.Unmarshal(resp, &sms); err != nil {
		return err, 0
	}
	if sms.Status != "success" {
		return error, 0
	}
	total, err = strconv.Atoi(sms.Balance)
	return nil, total
}

// 发送短信
type SendSmsResponse struct {
	Status                  string `json:"status"`
	SendID                  string `json:"send_id"`
	Fee                     int    `json:"fee"`
	SmsCredits              string `json:"sms_credits"`
	TransactionalSmsCredits string `json:"transactional_sms_credits"`
}

func SendSmsBySubmail(to string) bool {
	// content
	// 正文中必须提交有效的短信签名，且您的短信签名必须放在短信的最前端，e.g.
	// 【SUBMAIL】您的短信验证码：4438，请在10分钟内输入。
	// content 参数将会与您账户中的短信模板进行匹配，
	// 如无匹配 API 会创建一个短信模板并提交审核，如审核失败则返回 420 错误
	// 请将 短信正文控制在 350 个字符以内。

	sendSmslUrl := "https://api.submail.cn/message/send"
	err, total := GetSubmailSmsBalance()
	if err != nil {
		return false
	}
	if total <= 0 {
		return false
	}
	data := &fasthttp.Args{}
	data.Add("appid", smsAppID)
	data.Add("signature", smsAppkey)
	data.Add("to", to)
	data.Add("content", "【SUBMAIL】您的短信验证码：4438，请在10分钟内输入")
	status, resp, err := fasthttp.Post(nil, sendSmslUrl, data)
	if err != nil {
		return false
	}
	if status != fasthttp.StatusOK {
		return false
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var send SendSmsResponse
	if err := json.Unmarshal(resp, &send); err != nil {
		return false
	}
	if send.Status != "success" {
		return false
	}
	return true
}
