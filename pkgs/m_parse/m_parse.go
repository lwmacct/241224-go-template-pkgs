package m_parse

import "strings"

// CommaSeparatedKVPairs 解析逗号分隔的 key=value 字符串，并返回对应的 map。
// 示例:
//
//	input:  "key1=value1,key2=value2"
//	output: map[string]string{"key1":"value1","key2":"value2"}
func CommaSeparatedKVPairs(input string) map[string]string {
	pairs := strings.Split(input, ",")
	result := make(map[string]string, len(pairs))
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])
			result[key] = val
		}
	}
	return result
}
