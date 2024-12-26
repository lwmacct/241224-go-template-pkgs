package m_new

import (
	"sync"
	"time"
)

// DebouncerOption 定义 Debouncer 的可选参数
type DebouncerOption func(*Debouncer)

// WithLeadingEdge 设置是否在 leading edge 执行 action
func WithLeadingEdge(enabled bool) DebouncerOption {
	return func(d *Debouncer) {
		d.leading = enabled
	}
}

// Debouncer 封装防抖逻辑
type Debouncer struct {
	mu       sync.Mutex
	timer    *time.Timer
	interval time.Duration
	leading  bool
	action   func()
}

// NewDebouncer 创建一个新的 Debouncer 实例
func NewDebouncer(interval time.Duration, action func(), opts ...DebouncerOption) *Debouncer {
	d := &Debouncer{
		interval: interval,
		action:   action,
	}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

// Trigger 触发防抖逻辑
func (d *Debouncer) Trigger() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.leading && d.timer == nil {
		// Leading edge: 立即执行
		d.action()
	}

	if d.timer != nil {
		// 重置定时器
		d.timer.Stop()
	}

	// 设置或重置计时器
	d.timer = time.AfterFunc(d.interval, func() {
		d.mu.Lock()
		defer d.mu.Unlock()

		if !d.leading || d.timer != nil {
			// Trailing edge: 执行
			d.action()
			d.timer = nil
		}
	})
}

// Stop 停止 Debouncer，清理资源
func (d *Debouncer) Stop() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
		d.timer = nil
	}
}
