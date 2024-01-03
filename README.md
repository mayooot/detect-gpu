# Detect-GPU

![license](https://img.shields.io/hexpm/l/plug.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mayooot/detect-gpu)](https://goreportcard.com/report/github.com/mayooot/detect-gpu)

[简体中文](/docs/zh-cn.md)

# Overview

`detect-gpu` is a HTTP server that calls [go-nvml](https://github.com/NVIDIA/go-nvml) and provides an api to get
information about the NVIDIA GPU on a Linux server.

Because we may use macOS or Windows for development, but [go-nvml](https://github.com/NVIDIA/go-nvml) needs Linux NVIDIA
driver, so it will report error during development.

And we split it out as a standalone HTTP service so that we can compile and run the main application successfully even
in an environment without NVIDIA drivers.

- [Detect-GPU](#detect-gpu)
- [Overview](#Overview)
- [Quick Start](#quick-start)
- [Build from source](#build-from-source)
- [Test](#test)
- [Installation](#installation)
- [Usage](#usage)
- [Contribute](#contribute)

# Quick Start

Downloading the binary executable from [release](https://github.com/mayooot/detect-gpu/releases).
And run it.

```shell
$ ./detect-gpu-linux-amd64 -h
Usage of ./detect-gpu-linux-amd64:
  -r, --pattern string   Pattern of detect server (default "/api/v1/detect/gpu")
  -p, --port string      Port of detect server, format :port (default ":2376")
  -t, --td duration      Timeout duration for detect gpu (default 5s)
pflag: help requested

$ ./detect-gpu-linux-amd64
2024/01/03 22:36:26 stat.go:30: [info] detect server start success, listen on :2376
```

Send a GET request using cURL or any language.

```shell
$ curl 127.0.0.1:2376/api/v1/detect/gpu
[
    {
        "index":0,
        "uuid":"GPU-uuid",
        "name":"NVIDIA A100 80GB PCIe",
        "memoryInfo":{
            "Total":85899345920,
            "Free":63216877568,
            "Used":22682468352
        },
        "powerUsage":74634,
        "powerState":0,
        "powerManagementDefaultLimit":300000,
        "informImageVersion":"1001.0230.00.03",
        "systemGetDriverVersion":"525.85.12",
        "systemGetCudaDriverVersion":12000,
        "tGraphicsRunningProcesses":[]
    },
    {
        "index":1,
        "uuid":"GPU-uuid",
        "name":"NVIDIA A100 80GB PCIe",
        "memoryInfo":{
            "Total":85899345920,
            "Free":30687952896,
            "Used":55211393024
        },
        "powerUsage":65507,
        "powerState":0,
        "powerManagementDefaultLimit":300000,
        "informImageVersion":"1001.0230.00.03",
        "systemGetDriverVersion":"525.85.12",
        "systemGetCudaDriverVersion":12000,
        "tGraphicsRunningProcesses":[]
    }
]
```

# Build from source

```shell
$ git clone https://github.com/mayooot/detect-gpu
$ cd detect-gpu
$ make linux
```

# Test

```shell
go test -v pkg/detect/*
```

# Installation

`detect-gpu` is available using the standard go get command.

Install by running:

```shell
$ go get github.com/mayooot/detect-gpu/pkg/detect 
```

# Usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mayooot/detect-gpu/pkg/detect"
)

func main() {
	infos, err := detect.DetectGpu(1 * time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, info := range infos {
		fmt.Printf("%+v\n", info)
	}
}
```

# Contribute

Feel free to open issues and pull requests. Any feedback is highly appreciated!