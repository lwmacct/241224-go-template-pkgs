package m_time

import (
	"regexp"
	"time"
)

type Mux struct {
	Location           *time.Location
	FormatDecimalRegex *regexp.Regexp
}

var mux *Mux = &Mux{
	Location: time.FixedZone("CST", 8*3600), // CST (China Standard Time) UTC+8

	// 正则表达式匹配数字和单位
	// 解释:
	// (-?\d+)    : 捕获整数部分（可选负号）
	// (\.\d+)?   : 可选的小数部分
	// ([a-zA-Z]+): 捕获单位
	FormatDecimalRegex: regexp.MustCompile(`(-?\d+)(\.\d+)?([a-zA-Z]+)`),
}

func GetMux() *Mux {
	return mux
}
