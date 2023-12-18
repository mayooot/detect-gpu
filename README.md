# detect-gpu

使用 [go-nvml](https://github.com/NVIDIA/go-nvml) 库，获取宿主机的 GPU 信息。

## Q&A

**Q: 为什么不直接在代码中调用 `go-nvml` 库？**

A: 直接调用 `go-nvml` 库需要在Linux服务器上安装NVIDIA驱动。然而，在开发阶段引入此库可能导致编译和运行失败。

**Q: 为什么选择通过HTTP服务获取GPU信息？**

A: 为了避免在开发过程中由于缺少NVIDIA驱动而导致编译和运行问题，我们将与GPU相关的代码拆分成一个独立的HTTP服务。这样，即使在没有NVIDIA驱动的环境中，我们仍然可以成功编译和运行主要应用程序，而不必担心与GPU相关的依赖关系。

## 使用

### 构建

**必须在 Linux 系统上构建**

```shell
git clone https://github.com/mayooot/detect-gpu
cd detect-gpu
make linux
```

### 运行

```
./detect-gpu-linux-amd64
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

使用标准的go get命令可以获得 `detect-gpu`

通过运行：
```
go get github.com/mayooot/detect-gpu/pkg/detect 
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