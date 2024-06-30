package main

import (
	"blog/controller"
	"blog/cron"
	"blog/middleware/cors"
	"blog/middleware/metrics"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func StartCronTask() {
	manager := cron.NewCronManager()
	manager.EquipmentTask(cron.NewAccessConsumerCron(), cron.NewLikeConsumerCron())
	go manager.Run()
}

func main() {
	//starting the cron task
	StartCronTask()
	//registe  gin router
	engine := gin.Default()
	engine.Use(cors.CORS())
	engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	engine.Use(gin.Recovery())
	//绑定流量监控
	metrics.BindMetrics(engine)
	//绑定业务路由
	bindArticleRoutes(engine)
	bindLikeRoutes(engine)
	bindCommentRoutes(engine)
	bindUserRoutes(engine)
	bindRewardRoutes(engine)
	bindTagRoutes(engine)
	engine.Run(":8080")
}

func bindArticleRoutes(engine *gin.Engine) {
	route := engine.Group("/article")

	route.POST("/image/upload", func(context *gin.Context) {
		controller.GetArticle().UploadImage(context)
	})
	route.GET("/image/download", func(ctx *gin.Context) {
		controller.GetArticle().DownloadImage(ctx)
	})
	route.POST("/publish", func(ctx *gin.Context) {
		controller.GetArticle().PublishArticle(ctx)
	})
	route.GET("/findbymaxaccess", func(ctx *gin.Context) {
		controller.GetArticle().FindArticleByMaxAccessNum(ctx)
	})
	route.GET("/findbycreatetime", func(ctx *gin.Context) {
		controller.GetArticle().FindArticleByCreateTime(ctx)
	})
	route.GET("/find", func(ctx *gin.Context) {
		controller.GetArticle().FindArticle(ctx)
	})
	route.GET("/search", func(ctx *gin.Context) {
		controller.GetArticle().SearchArticleByPage(ctx)
	})
}

func bindLikeRoutes(engine *gin.Engine) {
	router := engine.Group("/like")
	router.GET("/confirm", func(ctx *gin.Context) {
		controller.GetLike().SetAsLike(ctx)
	})
	router.GET("/cancel", func(ctx *gin.Context) {
		controller.GetLike().CancelLike(ctx)
	})
}
func bindCommentRoutes(engine *gin.Engine) {
	_ = engine.Group("/comment")
	{

	}
}

func bindUserRoutes(engine *gin.Engine) {
	_ = engine.Group("/user")
	{

	}
}

func bindRewardRoutes(engine *gin.Engine) {
	_ = engine.Group("/reward")
	{

	}
}

func bindTagRoutes(engine *gin.Engine) {
	router := engine.Group("/tag")
	router.GET("/findall", func(ctx *gin.Context) {
		controller.GetTag().GetAllTags(ctx)
	})
}
