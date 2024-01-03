package detect

import (
	"context"
	"fmt"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/ngaut/log"
)

type gpuInfo struct {
	Index                       int                `json:"index"`
	UUID                        string             `json:"uuid"`
	Name                        string             `json:"name"`
	MemoryInfo                  nvml.Memory        `json:"memoryInfo"`
	PowerUsage                  uint32             `json:"powerUsage"`
	PowerState                  nvml.Pstates       `json:"powerState"`
	PowerManagementDefaultLimit uint32             `json:"powerManagementDefaultLimit"`
	InformImageVersion          string             `json:"informImageVersion"`
	DriverVersion               string             `json:"systemGetDriverVersion"`
	CUDADriverVersion           int                `json:"systemGetCudaDriverVersion"`
	GraphicsRunningProcesses    []nvml.ProcessInfo `json:"tGraphicsRunningProcesses"`
}

func DetectGpu(td time.Duration) ([]*gpuInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), td)
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
	if ret := nvml.Init(); ret != nvml.SUCCESS {
		return nil, fmt.Errorf("unable to initialize NVML: %v", nvml.ErrorString(ret))
	}
	defer func() {
		ret := nvml.Shutdown()
		if ret != nvml.SUCCESS {
			log.Errorf("unable to shutdown NVML: %v", nvml.ErrorString(ret))
		}
	}()

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

		gpuInfos = append(gpuInfos, info)
	}
	return gpuInfos, nil
}
