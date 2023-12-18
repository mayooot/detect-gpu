# detect-gpu

使用 [go-nvml](https://github.com/NVIDIA/go-nvml) 库，获取宿主机的 GPU 信息。


## Q&A

**Q: 为什么不直接在代码中调用 `go-nvml` 库？**

A: 直接调用 `go-nvml` 库需要在 Linux 服务器上安装 NVIDIA 驱动。然而，在开发阶段引入此库可能导致编译和运行失败。

**Q: 为什么选择通过HTTP服务获取GPU信息？**

A: 为了避免在开发过程中由于缺少 NVIDIA 驱动而导致编译和运行问题，我们将与 GPU 相关的代码拆分成一个独立的 HTTP 服务。这样，即使在没有 NVIDIA 驱动的环境中，我们仍然可以成功编译和运行主要应用程序，而不必担心与 GPU 相关的依赖关系。

## 使用

### 构建

**必须在 Linux 系统上构建**

```shell
$ git clone https://github.com/mayooot/detect-gpu
$ cd detect-gpu
$ make linux
```

### 运行

```shell
$ ./detect-gpu-linux-amd64
```

可以通过传递参数来修改默认的端口和请求路径，默认端口为 `:2376`，请求路径为 `/api/v1/detect/gpu`

```shell
$ ./detect-gpu-linux-amd64 -h
Usage of ./detect-gpu-linux-amd64:
      --pattern string   Pattern of detect server (default "/api/v1/detect/gpu")
      --port string      Port of detect server, format :port (default ":2376")
pflag: help requested
```



## 在 Go 项目中使用

使用标准的 go get 命令可以获得 `detect-gpu`。

**需要注意的是**：运行下面代码的机器也必须得是 Linux 系统且有 NVIDIA 相关驱动。

通过运行：
```shell
$ go get github.com/mayooot/detect-gpu/pkg/detect 
```
### 简单的例子

```go
package main

import (
	"fmt"
	"log"

	"github.com/mayooot/detect-gpu/pkg/detect"
)

func main() {
	infos, err := detect.DetectGpu()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, info := range infos {
		fmt.Printf("%+v\n", info)
	}
}

```

## 示例
请求：
```shell
$ curl --location 'your_addr/api/v1/detect/gpu' 
```

响应：
```json
[
  {
    "UUID": "GPU-7a42be89-64fe-5383-c7be-49d199a96b3d",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 24386076672,
      "Used": 61513269248
    },
    "powerUsage": 74173,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  },
  {
    "UUID": "GPU-dc6d913c-8df4-a9a4-49e6-b82fcba5a6f9",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 30679629824,
      "Used": 55219716096
    },
    "powerUsage": 65251,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  },
  {
    "UUID": "GPU-82fbe07b-200b-1d4c-4fbe-b0b54db86be5",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 71479721984,
      "Used": 14419623936
    },
    "powerUsage": 71709,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  },
  {
    "UUID": "GPU-36009026-9470-a2e0-73d3-222a63b82e4e",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 71945289728,
      "Used": 13954056192
    },
    "powerUsage": 62076,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  },
  {
    "UUID": "GPU-bc85a406-0357-185f-a56c-afb49572bdbe",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 57613352960,
      "Used": 28285992960
    },
    "powerUsage": 71465,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  },
  {
    "UUID": "GPU-c6b3ca5f-c1ac-8171-582b-737b70a6bbce",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 85021032448,
      "Used": 878313472
    },
    "powerUsage": 45968,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  },
  {
    "UUID": "GPU-04adce59-e7fc-19ed-6800-bc09e5f8fa31",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 85021032448,
      "Used": 878313472
    },
    "powerUsage": 46814,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  },
  {
    "UUID": "GPU-281d9730-5a26-7c56-12fb-3a3d5a24ab68",
    "name": "NVIDIA A100 80GB PCIe",
    "memoryInfo": {
      "Total": 85899345920,
      "Free": 85021032448,
      "Used": 878313472
    },
    "powerUsage": 46229,
    "powerState": 0,
    "powerManagementDefaultLimit": 300000,
    "infoImageVersion": "1001.0230.00.03",
    "inforomImageVersion": "",
    "systemGetDriverVersion": "525.85.12",
    "systemGetCudaDriverVersion": 12000,
    "tGraphicsRunningProcesses": []
  }
]
```