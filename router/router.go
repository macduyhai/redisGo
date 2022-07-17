package router

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/macduyhai/redisGo/models"
	"github.com/macduyhai/redisGo/rdcli"
)

var (
	router *Router
	one    sync.Once
)

type Router struct {
	db       *gorm.DB
	rediscli *rdcli.RedisCli
}

func NewRouter(db *gorm.DB, recli *rdcli.RedisCli) *Router {
	one.Do(func() {
		router = &Router{
			db:       db,
			rediscli: recli,
		}
	})
	return router
}

func (router *Router) InitGin() (*gin.Engine, error) {
	engine := gin.Default()
	engine.GET("/ping", Ping)
	engine.POST("/add", AddUser)
	return engine, nil
}
func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "PongPong"})
}
func AddUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := router.rediscli.SaveRedisDB(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	router.rediscli.LoadAllData()
	ctx.JSON(http.StatusOK, gin.H{"message": "Add success"})
}
