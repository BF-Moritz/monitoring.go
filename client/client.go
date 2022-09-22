package main

import (
	loglibgo "github.com/BF-Moritz/log.lib.go"
	"github.com/BF-Moritz/log.lib.go/enum"
	"github.com/BF-Moritz/monitoring.go/client/services/stats"
	"github.com/BF-Moritz/monitoring.go/client/vars"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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

	vars.Logger = loglibgo.NewLogger(enum.LogLevel(*verboseLevel), true, true)

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

	// --- connect to server

	vars.Logger.LogInfo("main", "connecting to gRPC server on %s", vars.Config.ServerURL)
	conn, err := grpc.Dial(vars.Config.ServerURL, grpc.WithInsecure())
	if err != nil {
		vars.Logger.LogFatal("main", "failed to connect to %s: %v", vars.Config.ServerURL, err)
	}

	defer conn.Close()

	// start stats service

	m := stats.NewStatsModule(conn, vars.Config.Name)
	m.Init()
}
