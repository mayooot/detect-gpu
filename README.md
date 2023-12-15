# detect-gpu

调用 [go-nvml](https://github.com/NVIDIA/go-nvml) 获取宿主机的 GPU 信息

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