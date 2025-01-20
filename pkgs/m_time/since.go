package m_time

import (
	"strings"
	"time"
)

// Since 返回自 timeSince 以来经过的时间, 去除小数点, 格式化为字符串。
func Since(timeSince time.Time, m ...Mux) string {
	value := time.Since(timeSince).String()
	digit := 0
	if digit < 0 {
		// 如果小数位数为负数，视为无效，直接返回原始字符串
		return value
	}

	regex := mux.FormatDecimalRegex
	if len(m) > 0 && m[0].FormatDecimalRegex != nil {
		regex = m[0].FormatDecimalRegex
	}
	matches := regex.FindAllStringSubmatchIndex(value, -1)
	if matches == nil {
		// 如果没有匹配项，直接返回原始字符串
		return value
	}

	var result strings.Builder
	prevEnd := 0

	for _, match := range matches {
		start := match[0]
		end := match[1]

		// 将非匹配部分直接写入结果
		if start > prevEnd {
			result.WriteString(value[prevEnd:start])
		}

		// 提取各个捕获组的内容
		intStart, intEnd := match[2], match[3]
		decimalStart, decimalEnd := match[4], match[5]
		unitStart, unitEnd := match[6], match[7]

		intPart := value[intStart:intEnd]
		var decimalPart string
		if decimalStart != -1 && decimalEnd != -1 {
			decimalPart = value[decimalStart+1 : decimalEnd] // 去掉点号
		}
		unit := value[unitStart:unitEnd]

		// 根据 n 处理小数部分
		if digit == 0 {
			// 不保留小数位，直接使用整数部分
			result.WriteString(intPart + unit)
		} else {
			if decimalPart != "" {
				if len(decimalPart) > digit {
					// 保留前 n 位
					decimalPart = decimalPart[:digit]
				} else {
					// 如果小数位数不足 n 位，保留现有的小数部分
					// 可以选择是否补零，这里选择不补零
				}
				// 拼接整数部分、小数点和处理后的小数部分
				result.WriteString(intPart + "." + decimalPart + unit)
			} else {
				// 没有小数部分，直接拼接整数部分和单位
				result.WriteString(intPart + unit)
			}
		}

		prevEnd = end
	}

	// 将最后一部分写入结果
	if prevEnd < len(value) {
		result.WriteString(value[prevEnd:])
	}

	return result.String()
}
