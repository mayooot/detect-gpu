package main

import (
	goflag "flag"

	flag "github.com/spf13/pflag"

	"github.com/mayooot/detect-gpu/internal/stat"
)

var port *string = flag.String("port", ":2376", "Port of detect server, format :port")

var pattern *string = flag.String("pattern", "/api/v1/detect/gpu", "Pattern of detect server")

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	stat := &stat.Stat{}
	stat.Run(*port, *pattern)
}
