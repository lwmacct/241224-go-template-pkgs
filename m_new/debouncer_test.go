package m_new

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestDebouncer 测试 Debouncer 的功能
func TestDebouncer(t *testing.T) {
	tests := []struct {
		name            string
		interval        time.Duration
		triggerInterval time.Duration
		triggerCount    int
		leading         bool
		expectedExec    int
		waitDuration    time.Duration
	}{
		{
			name:            "Trailing Edge Only",
			interval:        500 * time.Millisecond,
			triggerInterval: 100 * time.Millisecond,
			triggerCount:    10,
			leading:         false,
			expectedExec:    1,
			waitDuration:    1 * time.Second,
		},
		{
			name:            "Leading Edge Only",
			interval:        500 * time.Millisecond,
			triggerInterval: 100 * time.Millisecond,
			triggerCount:    10,
			leading:         true,
			expectedExec:    2, // 首次触发和最后一次触发
			waitDuration:    1 * time.Second,
		},
		{
			name:            "Leading and Trailing Edge",
			interval:        500 * time.Millisecond,
			triggerInterval: 100 * time.Millisecond,
			triggerCount:    10,
			leading:         true,
			expectedExec:    2, // 根据实现，可能是 2
			waitDuration:    1 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mu sync.Mutex
			execCount := 0
			action := func() {
				mu.Lock()
				defer mu.Unlock()
				execCount++
				fmt.Printf("Action executed at %v\n", time.Now())
			}

			debouncer := NewDebouncer(tt.interval, action, WithLeadingEdge(tt.leading))

			var wg sync.WaitGroup
			// 并发触发操作，每隔 triggerInterval 触发一次，共触发 triggerCount 次
			for i := 0; i < tt.triggerCount; i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					debouncer.Trigger()
					fmt.Printf("Trigger %d at %v\n", i+1, time.Now())
				}(i)
				time.Sleep(tt.triggerInterval)
			}
			// 等待所有 goroutine 完成
			wg.Wait()

			// 再等待一段时间来确保观察到防抖效果
			time.Sleep(tt.waitDuration)

			// 停止 Debouncer
			debouncer.Stop()

			// 检查执行次数是否符合预期
			mu.Lock()
			if execCount != tt.expectedExec {
				t.Errorf("Expected exec count %d, but got %d", tt.expectedExec, execCount)
			}
			mu.Unlock()
		})
	}
}
