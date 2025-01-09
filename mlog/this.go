package mlog

import (
	"os"
)

var ts *Ts = New(nil)

func NewTs(config *Config) *Ts {
	ts = New(config)
	return ts
}

func Fatal(fields H) *Ts {
	if ts.config.Level < levelFatal {
		return ts
	}

	fields["level"] = levelToString(levelFatal)
	ts.logWithLevel(fields, 2)
	os.Exit(1)
	return ts
}

func Error(fields H) *Ts {
	if ts.config.Level < levelError {
		return ts
	}
	fields["level"] = levelToString(levelError)
	return ts.logWithLevel(fields, 2)
}

func Warn(fields H) *Ts {
	if ts.config.Level < levelWarn {
		return ts
	}
	fields["level"] = levelToString(levelWarn)
	return ts.logWithLevel(fields, 2)
}

func Info(fields H) *Ts {
	if ts.config.Level < levelInfo {
		return ts
	}
	fields["level"] = levelToString(levelInfo)
	return ts.logWithLevel(fields, 2)
}

func Debug(fields H) *Ts {
	if ts.config.Level < levelDebug {
		return ts
	}
	fields["level"] = levelToString(levelDebug)
	return ts.logWithLevel(fields, 2)
}

func Trace(fields H) *Ts {
	if ts.config.Level < levelTrace {
		return ts
	}
	fields["level"] = levelToString(levelTrace)
	return ts.logWithLevel(fields, 2)
}

func ShowLevel() {
	println(ts.config.Level)
}
