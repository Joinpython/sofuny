package controls

import (
	"github.com/labstack/echo/v4"
	"sofuny/config"
	"sofuny/utils"
	"time"
)

func SettingConfig(ctx echo.Context) error {
	var setting = config.Config()
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Data:       setting,
		Time:       time.Now().Local(),
	})
}
