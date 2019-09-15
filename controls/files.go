package controls

import (
	"github.com/labstack/echo/v4"
	"github.com/upyun/go-sdk/upyun"
	"os"
	"sofuny/config"
	"sofuny/utils"
	"strconv"
	"strings"
	"time"
)

// 上传文件
func UploadFile(ctx echo.Context) error {
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
	// 判断文件类型
	var filePath = "other"
	fileType := fileNames[len(fileNames)-1]
	switch strings.ToLower(fileType) {
	case "png":
		filePath = "/image/"
		break
	case "jpeg":
		filePath = "/image/"
		break
	case "jpg":
		filePath = "/image/"
		break
	case "gif":
		filePath = "/image/"
		break
	case "svg":
		filePath = "/image/"
		break
	case "webp":
		filePath = "/image/"
		break
	case "mp4":
		filePath = "/video/"
		break
	case "avi":
		filePath = "/video/"
		break
	case "3gp":
		filePath = "/video/"
		break
	case "mov":
		filePath = "/video/"
		break
	case "flv":
		filePath = "/video/"
		break
	case "md":
		filePath = "/markdown/"
		break
	case "css":
		filePath = "/css/"
		break
	case "js":
		filePath = "/js/"
		break
	case "html":
		filePath = "/html/"
		break
	case "txt":
		filePath = "/txt/"
		break
	case "json":
		filePath = "/json/"
		break
	case "vue":
		filePath = "/vue/"
		break
	}
	// 上传 图片
	if err := up.Put(&upyun.PutObjectConfig{
		Path:   filePath + imgName,
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
type Files struct {
	FileUrl string `json:"fileUrl"`
}

func DropFile(ctx echo.Context) error {
	file := Files{}
	if err := ctx.Bind(&file); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err.Error(),
		})
	}
	paths := strings.Split(file.FileUrl, "/")
	currentPath, err := os.Getwd()
	if err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err.Error(),
		})
	}
	err = os.Remove(currentPath + "/static/image/" + paths[len(paths)-1])
	if err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err.Error(),
		})
	} else {
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Msg:        "删除成功!",
			Data:       file.FileUrl,
		})
	}
}
