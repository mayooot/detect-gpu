package detect

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetectGpu(t *testing.T) {
	client := NewClient(WithTimeout(500 * time.Millisecond))
	err := client.Init()
	require.Nil(t, err)
	defer client.Close()

	t.Run("Get All Gpus", func(t *testing.T) {
		gpus, err := client.DetectGpu()
		assert.Nil(t, err)
		assert.NotNil(t, gpus)
		assert.NotEqual(t, 0, len(gpus))
	})
}

func TestDetectGpu_Timeout(t *testing.T) {
	allowTimeDuration := 50 * time.Millisecond
	sleepTimeDuration := 100 * time.Millisecond

	client := NewClient(WithTimeout(allowTimeDuration))
	err := client.Init()
	require.Nil(t, err)
	defer client.Close()

	t.Run("Test Timeout Control", func(t *testing.T) {
		timeOutDetectGpu := func(td time.Duration) ([]*gpuInfo, error) {
			time.Sleep(sleepTimeDuration)
			return client.DetectGpu()
		}

		gpus, err := timeOutDetectGpu(allowTimeDuration)
		assert.NotNil(t, err, "Timeout error should be returned")
		assert.Nil(t, gpus)
	})
}
