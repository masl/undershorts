package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/utils"
	"github.com/masl/undershorts/internal/web/controllers"
)

func Serve() (err error) {
	// Set gin mode
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// Serve static files
	router.Static("/assets", "./web/assets")

	router.LoadHTMLFiles("./web/index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// Route shortening redirect endpoint
	router.GET("/:path", func(ctx *gin.Context) {
		path := ctx.Param("path")

		if !db.Exist(path) {
			ctx.Writer.WriteHeader(http.StatusNotFound)
			return
		}

		url, err := db.GetURL(path)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx.Redirect(http.StatusFound, url)
	})

	// Route API endpoints
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", controllers.GetHealth)
			// TODO: get path
			v1.POST("/shorten", controllers.PostShorten)
		}
	}

	webAddress := utils.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
	return router.Run(webAddress)
}
