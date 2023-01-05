package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pokrovsky-io/nats-store/internal/usecase"
	"net/http"
	"strconv"
)

// TODO: Здесь добавим и настроим роутер, а веб-сервер уже в gin

// TODO: Добавить инъекцию логгера в роутер
func NewRouter(router *gin.Engine, uc usecase.Order) {
	// Options
	//handler.Use(gin.Logger())
	//handler.Use(gin.Recovery())

	router.Static("/static", "./web/static")
	router.LoadHTMLFiles("web/index.html")

	router.GET("/", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("order_id"))
		order, _ := uc.Get(id)
		jsonOrder, _ := json.MarshalIndent(order, "", "  ")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"data": string(jsonOrder),
		})
	})
}
