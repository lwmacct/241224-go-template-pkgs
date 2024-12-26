package m_syscall

import (
	"syscall"
	"time"
)

type Ts struct {
}

// GetSystemUptime 获取系统启动后的时间
func GetSystemUptime() int64 {
	var sysInfo syscall.Sysinfo_t
	syscall.Sysinfo(&sysInfo)
	uptime := time.Duration(sysInfo.Uptime) * time.Second
	return int64(uptime.Seconds()) // 返回启动后的时间，以秒为单位
}
