package controls

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/mojocn/base64Captcha"
	"math/rand"
	"time"
)

// redis 使用
// 验证码使用redis
type customizeRdsStore struct {
	redisClient *redis.Client
}

// set
func (s *customizeRdsStore) Set(id string, value string) {
	err := s.redisClient.Set(id, value, time.Minute*1).Err()
	if err != nil {
		fmt.Println(err, "1")
		panic(err)
	}
}

// get
func (s *customizeRdsStore) Get(id string, clear bool) (value string) {
	val, err := s.redisClient.Get(id).Result()
	if err != nil {
		//fmt.Println(err,"2")
		//panic(err)
		return "false"
	}
	// TODO  验证通过以后 会删除redis存储的id字段　所以会报错 此处需要更加完善的处理方案
	//if clear {
	//	err := s.redisClient.Del(id).Err()
	//	if err != nil {
	//		fmt.Println(err,"3")
	//		panic(err)
	//	}
	//}
	return val
}

type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

// 使用go验证码库验证
type CaptchaJson struct {
	CaptchaId   string `json:"id" form:"id"`
	CaptchaType string `json:"type" form:"type"`
	VerifyValue string `json:"code" form:"code"`
	//ConfigAudio     base64Captcha.ConfigAudio `json:"config_audio"`
	//ConfigCharacter base64Captcha.ConfigCharacter `json:"config_character"`
	//ConfigDigit     base64Captcha.ConfigDigit `json:"config_digit"`
}

//字符,公式,验证码配置
var ConfigCharacter = base64Captcha.ConfigCharacter{
	Height: 40,
	Width:  200,
	//const CaptchaModeNumber:数字,
	// CaptchaModeAlphabet:字母,
	// CaptchaModeArithmetic:算术,
	// CaptchaModeNumberAlphabet:数字字母混合.
	Mode:               base64Captcha.CaptchaModeArithmetic,
	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	//IsShowHollowLine:   true,
	//IsShowNoiseDot:     true,
	//IsShowNoiseText:    true,
	IsShowSlimeLine: false,
	IsShowSineLine:  false,
	CaptchaLen:      6,
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Data       string `json:"data"`
	Msg        string `json:"msg"`
	Id         string `json:"id"`
	Type       string `json:"type"`
}

// 初始化redis
func init() {
	//create redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//init redis store
	customeStore := customizeRdsStore{client}
	base64Captcha.SetCustomStore(&customeStore)
}

// 获取验证码
func GetCaptcha(ctx echo.Context) error {
	captchaType := rand.Intn(4)
	var captchaTypeStr string
	switch captchaType {
	case 0:
		ConfigCharacter.Mode = base64Captcha.CaptchaModeNumberAlphabet
		captchaTypeStr = "NumberAlphabet"
	case 1:
		ConfigCharacter.Mode = base64Captcha.CaptchaModeNumber
		captchaTypeStr = "Number"
	case 2:
		ConfigCharacter.Mode = base64Captcha.CaptchaModeAlphabet
		captchaTypeStr = "Alphabet"
	default:
		ConfigCharacter.Mode = base64Captcha.CaptchaModeArithmetic
		captchaTypeStr = "Arithmetic"
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKey, caps := base64Captcha.GenerateCaptcha("", ConfigCharacter)

	//以base64编码
	base64string := base64Captcha.CaptchaWriteToBase64Encoding(caps)
	response := Response{
		StatusCode: 200,
		Id:         idKey,
		Type:       captchaTypeStr,
		Data:       base64string,
		Msg:        "success"}
	return ctx.JSON(200, response)
}

// 验证验证码
func VerifyCaptcha(ctx echo.Context) error {
	capt := CaptchaJson{}
	if err := ctx.Bind(&capt); err != nil {
		response := Response{
			StatusCode: 500,
			Msg:        "parse data error!"}
		return ctx.JSON(200, response)
	} else {
		//比较图像验证码
		verifyResult := base64Captcha.VerifyCaptcha(capt.CaptchaId, capt.VerifyValue)
		if verifyResult {
			response := Response{
				StatusCode: 200,
				Msg:        "success"}
			return ctx.JSON(200, response)
		} else {
			response := Response{
				StatusCode: 201,
				Msg:        "failed"}
			return ctx.JSON(200, response)
		}
	}
}
