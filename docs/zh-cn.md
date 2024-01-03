# Detect-GPU

![license](https://img.shields.io/hexpm/l/plug.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mayooot/detect-gpu)](https://goreportcard.com/report/github.com/mayooot/detect-gpu)

[English](..%2FREADME.md)

# 概述

调用 [go-nvml](https://github.com/NVIDIA/go-nvml)，获取 Linux GPU服务器上的 NVIDIA GPU 信息，并以 API 的形式暴露。

因为开发时可能使用 mac 或者 windows 系统，但是 go-nvml 需要 Linux NVIDIA 的驱动，所以开发时会报错。

所以把它拆分出来，作为一个独立的 HTTP 服务，这样即使在没有 NVIDIA 驱动的环境中，也可以成功编译和运行主要应用程序。

- [Detect-GPU](#detect-gpu)
- [概述](#概述)
- [快速开始](#快速开始)
    - [从源码构建](#从源码构建)
    - [运行](#运行)
    - [测试](#测试)
    - [使用](#使用)
- [在Go项目中引用](#在Go项目中引用)
    - [简单的例子](#简单的例子)
- [贡献代码](#贡献代码)

# 快速开始

## 从源码构建

```shell
$ git clone https://github.com/mayooot/detect-gpu
$ cd detect-gpu
$ make linux
```

## 运行

如果需要的话，可以指定参数

```shell
$ ./detect-gpu-linux-amd64 -h
Usage of ./detect-gpu-linux-amd64:
  -r, --pattern string   Pattern of detect server (default "/api/v1/detect/gpu")
  -p, --port string      Port of detect server, format :port (default ":2376")
  -t, --td duration      Timeout duration for detect gpu (default 5s)
pflag: help requested

$ ./detect-gpu-linux-amd64
```

## 测试

```shell
go test -v pkg/detect/*
```

## 使用

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

# 在Go项目中引用

使用标准的 go get 命令可以获得 `detect-gpu`。

```shell
$ go get github.com/mayooot/detect-gpu/pkg/detect 
```

## 简单的例子

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

# 贡献代码

欢迎贡献代码或 issue!
