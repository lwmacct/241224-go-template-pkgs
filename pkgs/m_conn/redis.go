package conn

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Err error
	Ctx context.Context
	Raw *redis.Client
}

// https://github.com/redis/go-redis?tab=readme-ov-file#connecting-via-a-redis-url
func NewRedis(url string) (*Redis, error) {
	t := &Redis{}
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	t.Raw = redis.NewClient(opts)
	t.Ctx = context.Background()
	_, err = t.Raw.Ping(t.Ctx).Result()
	if err != nil {
		return nil, err
	}
	return t, nil
}

// SetExpiration 设置指定键的过期时间
func (rj *Redis) SetExpiration(key string, expiration time.Duration) error {
	cmd := rj.Raw.Expire(rj.Ctx, key, expiration)
	return cmd.Err()
}

// JSON.SET 命令封装
func (rj *Redis) JsonSet(key string, path string, json string) error {
	cmd := rj.Raw.Do(rj.Ctx, "JSON.SET", key, path, json)
	return cmd.Err()
}

// JSON.GET 命令封装
func (rj *Redis) JsonGet(key string, path ...string) (interface{}, error) {
	// 初始化 args 并将 "JSON.GET" 和 key 添加到其中
	args := []interface{}{"JSON.GET", key}

	// 逐个追加 path 切片的元素，确保 args 是扁平化的
	for _, p := range path {
		args = append(args, p)
	}

	// 使用展开操作符将 args 作为可变参数传递
	cmd := rj.Raw.Do(rj.Ctx, args...)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return cmd.Result()
}

// 取得 JSON 字符串
func (rj *Redis) JsonGetString(key string, path ...string) (string, error) {
	data, err := rj.JsonGet(key, path...)
	if err != nil {
		return "", err
	}
	jsonStr, ok := data.(string)
	if !ok {
		return "", fmt.Errorf("value is not a string")
	}
	return jsonStr, nil
}

// 取得 JSON 数组
func (rj *Redis) JsonGetStringArr(key string, path ...string) ([]string, error) {
	var result []string
	jsonStr, err := rj.JsonGetString(key, path...)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return result, err
	}
	return result, nil
}

// JSON.DEL 命令封装
func (rj *Redis) JsonDel(key string, path string) error {
	cmd := rj.Raw.Do(rj.Ctx, "JSON.DEL", key, path)
	return cmd.Err()
}

// JSON.INCRBY 命令封装（示例）
func (rj *Redis) IncrBy(key string, path string, increment int64) (int64, error) {
	cmd := rj.Raw.Do(rj.Ctx, "JSON.NUMINCRBY", key, path, increment)
	if err := cmd.Err(); err != nil {
		return 0, err
	}
	newValue, err := cmd.Int64()
	if err != nil {
		return 0, err
	}
	return newValue, nil
}
