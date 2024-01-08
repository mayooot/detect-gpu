package main

import (
	goflag "flag"
	"fmt"
	"syscall"
	"time"

	"github.com/judwhite/go-svc"
	"github.com/ngaut/log"
	flag "github.com/spf13/pflag"

	"github.com/mayooot/detect-gpu/internal/stat"
	"github.com/mayooot/detect-gpu/pkg/detect"
)

var (
	addr    = flag.StringP("addr", "a", "0.0.0.0:2376", "Address of detect server, format: ip:port, default: 0.0.0.0:2376")
	path    = flag.StringP("path", "p", "/api/v1/detect/gpu", "Path of detect server, default: /api/v1/detect/gpu")
	timeout = flag.DurationP("timeout", "t", 1*time.Second, "Timeout of detect gpu, default: 1s")
)

type program struct {
	stat   *stat.Stat
	client *detect.Client
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(svc.Environment) error {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	if len(*addr) == 0 || len(*path) == 0 || *timeout == 0 {
		return fmt.Errorf("addr, path, timeout must be set, "+
			"addr: %s, path: %s, timeout: %d",
			*addr, *path, *timeout)
	}

	p.stat = stat.NewStat(stat.WithAddr(*addr),
		stat.WithPath(*path),
		stat.WithClient(
			detect.NewClient(detect.WithTimeout(*timeout)),
		))

	if err := p.client.Init(); err != nil {
		return err
	}
	return nil
}

func (p *program) Start() error {
	go func() {
		p.stat.Run()
	}()
	return nil
}

func (p *program) Stop() error {
	if err := p.client.Close(); err != nil {
		log.Error(err)
		return err
	}
	log.Info("shutdown nvml success")
	log.Info("detect gpu server close")
	return nil
}
