package controls

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"sofuny/models"
	"sofuny/utils"
	"time"
)

// 信件
var user models.User

// find
func FindUser(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}

// create
func CreateUser(ctx echo.Context) error {
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(500, utils.Response{
			StatusCode: 200,
			Msg:        "",
			Data:       nil,
			Total:      0,
			Time:       time.Now().Local(),
		})
	} else {
		var lastUser models.User
		user.Uuid = xid.New().String()
		user.Status = true
		user.IsSuper = false
		db.Last(&lastUser)
		user.ID = lastUser.ID + 1
		if err := db.Create(&user).Error; err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Data:       nil,
				Total:      0,
				Time:       time.Now().Local(),
			})
		}
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Msg:        "注册成功!",
			Data:       user.Name,
			Total:      0,
			Time:       time.Now().Local(),
		})
	}
}

// update
func UpdateUser(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}

// delete
func DeleteUser(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}

// 用户登录
type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func UserLogin(ctx echo.Context) error {
	var login Login
	if err := ctx.Bind(&login); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err.Error(),
			Time:       time.Now().Local(),
		})
	}
	if err := db.Limit(1).Where("name = ? AND is_super=?", login.Name, true).Find(&user).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        "用户不存在!",
			Time:       time.Now().Local(),
		})
	}
	if login.Password != user.Password {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        "密码错误",
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "登录成功!",
		Data:       user.Name,
		Time:       time.Now().Local(),
	})
}
