package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"sofuny/config"
	"sofuny/controls"
	"sofuny/models"
)

func init() {
	var db = models.Connection()
	// 自动迁移数据库
	if err := db.AutoMigrate(
		&models.Letter{},           // 信件
		&models.User{},             // 用户
		&models.Bookmark{},         // 书签
		&models.Article{},          // 文章
		&models.Subscribe{},        // 订阅
		&models.Comment{},          // 评论
		&models.BigCategory{},      // 书签　大分类
		&models.BookmarkCategory{}, // 书签　小分类
	).Error; err != nil {
		fmt.Print(err)
	} else {
		user := &models.User{Name: "sofuny", Password: "admin123456", IsSuper: true}
		db.Create(&user)
		fmt.Print("迁移成功!")
	}
}

// 模板
type Template struct{ templates *template.Template }

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	app := echo.New()

	// 模板
	temp := &Template{templates: template.Must(template.ParseGlob("templates/*.html"))}
	app.Renderer = temp

	// pre
	app.Pre(middleware.AddTrailingSlash())
	// Middleware
	app.Use(middleware.Logger())    // logging
	app.Use(middleware.Recover())   // recover
	app.Use(middleware.Gzip())      // gzip
	app.Use(middleware.RequestID()) // request id
	app.Use(middleware.Secure())    // 安全
	app.Use(middleware.CORS())      // cors
	//app.Use(middleware.CSRF()) // csrf

	// 基本认证
	//app.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
	//	if username == config.Config().BasicAuth.Username && password == config.Config().BasicAuth.Password{
	//		return true, nil
	//	}
	//	return false, nil
	//}))

	// Routes
	app.GET("/", hello)

	// key 验证
	app.Use(middleware.KeyAuth(func(key string, ctx echo.Context) (bool, error) {
		if key != config.Config().BasicAuth.Key {
			return false, nil
		}
		return true, nil
	}))

	// 接口 v1 版
	v1 := app.Group("/api/v1/")
	// 验证码
	v1.GET("captcha/", controls.GetCaptcha)
	v1.POST("captcha/", controls.VerifyCaptcha)

	// 信件
	v1.GET("letter/", controls.FindLetter)
	v1.POST("letter/", controls.CreateLetter)
	v1.PUT("letter/", controls.UpdateLetter)
	v1.DELETE("letter/", controls.DeleteLetter)

	// 书签
	// 获取书签页面所需所有数据
	v1.GET("bookmark/allData/", controls.GetAllDataForBookmark)
	v1.GET("bookmark/", controls.FindBookmark)
	v1.POST("bookmark/", controls.CreateBookmark)
	v1.PUT("bookmark/", controls.UpdateBookmark)
	v1.DELETE("bookmark/", controls.DeleteBookmark)
	// 书签 大分类
	v1.POST("bookmark/big/", controls.CreateBigCate)
	v1.PUT("bookmark/big/", controls.UpdateBigCate)
	v1.DELETE("bookmark/big/", controls.DeleteBigCate)

	// 书签　小分类
	v1.POST("bookmark/cate/", controls.CreateCate)
	v1.PUT("bookmark/cate/", controls.UpdateCate)
	v1.DELETE("bookmark/cate/", controls.DeleteCate)

	// 用户
	v1.GET("user/", controls.FindUser)
	v1.POST("user/", controls.CreateUser)
	v1.PUT("user/", controls.UpdateUser)
	v1.DELETE("user/", controls.DeleteUser)
	v1.POST("user/login/", controls.UserLogin)

	// 上传图片
	v1.POST("articleFiles/", controls.CreateArticleFiles)
	v1.DELETE("articleFiles/", controls.DeleteArticleFiles)

	// 文件
	v1.POST("file/", controls.UploadFile)
	v1.DELETE("file/", controls.DropFile)

	// 文章
	v1.GET("article/", controls.FindArticle)
	v1.POST("article/", controls.CreateArticle)
	v1.PUT("article/", controls.ChangeArticle)
	v1.DELETE("article/", controls.DeleteArticle)
	v1.POST("article/like/", controls.LikeArticle)

	// 文章订阅
	v1.POST("subscribe/", controls.CreateSubscribe)

	// 基本设置
	v1.GET("setting/", controls.SettingConfig)

	// 分类
	v1.GET("category/", controls.FindCategory)

	// auth
	v1.POST("auth/", controls.GetAuth)
	v1.GET("auth/login/", controls.AuthLogin)

	// 评论
	v1.GET("comment/", controls.FindComment)
	v1.POST("comment/", controls.CreateComment)
	v1.POST("comment/like/", controls.LikeComment)
	v1.POST("comment/dislike/", controls.DislikeComment)

	// Start server
	app.Logger.Fatal(app.Start(":1323"))
}

// Handler
func hello(ctx echo.Context) error {
	//msg := utils.SendMailBy163([]string{"sofuny@aliyun.com"},
	//	[]byte(""+
	//		"<h2 style='color:red'> 嗨，新朋友!</h2>"+
	//		"<p>相信自己，你可以做的很好的，别迷茫了!</p>"+
	//		""), config.EMAIL_SUBJECT)
	//fmt.Println(msg)

	return ctx.String(200, "email send success >>> Hello, World!")
}
