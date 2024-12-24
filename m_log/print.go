package m_log

import "os"

func (t *Ts) Fatal(fields H) *Ts {
	if t.config.Level <= levelFatal {
		return t
	}
	fields["level"] = levelToString(levelFatal)
	t.logWithLevel(fields, 2)
	os.Exit(1)
	return t
}

func (t *Ts) Error(fields H) *Ts {
	if t.config.Level <= levelError {
		return t
	}
	fields["level"] = levelToString(levelError)
	return t.logWithLevel(fields, 2)
}

func (t *Ts) Warn(fields H) *Ts {
	if t.config.Level <= levelWarn {
		return t
	}
	fields["level"] = levelToString(levelWarn)
	return t.logWithLevel(fields, 2)
}

func (t *Ts) Info(fields H) *Ts {
	if t.config.Level <= levelInfo {
		return t
	}
	fields["level"] = levelToString(levelInfo)
	return t.logWithLevel(fields, 2)
}

func (t *Ts) Debug(fields H) *Ts {
	if t.config.Level <= levelDebug {
		return t
	}
	fields["level"] = levelToString(levelDebug)
	return t.logWithLevel(fields, 2)
}

func (t *Ts) Trace(fields H) *Ts {
	if t.config.Level <= levelTrace {
		return t
	}
	fields["level"] = levelToString(levelTrace)
	return t.logWithLevel(fields, 2)
}
