package m_time

import (
	"time"

	"github.com/lwmacct/241224-go-template-pkgs/pkgs/m_to"
)

/*
将时间n字符串转换为秒级 Unix 时间戳，并以字符串形式返回。

input:
  - dateTime: 时间字符串，格式为 "2006-01-02 15:04:05"
  - m: 可选的 Mux 指针，用于指定自定义时区，默认为 CST (UTC+8)

output:
  - string: 秒级 Unix 时间戳字符串，若解析失败则返回 "0"
*/

func ToUnixString(dateTime string, m ...*Mux) string {
	return m_to.String(ToUnixInt64(dateTime, m...))
}

/*
将时间字符串转换为秒级 Unix 时间戳，并返回 int64 类型的时间戳。

input:
  - dateTime: 时间字符串，格式为 "2006-01-02 15:04:05"
  - m: 可选的 Mux 指针，用于指定自定义时区，默认为 CST (UTC+8)

output:
  - int64: 秒级 Unix 时间戳，若解析失败则返回 0
*/
func ToUnixInt64(dateTime string, m ...*Mux) int64 {
	var loc = mux.Location
	if len(m) > 0 && m[0].Location != nil {
		loc = m[0].Location
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", dateTime, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}
