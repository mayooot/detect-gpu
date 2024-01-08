package main

import (
	"fmt"
	"time"

	"github.com/mayooot/detect-gpu/pkg/detect"
)

func main() {
	timeOutDuration := 500 * time.Millisecond

	testClient := detect.NewClient(detect.WithTimeout(timeOutDuration))
	if err := testClient.Init(); err != nil {
		panic(err)
	}
	defer testClient.Close()

	gpus, err := testClient.DetectGpu()
	if err != nil {
		panic(err)
	}
	for _, gpu := range gpus {
		fmt.Printf("%#+v\n", gpu)
	}
}
