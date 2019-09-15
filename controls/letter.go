package controls

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"sofuny/models"
	"sofuny/utils"
	"time"
)

// 信件
var letter models.Letter

// find
func FindLetter(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}

// create
func CreateLetter(ctx echo.Context) error {
	if err := ctx.Bind(&letter); err != nil {
		return ctx.JSON(500, utils.Response{
			StatusCode: 200,
			Msg:        "",
			Data:       nil,
			Total:      0,
			Time:       time.Now().Local(),
		})
	} else {
		fmt.Println(letter)
		letter.Uuid = xid.New().String()
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Msg:        "",
			Data:       letter,
			Total:      0,
			Time:       time.Now().Local(),
		})
	}
}

// update
func UpdateLetter(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}

// delete
func DeleteLetter(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}
