package log

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"common/log/log4go"
)

const DEEP int = 2

func getValue(ctx context.Context, key string) string {
	var value string
	if ctx != nil {
		val := ctx.Value(key)

		if val != nil {
			value = val.(string)
		}
	}
	return value
}

// getTraceId 返回PlayerID:SessionID:TraceID的组合，冒号分隔。
func getTraceId(ctx context.Context) string {
	tid := getValue(ctx, "trace_id")
	if tid == "" {
		return ""
	} else {
		return fmt.Sprintf(tid)
	}
}

func Finest(ctx context.Context, arg0 interface{}, args ...interface{}) {
	if !log4go.IsFinestEnabled() {
		return
	}
	tid := getTraceId(ctx)
	const (
		lvl = log4go.FINEST
	)
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first, args...)
	case func() string:
		// Log the closure (no other arguments used)
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first())
	default:
		// Build a format string so that it will be similar to Sprint
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, fmt.Sprint(arg0)+strings.Repeat(" %v", len(args)), args...)
	}
}

func Fine(ctx context.Context, arg0 interface{}, args ...interface{}) {
	if !log4go.IsFineEnabled() {
		return
	}
	tid := getTraceId(ctx)

	const (
		lvl = log4go.FINE
	)
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first, args...)
	case func() string:
		// Log the closure (no other arguments used)
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first())
	default:
		// Build a format string so that it will be similar to Sprint
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, fmt.Sprint(arg0)+strings.Repeat(" %v", len(args)), args...)
	}
}

func Debug(ctx context.Context, arg0 interface{}, args ...interface{}) {
	if !log4go.IsDebugEnabled() {
		return
	}
	tid := getTraceId(ctx)

	const (
		lvl = log4go.DEBUG
	)
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first, args...)
	case func() string:
		// Log the closure (no other arguments used)
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first())
	default:
		// Build a format string so that it will be similar to Sprint
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, fmt.Sprint(arg0)+strings.Repeat(" %v", len(args)), args...)
	}
}

func Trace(ctx context.Context, arg0 interface{}, args ...interface{}) {
	if !log4go.IsTraceEnabled() {
		return
	}
	tid := getTraceId(ctx)

	const (
		lvl = log4go.TRACE
	)
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first, args...)
	case func() string:
		// Log the closure (no other arguments used)
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first())
	default:
		// Build a format string so that it will be similar to Sprint
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, fmt.Sprint(arg0)+strings.Repeat(" %v", len(args)), args...)
	}
}

func Info(ctx context.Context, arg0 interface{}, args ...interface{}) {
	if !log4go.IsInfoEnabled() {
		return
	}
	tid := getTraceId(ctx)

	const (
		lvl = log4go.INFO
	)
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first, args...)
	case func() string:
		// Log the closure (no other arguments used)
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first())
	default:
		// Build a format string so that it will be similar to Sprint
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, fmt.Sprint(arg0)+strings.Repeat(" %v", len(args)), args...)
	}
}

func Warn(ctx context.Context, arg0 interface{}, args ...interface{}) error {
	if !log4go.IsWarnEnabled() {
		return nil
	}
	tid := getTraceId(ctx)

	const (
		lvl = log4go.WARNING
	)

	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, first, args...)
		return errors.New(fmt.Sprintf(first, args...))
	case func() string:
		// Log the closure (no other arguments used)
		str := first()
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, str)
		return errors.New(str)
	default:
		// Build a format string so that it will be similar to Sprint
		log4go.Global.IntLogTraceId(tid, nil, DEEP, lvl, fmt.Sprint(first)+strings.Repeat(" %v", len(args)), args...)
		return errors.New(fmt.Sprint(first) + fmt.Sprintf(strings.Repeat(" %v", len(args)), args...))
	}
	return nil
}

func Error(ctx context.Context, arg0 interface{}, args ...interface{}) error {
	if !log4go.IsErrorEnabled() {
		return nil
	}
	tid := getTraceId(ctx)

	const (
		lvl = log4go.ERROR
	)
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, first, args...)
		return errors.New(fmt.Sprintf(first, args...))
	case func() string:
		// Log the closure (no other arguments used)
		str := first()
		log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, str)
		return errors.New(str)
	default:
		// Build a format string so that it will be similar to Sprint
		log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, fmt.Sprint(first)+strings.Repeat(" %v", len(args)), args...)
		return errors.New(fmt.Sprint(first) + fmt.Sprintf(strings.Repeat(" %v", len(args)), args...))

	}
}

func Critical(ctx context.Context, arg0 interface{}, args ...interface{}) error {
	tid := getTraceId(ctx)

	const (
		lvl = log4go.CRITICAL
	)
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		msg := fmt.Sprintf("%s\n%s", fmt.Sprintf(first, args...), log4go.CallStack(DEEP))
		log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, msg)
		return errors.New(fmt.Sprintf(first, args...))
	case func() string:
		str := first()
		log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, "%s\n%s", str, log4go.CallStack(DEEP))
		return errors.New(str)
	case func(interface{}) string:
		str := first(args[0])
		log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, "%s\n%s", str, log4go.CallStack(DEEP))
		return errors.New(str)
	default:
		// Build a format string so that it will be similar to Sprint
		msg := fmt.Sprintf("%s\n%s", fmt.Sprint(first)+fmt.Sprintf(strings.Repeat(" %v", len(args)), args...), log4go.CallStack(DEEP))
		log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, msg)
		return errors.New(fmt.Sprint(first) + fmt.Sprintf(strings.Repeat(" %v", len(args)), args...))
	}
	return nil
}

func Recover(ctx context.Context, arg0 interface{}, args ...interface{}) {
	tid := getTraceId(ctx)

	const (
		lvl = log4go.CRITICAL
	)
	if err := recover(); err != nil {
		switch first := arg0.(type) {
		case func(interface{}) string:
			// the recovered err will pass to this func
			//Critical(arg0, append([]interface{}{err}, args)...)
			log4go.Global.IntLogTraceId(tid, ctx, DEEP+2, lvl, "%s\n%s", first(err), log4go.CallStack(DEEP+2))
		case string:
			//Critical(a+"\n%v", append(args, err)...)
			msg := fmt.Sprintf("%s\n%s", fmt.Sprintf(first, args...), log4go.CallStack(DEEP))
			log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, msg)
		default:
			//Critical(arg0, append(args, err)...)
			msg := fmt.Sprintf("%s\n%s", fmt.Sprint(first)+fmt.Sprintf(strings.Repeat(" %v", len(args)), args...), log4go.CallStack(DEEP))
			log4go.Global.IntLogTraceId(tid, ctx, DEEP, lvl, msg)
		}
	}
}

func IsFinestEnabled() bool {
	return log4go.IsLevelEnabled(log4go.FINEST)
}

func IsFineEnabled() bool {
	return log4go.IsLevelEnabled(log4go.FINE)
}

func IsDebugEnabled() bool {
	return log4go.IsLevelEnabled(log4go.DEBUG)
}

func IsTraceEnabled() bool {
	return log4go.IsLevelEnabled(log4go.TRACE)
}

func IsInfoEnabled() bool {
	return log4go.IsLevelEnabled(log4go.INFO)
}

func IsWarnEnabled() bool {
	return log4go.IsLevelEnabled(log4go.WARNING)
}

func IsErrorEnabled() bool {
	return log4go.IsLevelEnabled(log4go.ERROR)
}
