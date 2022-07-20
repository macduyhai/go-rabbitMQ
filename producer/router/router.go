package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/macduyhai/go-rabbitMQ/producer/config"
	"github.com/macduyhai/go-rabbitMQ/producer/controller"
	"github.com/macduyhai/go-rabbitMQ/producer/middleware"
)

type Router struct {
	config *config.Config
}

var (
	router Router
	one    sync.Once
)

func NewRouter(cfg *config.Config) Router {
	one.Do(func() {
		router.config = cfg
	})

	return router
}
func (router *Router) InitGin() (*gin.Engine, error) {
	controller := controller.NewController(router.config)
	engine := gin.Default()

	engine.Use(middleware.CORSMiddleware())
	engine.GET("/ping", controller.Ping)

	userAuthenMiddleWare := middleware.CheckAPIkey{Apikey: router.config.APIKEY}
	{
		user := engine.Group("/api/v1/user")
		user.Use(userAuthenMiddleWare.Check)
		user.POST("/buy", controller.UserBuy)
	}

	return engine, nil
}
