package detect

import (
	"fmt"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

type gpuInfo struct {
	UUID                        string             `json:"UUID"`
	Name                        string             `json:"name"`
	MemoryInfo                  nvml.Memory        `json:"memoryInfo"`
	PowerUsage                  uint32             `json:"powerUsage"`
	PowerState                  nvml.Pstates       `json:"powerState"`
	PowerManagementDefaultLimit uint32             `json:"powerManagementDefaultLimit"`
	InfoImageVersion            string             `json:"infoImageVersion"`
	InforomImageVersion         string             `json:"inforomImageVersion"`
	DriverVersion               string             `json:"systemGetDriverVersion"`
	CUDADriverVersion           int                `json:"systemGetCudaDriverVersion"`
	GraphicsRunningProcesses    []nvml.ProcessInfo `json:"tGraphicsRunningProcesses"`
}

func DetectGpu() (gpuInfos []gpuInfo, err error) {
	if ret := nvml.Init(); ret != nvml.SUCCESS {
		return gpuInfos, fmt.Errorf("unable to initialize NVML: %v", nvml.ErrorString(ret))
	}
	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		return gpuInfos, fmt.Errorf("unable to get gpuInfo count: %v", nvml.ErrorString(ret))
	}

	for i := 0; i < count; i++ {
		var info gpuInfo
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
		info.InfoImageVersion = version

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
