package controls

import (
	"github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
	"regexp"
	"sofuny/config"
	"sofuny/utils"
	"time"
)

// auth
var github = config.Config().Auth.Github
type AuthType struct {
	Type string `json:"type"`
}
func GetAuth(ctx echo.Context) error {
	var auth AuthType
	if err := ctx.Bind(&auth);err != nil {

	}
	//github.State = xid.New().String()
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Data: github,
	})
}

// auth login
//获取登录地址
type UserLoginInfo struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	HtmlUrl   string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
}
func AuthLogin(ctx echo.Context) error {
	var user UserLoginInfo
	appCode := ctx.QueryParam("code")
	_, resp, err := fasthttp.Get(nil,
		"https://github.com/login/oauth/access_token?client_id="+github.AppID+
			"&client_secret="+github.AppKey+"&code="+appCode)
	if err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:err,
			Time: time.Now().Local(),
		})
	}
	isMatch, _ := regexp.MatchString("error", string(resp))
	if isMatch {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:"error",
			Time: time.Now().Local(),
		})
	}else {
			re, _ := regexp.Compile("access_token=(.*)&scope")
			newStr := re.FindStringSubmatch(string(resp))
			if len(newStr) >= 2 {
				_, userInfo, err := fasthttp.Get(nil, "https://api.github.com/user?access_token="+newStr[1])
				if err != nil {
					return ctx.JSON(200, utils.Response{
						StatusCode: 201,
						Msg:err,
						Time: time.Now().Local(),
					})
				}
				var json = jsoniter.ConfigCompatibleWithStandardLibrary
				if err := json.Unmarshal(userInfo, &user);err != nil{
					return ctx.JSON(200, utils.Response{
						StatusCode: 201,
						Msg:err,
						Time: time.Now().Local(),
					})
				}
			}
		}
	return ctx.Render(200,"result", user)
}