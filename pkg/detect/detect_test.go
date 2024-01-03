package detect

import (
	"testing"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/assert"
)

func TestDetectGpu(t *testing.T) {
	ret := nvml.Init()
	assert.Equal(t, nvml.SUCCESS, ret)
	count, ret := nvml.DeviceGetCount()
	assert.Equal(t, nvml.SUCCESS, ret)
	defer nvml.Shutdown()

	t.Run("Detect gpus", func(t *testing.T) {
		gpus, err := DetectGpu(500 * time.Millisecond)
		assert.Nil(t, err)
		assert.NotNil(t, gpus)
		assert.Equal(t, count, len(gpus))
	})
}

func TestDetectGpu_Timeout(t *testing.T) {
	allowTimeDuration := 50 * time.Millisecond
	sleepTimeDuration := 100 * time.Millisecond
	timeOutDetectGpu := func(td time.Duration) ([]*gpuInfo, error) {
		time.Sleep(sleepTimeDuration)
		return DetectGpu(td)
	}

	gpus, err := timeOutDetectGpu(allowTimeDuration)
	assert.NotNil(t, err, "Timeout error should be returned")
	assert.Nil(t, gpus)
}
