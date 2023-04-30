package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/utils"
)

func Serve() (err error) {
	// router := mux.NewRouter()

	// Sets gin mode
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// Serve static files
	router.Static("/assets", "./web/assets")

	router.LoadHTMLFiles("./web/index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// pathsToUrls := map[string]string{
	// 	"undershorts": "https://github.com/masl/undershorts",
	// 	"author":      "https://github.com/masl",
	// }

	// Routes shortening redirect endpoint
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

	type PostBody struct {
		LongUrl   string `json:"longUrl"`
		ShortPath string `json:"shortPath"`
	}

	// Routes API endpoints
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/shorten", func(ctx *gin.Context) {
				var requestBody PostBody

				if err := ctx.BindJSON(&requestBody); err != nil {
					ctx.Writer.WriteHeader(http.StatusInternalServerError)
					return
				}

				if err := db.SetURL(requestBody.ShortPath, requestBody.LongUrl); err != nil {
					ctx.Writer.WriteHeader(http.StatusInternalServerError)
					return
				}

				ctx.Writer.WriteHeader(http.StatusCreated)
				return
			})
		}
	}

	webAddress := utils.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
	return router.Run(webAddress)

	// mapHandler := handler.MapHandler(pathsToUrls, router)

	// // Redis handler
	// redisContent, err := db.GetAllURLS()
	// if err != nil {
	// 	return
	// }

	// redisHandler, err := handler.RedisHandler(redisContent, mapHandler)
	// if err != nil {
	// 	return
	// }
	/*

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
