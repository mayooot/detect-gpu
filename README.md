# Detect-GPU

![license](https://img.shields.io/hexpm/l/plug.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mayooot/detect-gpu)](https://goreportcard.com/report/github.com/mayooot/detect-gpu)

[简体中文](/docs/zh-cn.md)

# Overview

`Detect-GPU` is an HTTP server that calls [go-nvml](https://github.com/NVIDIA/go-nvml) and provides an api to get
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
$ ./detect-gpu-linux-amd64 
2024/01/08 06:46:03 stat.go:60: [info] detect server start success, listen on 0.0.0.0:2376
2024/01/08 06:46:03 stat.go:61: [info] detect gpu timeout: 1000 ms
2024/01/08 06:46:03 stat.go:62: [info] ROUTES: 
2024/01/08 06:46:03 stat.go:63: [info] GET              -->             /api/v1/detect/gpu

$ ./detect-gpu-linux-amd64 -h
Usage of ./detect-gpu-linux-amd64:
  -a, --addr string        Address of detect server, format: ip:port, default: 0.0.0.0:2376 (default "0.0.0.0:2376")
  -p, --path string        Path of detect server, default: /api/v1/detect/gpu (default "/api/v1/detect/gpu")
  -t, --timeout duration   Timeout of detect gpu, default: 1s (default 1s)
pflag: help requested
```

Send a GET request using cURL or any language.

```shell
$ curl 127.0.0.1:2376/api/v1/detect/gpu
[
  {
    "index": 0,
    "uuid": "GPU-3d55828f-bc3b-41cb-452d-30189d49dbeb",
    "name": "NVIDIA vGPU-48GB",
    "memoryInfo": {
      "Total": 51527024640,
      "Free": 50875727872,
      "Used": 651296768
    },
    "memoryInfoV2": {
      "Version": 33554472,
      "Total": 51527024640,
      "Reserved": 651100160,
      "Free": 50875727872,
      "Used": 196608
    },
    "powerUsage": 94494,
    "powerState": 0,
    "powerManagementDefaultLimit": 450000,
    "informImageVersion": "G002.0000.00.03",
    "systemGetDriverVersion": "570.124.04",
    "systemGetCudaDriverVersion": 12080,
    "tGraphicsRunningProcesses": [],
    "utilization": {
      "Gpu": 0,
      "Memory": 0
    }
  }
]
```
Simultaneously, run `nvidia-smi` to compare both outputs.
```shell
nvidia-smi 
Fri Aug  8 17:54:21 2025       
+-----------------------------------------------------------------------------------------+
| NVIDIA-SMI 570.124.04             Driver Version: 570.124.04     CUDA Version: 12.8     |
|-----------------------------------------+------------------------+----------------------+
| GPU  Name                 Persistence-M | Bus-Id          Disp.A | Volatile Uncorr. ECC |
| Fan  Temp   Perf          Pwr:Usage/Cap |           Memory-Usage | GPU-Util  Compute M. |
|                                         |                        |               MIG M. |
|=========================================+========================+======================|
|   0  NVIDIA vGPU-48GB               On  |   00000000:C9:00.0 Off |                  Off |
| 46%   49C    P0             94W /  450W |       1MiB /  49140MiB |      0%      Default |
|                                         |                        |                  N/A |
+-----------------------------------------+------------------------+----------------------+
                                                                                         
+-----------------------------------------------------------------------------------------+
| Processes:                                                                              |
|  GPU   GI   CI              PID   Type   Process name                        GPU Memory |
|        ID   ID                                                               Usage      |
|=========================================================================================|
|  No running processes found                                                             |
+-----------------------------------------------------------------------------------------+
```
# Build from source

```shell
$ git clone https://github.com/mayooot/detect-gpu
$ cd detect-gpu
$ make linux
```

# Test

```shell
$ make test
```

# Installation

`Detect-GPU` is available using the standard go get command.

Install by running:

```shell
$ go get github.com/mayooot/detect-gpu/pkg/detect 
```

# Usage

You can refer to the [example](./examples/main.go) for usage.

Like this:

```go
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

```

# Contribute

Feel free to open issues and pull requests. Any feedback is highly appreciated!