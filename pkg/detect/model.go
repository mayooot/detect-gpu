package detect

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

type gpuInfo struct {
	Index                       int                `json:"index"`
	UUID                        string             `json:"uuid"`
	Name                        string             `json:"name"`
	MemoryInfo                  nvml.Memory        `json:"memoryInfo"`
	MemoryInfoV2                nvml.Memory_v2     `json:"memoryInfoV2"`
	PowerUsage                  uint32             `json:"powerUsage"`
	PowerState                  nvml.Pstates       `json:"powerState"`
	PowerManagementDefaultLimit uint32             `json:"powerManagementDefaultLimit"`
	InformImageVersion          string             `json:"informImageVersion"`
	DriverVersion               string             `json:"systemGetDriverVersion"`
	CUDADriverVersion           int                `json:"systemGetCudaDriverVersion"`
	GraphicsRunningProcesses    []nvml.ProcessInfo `json:"tGraphicsRunningProcesses"`
	Utilization                 nvml.Utilization   `json:"utilization"`
}
