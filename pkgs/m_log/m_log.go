package m_log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

type H map[string]interface{}

type Ts struct {
	config *Config

	H      map[string]any
	logger *log.Logger
}

type Config struct {
	File   string `group:"m_log" note:"日志文件名" default:""`
	Level  int    `group:"m_log" note:"日志级别" default:"3"`
	Stdout bool   `group:"m_log" note:"是否输出到标准输出" default:"true"`

	// lumberjack.Logger fields
	RotateMaxSize    int  `group:"m_log" note:"日志文件最大尺寸" default:"100"`
	RotateMaxAge     int  `group:"m_log" note:"日志文件最大保存天数" default:"30"`
	RotateMaxBackups int  `group:"m_log" note:"日志文件最大备份数" default:"3"`
	RotateLocalTime  bool `group:"m_log" note:"是否使用本地时间" default:"false"`
	RotateCompress   bool `group:"m_log" note:"是否压缩" default:"false"`
	// 不参与 m_cmd 解析的字段
	CallerClip  string
	OrderedKeys []string
}

func New(config *Config) *Ts {
	if config == nil {
		config = NewConfig()
		config.File = ""
	}

	ref := &Ts{config: config}
	if ref.config.File != "" {
		logger := &lumberjack.Logger{
			Filename:   ref.config.File,
			MaxSize:    ref.config.RotateMaxSize,
			MaxBackups: ref.config.RotateMaxBackups,
			MaxAge:     ref.config.RotateMaxAge,
			Compress:   ref.config.RotateCompress,
			LocalTime:  ref.config.RotateLocalTime,
		}
		ref.logger = log.New(logger, "", 0)
	}
	return ref
}

func NewConfig() *Config {
	execName := filepath.Base(os.Args[0])
	fileName := fmt.Sprintf("/var/log/%s.log", execName)
	return &Config{
		Stdout:           true,
		Level:            3,
		File:             fileName,
		OrderedKeys:      []string{"time", "level", "msg", "info", "error", "warn", "data", "flags"},
		CallerClip:       "",
		RotateMaxSize:    100,
		RotateMaxAge:     90,
		RotateMaxBackups: 3,
		RotateLocalTime:  true,
		RotateCompress:   true,
	}
}

func (e *Ts) logWithLevel(fields H, callDepth int) *Ts {
	e.setCaller(fields, callDepth+1)

	coloredOutput := colorizeJSONValues(fields, e.config.OrderedKeys)

	if e.config.File != "" {
		e.logger.Println(coloredOutput)
	}
	if e.config.Stdout {
		fmt.Println(coloredOutput)
	}

	return e
}

func colorizeJSONValues(fields H, orderedKeys []string) string {
	resetColor := "\033[0m"

	levelColors := map[string]string{
		"FATAL": "\033[95m",
		"ERROR": "\033[91m",
		"WARN":  "\033[93m",
		"INFO":  "\033[92m",
		"DEBUG": "\033[94m",
		"TRACE": "\033[90m",
	}

	keyColors := map[string]string{
		"time":  "\033[30m",
		"msg":   "\033[34m",
		"error": "\033[31m",
		"warn":  "\033[33m",
		"info":  "\033[34m",
		"data":  "\033[32m",
		"other": "\033[36m",
		"call":  "\033[35m",
	}

	var otherKeys []string
	for key := range fields {
		if key != "call" && !slices.Contains(orderedKeys, key) {
			otherKeys = append(otherKeys, key)
		}
	}
	sort.Strings(otherKeys)

	allKeys := append(orderedKeys, otherKeys...)

	if _, exists := fields["call"]; exists {
		allKeys = append(allKeys, "call")
	}

	var builder strings.Builder
	builder.WriteByte('{')

	first := true
	for _, key := range allKeys {
		value, exists := fields[key]
		if !exists {
			continue
		}

		if !first {
			builder.WriteByte(',')
		}
		first = false

		builder.WriteByte('"')
		builder.WriteString(key)
		builder.WriteString(`":`)

		valueBytes, err := json.Marshal(value)
		if err != nil {
			valueBytes = []byte(`"` + fmt.Sprintf("%v", value) + `"`)
		}

		if key == "level" {
			level := fmt.Sprintf("%v", value)
			if color, ok := levelColors[level]; ok {
				builder.WriteString(color)
				builder.Write(valueBytes)
				builder.WriteString(resetColor)
				continue
			}
		}

		if color, ok := keyColors[key]; ok {
			builder.WriteString(color)
			builder.Write(valueBytes)
			builder.WriteString(resetColor)
		} else {
			builder.WriteString(keyColors["other"])
			builder.Write(valueBytes)
			builder.WriteString(resetColor)
		}
	}
	builder.WriteByte('}')

	return builder.String()
}
