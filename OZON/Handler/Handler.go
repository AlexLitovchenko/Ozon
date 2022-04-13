package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	cors "github.com/rs/cors/wrapper/gin"
)

type Handle struct {
	rdb *redis.Client
}

func (h *Handle) Handles() {
	h.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})
	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"192.168.1.2"})

	router.POST("/", h.CreateShortURL)
	router.GET("/", h.GetShortURL)

	router.Run()
}
