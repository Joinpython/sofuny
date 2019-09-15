package controls

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"sofuny/models"
	"sofuny/utils"
	"time"
)

// 后台　书签页面所需数据
type AllDataForBookmark struct {
	BigCates  []models.BigCategory      `json:"big_cates"`
	Cates     []models.BookmarkCategory `json:"cates"`
	Bookmarks []models.Bookmark         `json:"bookmarks"`
}

func GetAllDataForBookmark(ctx echo.Context) error {
	var allData AllDataForBookmark
	if err := db.Find(&allData.BigCates).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Find(&allData.Cates).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Find(&allData.Bookmarks).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       allData,
		Time:       time.Now().Local(),
	})
}

// 书签　大分类
// create
func CreateBigCate(ctx echo.Context) error {
	var big models.BigCategory
	if err := ctx.Bind(&big); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	big.Uuid = xid.New().String()
	if err := db.Create(&big).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       big,
		Time:       time.Now().Local(),
	})
}

// update
func UpdateBigCate(ctx echo.Context) error {
	var big models.BigCategory
	var oldBig models.BigCategory
	if err := ctx.Bind(&big); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Limit(1).Where("id=?", big.ID).Find(&oldBig).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if big.Name == oldBig.Name && big.Status == oldBig.Status {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        "大分类 " + oldBig.Name + " 未做修改",
			Time:       time.Now().Local(),
		})
	}
	big.UpdatedAt = time.Now().Local()
	if err := db.Model(&models.BigCategory{}).Where("id=?", big.ID).Updates(&big).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       big,
		Time:       time.Now().Local(),
	})
}

// delete
func DeleteBigCate(ctx echo.Context) error {
	var big models.BigCategory
	if err := ctx.Bind(&big); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Where("id=?", big.ID).Unscoped().Delete(&big).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       big,
		Time:       time.Now().Local(),
	})
}

// 书签小分类
// create 小分类
func CreateCate(ctx echo.Context) error {
	var cate models.BookmarkCategory
	if err := ctx.Bind(&cate); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	cate.Uuid = xid.New().String()
	if err := db.Create(&cate).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       cate,
		Time:       time.Now().Local(),
	})
}

// update 小分类
func UpdateCate(ctx echo.Context) error {
	var cate models.BookmarkCategory
	var oldCate models.BookmarkCategory
	if err := ctx.Bind(&cate); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Limit(1).Where("id=?", cate.ID).Find(&oldCate).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if cate.Name == oldCate.Name && cate.Status == oldCate.Status && cate.BigID == oldCate.BigID {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        "小分类 " + oldCate.Name + " 未做修改",
			Time:       time.Now().Local(),
		})
	}
	cate.UpdatedAt = time.Now().Local()
	if err := db.Model(&models.BookmarkCategory{}).Where("id=?", cate.ID).Updates(&cate).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       cate,
		Time:       time.Now().Local(),
	})
}

// delete 小分类
func DeleteCate(ctx echo.Context) error {
	var cate models.BookmarkCategory
	if err := ctx.Bind(&cate); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Where("id=?", cate.ID).Unscoped().Delete(&cate).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       cate,
		Time:       time.Now().Local(),
	})
}

// 书签
// find
func FindBookmark(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}

// create
func CreateBookmark(ctx echo.Context) error {
	var bookmark models.Bookmark
	if err := ctx.Bind(&bookmark); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	} else {
		fmt.Println(bookmark)
		bookmark.Uuid = xid.New().String()
		bookmark.Status = true
		if err := db.Create(&bookmark).Error; err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Time:       time.Now().Local(),
			})
		}
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Msg:        "success",
			Data:       bookmark,
			Total:      0,
			Time:       time.Now().Local(),
		})
	}
}

// update
func UpdateBookmark(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}

// delete
func DeleteBookmark(ctx echo.Context) error {
	return ctx.JSON(200, utils.Response{
		StatusCode: 0,
		Msg:        "",
		Data:       nil,
		Total:      0,
		Time:       time.Now().Local(),
	})
}
