package controls

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/upyun/go-sdk/upyun"
	"sofuny/config"
	"sofuny/models"
	"sofuny/utils"
	"strconv"
	"strings"
	"time"
)

// 文章图片一律上传到第三方云存储平台
func CreateArticleFiles(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 0,
			Msg:        err.Error(),
		})
	}
	fileRead, err := file.Open()
	if err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 0,
			Msg:        err.Error(),
		})
	}
	defer fileRead.Close()
	// 根据时间戳生成图片名称
	timeUnix := time.Now().UnixNano()
	fileNames := strings.Split(file.Filename, ".")
	timeUnixStr := strconv.FormatInt(timeUnix, 10)
	imgName := timeUnixStr + "." + fileNames[len(fileNames)-1]
	// 获取 存储用量
	useCount, _ := up.Usage()
	if useCount > 10485760 {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        "存储空间已满, 不支持上传图片了!",
			Time:       time.Now().Local(),
		})
	}
	// 上传 图片
	if err := up.Put(&upyun.PutObjectConfig{
		Path:   "/image/article/" + imgName,
		Reader: fileRead,
		UseMD5: true,
	}); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "上传成功!",
		Data:       config.Config().Upyun.FileUrl + "/image/article/" + imgName,
		Time:       time.Now().Local(),
	})
}

// 文章 删除图片
type Image struct {
	ImgUrl string `json:"imgUrl"`
}

func DeleteArticleFiles(ctx echo.Context) error {
	image := Image{}
	if err := ctx.Bind(&image); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err.Error(),
			Time:       time.Now().Local(),
		})
	}
	paths := strings.Split(image.ImgUrl, "/")
	if err := up.Delete(&upyun.DeleteObjectConfig{
		Path:  "/image/article/" + paths[len(paths)-1],
		Async: true,
	}); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "删除成功!",
		Data:       image.ImgUrl,
		Time:       time.Now().Local(),
	})
}

// 文章
// 创建文章
func CreateArticle(ctx echo.Context) error {
	var article models.Article
	if err := ctx.Bind(&article); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	} else {
		var lastArticle models.Article
		article.Uuid = xid.New().String()
		article.UserID = 1
		db.Find(&lastArticle)
		article.ID = lastArticle.ID + 1
		if err := db.Create(&article).Error; err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Time:       time.Now().Local(),
			})
		} else {
			var subscribList []models.Subscribe
			var emailList []string
			db.Limit(1).First(&subscribList)
			for i := 0; i < len(subscribList); i++ {
				emailList = append(emailList, subscribList[i].Email)
			}
			subject := config.Config().Name + "发布了新文章(" + article.Title + ")"
			for i := 0; i < len(emailList); i++ {
				context := "<h2>你好，" + emailList[i] + "</h2><p>" +
					config.Config().Author + " | 发布了新文章: <a href='https://www.sofuny.cn/post/" +
					article.Uuid + "'>" + article.Title + "</a></p>"
				// TODO 未解决 for 循环　只发送最后一个邮箱地址的问题 需要加深go语言的学习
				flag := utils.SendEmailBySubmail(emailList[i], subject, context)
				if flag {
					fmt.Println("发送订阅成功!")
				}
			}
			return ctx.JSON(200, utils.Response{
				StatusCode: 200,
				Msg:        "success",
				Data:       article,
				Time:       time.Now().Local(),
			})
		}
	}
}

// 查找文章
// 具体那一篇文章的返回值
type ResponseArticle struct {
	Article       models.Article   `json:"article"`
	NextArticle   models.Article   `json:"nextArticle"`
	RelateArticle []models.Article `json:"relateArticle"`
}

func FindArticle(ctx echo.Context) error {
	var article models.Article
	var nextArticle models.Article
	var relateArticle []models.Article
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")
	uuid := ctx.QueryParam("uuid")
	// 查询单个文章
	if uuid != "" {
		if err := db.Limit(1).Where("uuid=?", uuid).Find(&article).Error; err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Time:       time.Now().Local(),
			})
		}
		article.ViewCounts += 1
		db.Save(&article)
		if err := db.Limit(1).Where("id=?", article.ID+1).Find(&nextArticle).Error; err != nil {
			nextArticle = models.Article{}
		}
		if err := db.Limit(3).Where("category=? AND id !=?", article.Category, article.ID).Find(&relateArticle).Error; err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Time:       time.Now().Local(),
			})
		}
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Time:       time.Now().Local(),
			Data: ResponseArticle{
				Article:       article,
				NextArticle:   nextArticle,
				RelateArticle: relateArticle,
			},
		})
	}
	// 文章列表
	var results []models.Article
	var hotArticle []models.Article
	if err := db.Limit(limit).Offset(offset).Where("status=?", true).Find(&results).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Limit(10).Where("status=? AND view_counts > ? OR like_counts > ? OR comment_counts > ?", true, 100, 10, 5).Find(&hotArticle).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data: map[string][]models.Article{
			"articleList":    results,
			"hotArticleList": hotArticle,
		},
		Total: len(results),
		Time:  time.Now().Local(),
	})
}

// 修改文章
func ChangeArticle(ctx echo.Context) error {
	var article models.Article
	var oldArticle models.Article
	if err := ctx.Bind(&article); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	} else {
		if err := db.Limit(1).Where("id=?", article.ID).First(&oldArticle).Error; err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Time:       time.Now().Local(),
			})
		}
		if article == oldArticle {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        "文章未做修改",
				Time:       time.Now().Local(),
			})
		}
		// 如果原来的媒体地址和现在的不一样 那么就删除原来的媒体地址
		if article.MediaUrl != oldArticle.MediaUrl {
			mediaPath := strings.Split(oldArticle.MediaUrl, "/")
			if err := up.Delete(&upyun.DeleteObjectConfig{
				Path:  "/image/article/" + mediaPath[len(mediaPath)-1],
				Async: true,
			}); err != nil {
				return ctx.JSON(200, utils.Response{
					StatusCode: 201,
					Msg:        err,
					Time:       time.Now().Local(),
				})
			}
		}
		// TODO 未实现改变文章内的图片 之后 删除文章内的图片 有时间再来解决 可能要更换markdown编辑器
		//if article.Images != oldArticle.Images {
		//
		//}

		if err := db.Model(&models.Article{}).Where("id=?", article.ID).Updates(&article).Error; err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Time:       time.Now().Local(),
			})
		}
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Data:       article,
			Msg:        "更新成功!",
			Time:       time.Now().Local(),
		})
	}
}

// 删除文章
func DeleteArticle(ctx echo.Context) error {
	var article models.Article
	if err := ctx.Bind(&article); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	// 删除文章所带文件
	if err := db.Limit(1).Where("id=?", article.ID).First(&article).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	var deleteFiles []string // 需要删除的文件列表
	mediaPath := strings.Split(article.MediaUrl, "/")
	imagePath := strings.Split(article.Images, ",")
	deleteFiles = append(deleteFiles, mediaPath[len(mediaPath)-1])
	for i := 0; i < len(imagePath); i++ {
		items := strings.Split(imagePath[i], "/")
		deleteFiles = append(deleteFiles, items[len(items)-1])
	}
	for i := 0; i < len(deleteFiles); i++ {
		if err := up.Delete(&upyun.DeleteObjectConfig{
			Path:  "/image/article/" + deleteFiles[i],
			Async: true,
		}); err != nil {
			return ctx.JSON(200, utils.Response{
				StatusCode: 201,
				Msg:        err,
				Time:       time.Now().Local(),
			})
		}
	}
	if err := db.Where("id=?", article.ID).Unscoped().Delete(&article).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       article,
		Time:       time.Now().Local(),
	})
}

// 文章 订阅
func CreateSubscribe(ctx echo.Context) error {
	var subscribe models.Subscribe
	var lastSubscribe models.Subscribe
	if err := ctx.Bind(&subscribe); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	db.Find(&lastSubscribe)
	subscribe.Uuid = xid.New().String()
	subscribe.ID = lastSubscribe.ID + 1
	subscribe.Status = true
	if err := db.Create(&subscribe).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Time:       time.Now().Local(),
	})
}

// 文章点赞
type LikeArticleStruct struct {
	ID int `json:"id"`
}

func LikeArticle(ctx echo.Context) error {
	var like LikeArticleStruct
	var article models.Article
	if err := ctx.Bind(&like); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Limit(1).Where("id=?", like.ID).First(&article).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	article.LikeCounts += 1
	db.Save(&article)
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       map[string]int{"newLikeCounts": article.LikeCounts},
		Time:       time.Now().Local(),
	})
}
