package api

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"yellowgreenorgreenyellow/config"
)

func Serve() {

	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	err := engine.SetTrustedProxies(config.TrustedProxies())
	if err != nil {
		panic(err)
	}
	//serve static files from dist
	engine.Use(static.Serve("/", static.LocalFile("./dist", false)))

	//Handle router for svelte (any unknown page brings back to /)
	engine.NoRoute(func(c *gin.Context) {
		log.WithField("path", c.Request.URL.Path).Debug("calling unknown path")
		c.File("./dist/index.html")
	})

	// create an apiGroup (prefix for every route '/api')
	// If you change this, think about changing it as well in vite.config.js and everywhere the api is called if it is hardcoded.
	apiGroup := engine.Group("/api")

	registerRoutes(apiGroup)

	log.WithFields(log.Fields{"address": config.HostAndPort()}).Info("Starting server")
	engine.Run(config.HostAndPort())
}

func registerRoutes(apiGroup *gin.RouterGroup) {
	// You can separate your routes in sub folders, feel free to organize as you want
	apiGroup.GET("/", func(c *gin.Context) {
		id, _ := uuid.NewUUID()
		c.JSON(http.StatusOK, gin.H{"uuid": id.String()})
	})
	apiGroup.POST("/store", Store)
	apiGroup.GET("/stats", Stats)

}
