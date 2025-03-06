package filelog

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFileLogger(t *testing.T) {

	SetLogFilePath("../test/log")
	l1 := Get()
	l2 := Get()

	logs := make([][]byte, 0, 10)

	for i := range 10 {
		logs = append(logs, []byte("log"+strconv.Itoa(i)))
	}

	t.Run("Instance", func(t *testing.T) {
		l1Ptr := fmt.Sprintf("%p", l1)
		l2Ptr := fmt.Sprintf("%p", l2)
		if l1Ptr != l2Ptr {
			t.Error("instances get diff address")
		}
	})

	t.Run("LogInfo", func(t *testing.T) {
		for _, v := range logs {
			l1.LogInfo(v)
		}
	})

	t.Run("LogWarning", func(t *testing.T) {
		for _, v := range logs {
			l1.LogWarning(v)
		}
	})

	t.Run("LogError", func(t *testing.T) {
		for _, v := range logs {
			l1.LogError(v)
		}
	})

	t.Run("CloseLogFile", func(t *testing.T) {
		l1.CloseLogFile()
	})
}
