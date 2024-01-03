package main

import (
	goflag "flag"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/mayooot/detect-gpu/internal/stat"
)

var (
	port    = flag.StringP("port", "p", ":2376", "Port of detect server, format :port")
	pattern = flag.StringP("pattern", "r", "/api/v1/detect/gpu", "Pattern of detect server")
	td      = flag.DurationP("td", "t", 5*time.Second, "Time duration for detect gpu")
)

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	stat := &stat.Stat{}
	stat.Run(*port, *pattern, *td)
}
