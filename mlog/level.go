package mlog

// Log levels

const (
	levelFatal = iota // 0
	levelError        // 1
	levelWarn         // 2
	levelInfo         // 3
	levelDebug        // 4
	levelTrace        // 5
)

func levelToString(level int) string {
	switch level {
	case levelFatal:
		return "FATAL"
	case levelError:
		return "ERROR"
	case levelWarn:
		return "WARN"
	case levelInfo:
		return "INFO"
	case levelDebug:
		return "DEBUG"
	case levelTrace:
		return "TRACE"
	default:
		return "UNKNOWN"
	}
}
