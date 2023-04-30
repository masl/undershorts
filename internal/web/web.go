package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/utils"
)

func Serve() (err error) {
	// router := mux.NewRouter()

	// Set gin mode
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// Serve static files
	router.Static("/assets", "./web/assets")

	router.LoadHTMLFiles("./web/index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	pathsToUrls := map[string]string{
		"undershorts": "https://github.com/masl/undershorts",
		"author":      "https://github.com/masl",
	}

	router.GET("/:path", func(ctx *gin.Context) {
		path := ctx.Param("path")

		for key := range pathsToUrls {
			if key == path {
				ctx.Redirect(http.StatusFound, pathsToUrls[key])
			}
		}
	})

	webAddress := utils.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
	return router.Run(webAddress)

	/*
		mapHandler := handler.MapHandler(pathsToUrls, router)

		// Redis handler
		redisContent, err := db.GetAllURLS()
		if err != nil {
			return
		}

		redisHandler, err := handler.RedisHandler(redisContent, mapHandler)
		if err != nil {
			return
		}

		// API handler
		apiRouter := router.PathPrefix("/api").Subrouter()

		// Register API Endpoints
		api.HealthCheckEndpoint(apiRouter)
		api.PathEndpoint(apiRouter)
		api.ShortenEndpoint(apiRouter, router)

		// Start http server
		webAddress := db.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
		srv := &http.Server{
			Handler: redisHandler,
			Addr:    webAddress,
		}

		log.Println("Starting web server on", webAddress)
		return srv.ListenAndServe()
	*/
}
