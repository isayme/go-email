package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-email/app"
	"github.com/isayme/go-email/app/conf"
	"github.com/isayme/go-email/app/manager"
	"github.com/isayme/go-email/app/middleware"
	"github.com/isayme/go-email/app/router"
	logger "github.com/isayme/go-logger"
)

var configPath = flag.String("c", "/etc/email.json", "config file path")
var showVersion = flag.Bool("v", false, "show version")
var showHelp = flag.Bool("h", false, "show help")

func main() {
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("%s: %s\n", app.Name, app.Version)
		os.Exit(0)
	}

	conf.SetPath(*configPath)
	config := conf.Get()

	if config.Logger.Level != "" {
		logger.SetLevel(config.Logger.Level)
	}

	manager.Init(config)

	r := gin.New()
	r.Use(middleware.Logger)
	r.Use(gin.Recovery())

	r.GET("/version", router.Version)
	r.POST("/send", router.Send)
	r.Run(fmt.Sprintf(":%d", config.HTTP.Port))
}
