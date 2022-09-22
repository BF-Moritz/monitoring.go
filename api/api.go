package main

import (
	"fmt"

	logger "github.com/BF-Moritz/log.lib.go"
	"github.com/BF-Moritz/log.lib.go/enum"
	"github.com/BF-Moritz/log.lib.go/middleware"
	"github.com/BF-Moritz/monitoring.go/api/router"
	"github.com/BF-Moritz/monitoring.go/api/vars"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verboseLevel = kingpin.Flag("verbose", "the level of verbosity (0-3)").Default("1").Short('v').Uint32()
	configFile   = kingpin.Flag("config", "config file").Required().Short('c').File()
)

func main() {

	// --- init kingpin

	kingpin.Parse()
	if *verboseLevel > 3 || *verboseLevel < 0 {
		kingpin.FatalUsage("the level of verbosity has to be between 0 and 3")
	}

	// --- init logger

	vars.Logger = logger.NewLogger(enum.LogLevel(*verboseLevel), true, true)

	// --- init config

	if configFile == nil {
		vars.Logger.LogFatal("main", "missing config file")
	}

	viper.SetConfigType("yaml")
	viper.ReadConfig(*configFile)
	err := viper.Unmarshal(&vars.Config)
	if err != nil {
		vars.Logger.LogFatal("main", "failed to read config (%s)", err)
	}

	// --- init influxDB

	influxClient := influxdb2.NewClient(vars.Config.Influx.URL, vars.Config.Influx.APIKey)
	defer influxClient.Close()
	vars.QueryAPI = influxClient.QueryAPI(vars.Config.Influx.Org)

	// --- init echo

	e := echo.New()
	e.Pre(echoMiddleware.AddTrailingSlash())
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.MakeEchoMiddleware(vars.Logger))

	router.MakeRoutes(e)

	vars.Logger.LogInfo("main", "start listening on port %d", vars.Config.Port)
	err = e.Start(fmt.Sprintf(":%d", vars.Config.Port))
	if err != nil {
		vars.Logger.LogFatal("main", "failed to listen on Port %d (%s)", vars.Config.Port, err)
	}
}
