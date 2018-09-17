package main

import (
	"net/http"
	"./config"
	"./router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"fmt"
	"sync"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

var moni map[string]string
var monitorHeart *monitor_heart

type monitor_heart struct{
	heartMap map[string]int64
	sync.RWMutex
}

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	moni = viper.GetStringMapString("monitor")
	fmt.Println("moni :",moni)

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}
	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.

		middlewares...,
	)

	monitorHeart = &monitor_heart{
		heartMap:make(map[string]int64),
	}
	pbs_c = &Pbs{}
	go monitorQuest()
	go runInfluxdb("http://10.0.0.71:8086","zhongxuan","gdjy")

	http.ListenAndServe(viper.GetString("addr"), g)
}



