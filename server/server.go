package main

import (
	"fmt"
	"net"

	loglibgo "github.com/BF-Moritz/log.lib.go"
	"github.com/BF-Moritz/log.lib.go/enum"
	api "github.com/BF-Moritz/monitoring.go/grpc_api"
	"github.com/BF-Moritz/monitoring.go/server/services/monitoring"
	"github.com/BF-Moritz/monitoring.go/server/vars"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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

	// --- init influxDB

	influxClient := influxdb2.NewClient(vars.Config.Influx.URL, vars.Config.Influx.APIKey)
	defer influxClient.Close()
	writeAPI := influxClient.WriteAPIBlocking("bf_moritz", "monitoring")

	// --- start grpc Server

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", vars.Config.Port))
	if err != nil {
		vars.Logger.LogFatal("main", "failed to listen on port %d (%v)", vars.Config.Port, err)
	}

	s := monitoring.Service{
		InfluxWriteAPI: writeAPI,
	}
	grpcServer := grpc.NewServer()

	api.RegisterChatServiceServer(grpcServer, &s)

	vars.Logger.LogInfo("main", "starting gRPC server on port %d", vars.Config.Port)
	err = grpcServer.Serve(lis)
	if err != nil {
		vars.Logger.LogFatal("main", "failed to serve gRPC server over port %d: %v", vars.Config.Port, err)
	}

}
