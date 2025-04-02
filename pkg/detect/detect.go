package detect

import (
	"context"
	"fmt"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

type Option func(c *Client)

// Client calls nvml to query gpus
type Client struct {
	Timeout time.Duration
}

func NewClient(opts ...Option) *Client {
	c := &Client{}
	for _, apply := range opts {
		apply(c)
	}
	return c
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

func (c *Client) Init() error {
	if ret := nvml.Init(); ret != nvml.SUCCESS {
		return fmt.Errorf("unable to initialize NVML: %v", nvml.ErrorString(ret))
	}
	return nil
}

func (c *Client) Close() error {
	if ret := nvml.Shutdown(); ret != nvml.SUCCESS {
		return fmt.Errorf("unable to shutdown NVML: %v", nvml.ErrorString(ret))
	}
	return nil
}

// DetectGpu return error if the timeout is exceeded
func (c *Client) DetectGpu() ([]*gpuInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	resultCh := make(chan []*gpuInfo, 1)
	errCh := make(chan error, 1)
	go func() {
		gpuInfos, err := invokeNvml()
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- gpuInfos
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errCh:
		return nil, err
	case gpuInfos := <-resultCh:
		return gpuInfos, nil
	}
}

func invokeNvml() ([]*gpuInfo, error) {
	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		return nil, fmt.Errorf("unable to get gpuInfo count: %v", nvml.ErrorString(ret))
	}
	gpuInfos := make([]*gpuInfo, 0, count)

	for i := 0; i < count; i++ {
		info := &gpuInfo{Index: i}
		device, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}

		uuid, ret := device.GetUUID()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get uuid of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.UUID = uuid

		name, ret := device.GetName()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get name of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.Name = name

		memoryInfo, ret := device.GetMemoryInfo()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get memory info of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.MemoryInfo = memoryInfo

		powerUsage, ret := device.GetPowerUsage()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get power usage of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.PowerUsage = powerUsage

		powerState, ret := device.GetPowerState()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get power state of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.PowerState = powerState

		managementDefaultLimit, ret := device.GetPowerManagementDefaultLimit()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get power management default limit of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.PowerManagementDefaultLimit = managementDefaultLimit

		version, ret := device.GetInforomImageVersion()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get info image version of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.InformImageVersion = version

		driverVersion, ret := nvml.SystemGetDriverVersion()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get system driver version: %v", nvml.ErrorString(ret))
		}
		info.DriverVersion = driverVersion

		cudaDriverVersion, ret := nvml.SystemGetCudaDriverVersion()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get CUDA driver version: %v", nvml.ErrorString(ret))
		}
		info.CUDADriverVersion = cudaDriverVersion

		computeRunningProcesses, ret := device.GetGraphicsRunningProcesses()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get graphics running processes of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.GraphicsRunningProcesses = computeRunningProcesses


		utilization, ret := device.GetUtilizationRates()
		if ret != nvml.SUCCESS {
			return gpuInfos, fmt.Errorf("unable to get utilization rates of gpuInfo at index %d: %v", i, nvml.ErrorString(ret))
		}
		info.Utilization = utilization

		gpuInfos = append(gpuInfos, info)
	}
	return gpuInfos, nil
}
