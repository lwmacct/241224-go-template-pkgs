package m_time

import (
	"time"

	"github.com/lwmacct/241220-go-pkgs/241222/m_to"
)

type Cfg struct {
	Location *time.Location
	Time     time.Time
}

type Ts struct {
	Cfg *Cfg
}

func New(cfg *Cfg) *Ts {
	if cfg == nil {
		cfg = &Cfg{}
		cfg.Location = time.FixedZone("CST", 8*3600) // CST (China Standard Time) UTC+8
		cfg.Time = time.Now().In(cfg.Location)
	}

	if cfg.Location == nil {
		cfg.Location = time.FixedZone("CST", 8*3600) // CST (China Standard Time) UTC+8
	}

	if cfg.Time.IsZero() {
		cfg.Time = time.Now().In(cfg.Location)
	}
	return &Ts{
		Cfg: cfg,
	}
}

func (t *Ts) Now() time.Time {
	return time.Now().In(t.Cfg.Location)
}

func (t *Ts) ToString(format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return t.Cfg.Time.In(t.Cfg.Location).Format(format)
}

func (t *Ts) ToUnix() int64 {
	return t.Cfg.Time.In(t.Cfg.Location).Unix()
}

func (t *Ts) ToUnixString() string {
	return m_to.String(t.ToUnix())
}

func (t *Ts) ToCST() time.Time {
	return t.Cfg.Time.In(t.Cfg.Location)
}

func (t *Ts) ToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).In(t.Cfg.Location)
}

// 将时间向前取整到最接近的 5 的倍数
func (t *Ts) Round5m() time.Time {
	// 获取分钟数
	minutes := t.Cfg.Time.Minute()
	// 向前取整到最接近的 5 的倍数
	roundedMinutes := (minutes / 5) * 5
	// 返回新时间
	return time.Date(t.Cfg.Time.Year(), t.Cfg.Time.Month(), t.Cfg.Time.Day(), t.Cfg.Time.Hour(), roundedMinutes, 0, 0, t.Cfg.Location)
}
