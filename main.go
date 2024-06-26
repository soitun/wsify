package main

import (
	"github.com/alash3al/wsify/broker"
	_ "github.com/alash3al/wsify/broker/drivers/memory"
	_ "github.com/alash3al/wsify/broker/drivers/redis"
	"github.com/alash3al/wsify/config"
	"github.com/alash3al/wsify/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	cfg, err := config.NewFromFlags()
	if err != nil {
		panic(err.Error())
	}

	brokerConn, err := broker.Connect(cfg.GetBrokerDriver(), cfg.GetBrokerDSN())
	if err != nil {
		panic(err.Error())
	}

	srv := echo.New()
	srv.HideBanner = true

	srv.Use(middleware.CORS())
	srv.Use(middleware.Logger())

	srv.GET("/ws", routes.WebsocketRouteHandler(cfg, brokerConn))
	srv.POST("/broadcast", routes.BroadcastHandler(cfg, brokerConn))

	log.Fatal(srv.Start(cfg.GetWebServerListenAddr()))
}
