package m_redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetListLtrim(rds *redis.Client, ctx context.Context, keyName string, count int64) ([]string, error) {
	script := fmt.Sprintf(`
		local values = redis.call('lrange', KEYS[1], 0, %d)
		redis.call('ltrim', KEYS[1], %d + 1, -1)
		return values
	`, count-1, count-1) // Lua脚本通过参数设置获取和删除的范围

	// 获取上下文
	val, err := rds.Eval(ctx, script, []string{keyName}, int64(count)-1).Result()
	if err != nil {
		return nil, err // 如果执行Lua脚本出错，则返回错误
	}

	// 处理Lua脚本返回的数据
	values, ok := val.([]interface{})
	if !ok {
		return nil, fmt.Errorf("error processing data from Redis")
	}
	var result []string
	for _, v := range values {
		str, isString := v.(string)
		if !isString {
			continue // 如果不是字符串，跳过这个元素
		}
		result = append(result, str)
	}
	return result, nil
}
