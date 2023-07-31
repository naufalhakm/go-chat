package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalhakm/go-chat/internal/user"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
}

func Start(addr string) error {
	return r.Run(addr)
}
