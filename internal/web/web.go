package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/storage"
	"github.com/masl/undershorts/internal/utils"
)

type WebServer struct {
	store storage.Storage
}

func NewWebServer(store storage.Storage) *WebServer {
	return &WebServer{
		store: store,
	}
}

func (w *WebServer) Serve() error {
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
			v1.GET("/health", w.GetHealth)
			v1.GET("/path/:path", w.GetPath)
			v1.POST("/shorten", w.PostShorten)
		}
	}

	// Route auth endpoints
	auth := router.Group("/auth")
	{
		auth.POST("/signup", w.PostSignup)
		auth.POST("/login", w.PostLogin)
	}

	// Route for testing authentication
	router.GET("/test", w.Auth, func(ctx *gin.Context) {
		ctx.Writer.WriteString("Hello, World!")
		return
	})

	webAddress := utils.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
	return router.Run(webAddress)
}
