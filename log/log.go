// Package log implements leveled logging.
//
// Example
//
//	package main
//
//	import (
//		"github.com/geektime007/util/log"
//	)
//
//	var logger = log.Get("ExampleName")
//
//	func main() {
//		logger.SetLevel(log.INFO)
//		logger.Debug("This is a debug message")
//		logger.Info("This is a info message")
//		logger.Warn("This is a warning message")
//		logger.Error("This is an error message")
//		logger.Warnf("This is a number %v", 1)
//	}
//
package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/geektime007/util/log/journal"
	"github.com/geektime007/util/orderedmap"
	"github.com/willf/pad"
)

// Level
const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
)

// Level to journal level
var journalLevel = map[int]journal.Priority{
	INFO:  journal.PriInfo,
	DEBUG: journal.PriDebug,
	WARN:  journal.PriWarning,
	ERROR: journal.PriErr,
}

// default
const (
	DefaultCallerDepth = 2
)

// Global registry
var m = sync.Mutex{}
var registry = make(map[string]*Logger, 0)

// Level name
var levelNames = [4]string{"DEBUG", "INFO", "WARN", "ERROR"}

// Logger abstraction.
type Logger struct {
	name                      string
	nameLength                int
	level                     int
	w                         io.Writer
	colored                   bool
	enabled                   bool
	callerDepth               int
	enableCallerSourceLogging bool
	metas                     *orderedmap.OrderedMap
}

// get log registry
func GetRegistry() map[string]*Logger {
	return registry
}

// New creates a new Logger.
func Get(name string) *Logger {
	m.Lock()
	defer m.Unlock()
	l, ok := registry[name]
	if ok {
		return l
	}
	l = &Logger{
		name:                      name,
		nameLength:                0,
		level:                     INFO,
		w:                         os.Stdout,
		colored:                   true,
		enabled:                   true,
		callerDepth:               DefaultCallerDepth,
		enableCallerSourceLogging: true,
		metas:                     orderedmap.NewOrderedMap(),
	}
	registry[name] = l
	return l
}

// colors to ansi code map
var colors = map[string]int{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,
}

// levelColors
var levelColors = map[int]string{
	DEBUG: "blue",
	INFO:  "green",
	WARN:  "yellow",
	ERROR: "red",
}

// SetColored sets the color enability.
func (l *Logger) SetColored(b bool) {
	l.colored = b
}

// SetLevel sets the logging level.
func (l *Logger) SetLevel(level int) {
	l.level = level % len(levelNames)
}

// GetLevel gets the logging level.
func (l *Logger) GetLevel() int {
	return l.level
}

// SetWriter sets the writer.
func (l *Logger) SetWriter(w io.Writer) {
	l.w = w
}

// SetNameLength sets the name length.
func (l *Logger) SetNameLength(n int) { l.nameLength = n }

// SetCallerDepth sets the caller depth for this logger.
func (l *Logger) SetCallerDepth(callerDepth int) {
	l.callerDepth = callerDepth
}

// DisableCallerSourceLogging disables the logging for caller source.
// This sets to true by default.
func (l *Logger) DisableCallerSourceLogging() {
	l.enableCallerSourceLogging = false
}

// Add logging meta adds logging meta to the logger.
func (l *Logger) SetMeta(name string, value interface{}) {
	l.metas.Set(name, value)
}

// Disable the logging.
func (l *Logger) Disable() {
	l.enabled = false
}

// Enable the logging.
func (l *Logger) Enable() {
	l.enabled = true
}

// Debug logs message with level DEBUG.
func (l *Logger) Debug(a ...interface{}) error {
	return l.log(DEBUG, fmt.Sprint(a...))
}

// Info logs message with level INFO.
func (l *Logger) Info(a ...interface{}) error {
	return l.log(INFO, fmt.Sprint(a...))
}

// Warn logs message with level WARN.
func (l *Logger) Warn(a ...interface{}) error {
	return l.log(WARN, fmt.Sprint(a...))
}

// Error logs message with level ERROR.
func (l *Logger) Error(a ...interface{}) error {
	return l.log(ERROR, fmt.Sprint(a...))
}

// Fatal and logs message with level FATAL.
func (l *Logger) Fatal(a ...interface{}) {
	l.log(ERROR, fmt.Sprint(a...))
	os.Exit(1)
}

// Debugf formats and logs message with level DEBUG.
func (l *Logger) Debugf(format string, a ...interface{}) error {
	return l.log(DEBUG, fmt.Sprintf(format, a...))
}

// Infof formats and logs message with level INFO.
func (l *Logger) Infof(format string, a ...interface{}) error {
	return l.log(INFO, fmt.Sprintf(format, a...))
}

// Warnf formats and logs message with level WARN.
func (l *Logger) Warnf(format string, a ...interface{}) error {
	return l.log(WARN, fmt.Sprintf(format, a...))
}

// Errorf formats and logs message with level ERROR.
func (l *Logger) Errorf(format string, a ...interface{}) error {
	return l.log(ERROR, fmt.Sprintf(format, a...))
}

// Fatalf formats and logs message with level FATAL.
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.log(ERROR, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Colored returns text in color.
func Colored(color string, text string) string {
	return fmt.Sprintf("\033[3%dm%s\033[0m", colors[color], text)
}

// log dose logging.
func (l *Logger) log(level int, msg string) error {
	if l.enabled && level >= l.level {
		// Caller pkg.
		_, fileName, line, _ := runtime.Caller(l.callerDepth)
		pkgName := path.Base(path.Dir(fileName))
		filepath := path.Join(pkgName, path.Base(fileName))
		// Datetime and pid.
		now := time.Now().String()[:19]
		// Message
		levelName := levelNames[level]
		levelName = pad.Right(levelName, 5, " ")
		// Name
		name := l.name
		if l.nameLength > 0 {
			name = pad.Right(l.name, l.nameLength, " ")
		}
		// logging meta
		lst := make([]string, 0)
		for _, key := range l.metas.Keys() {
			value, ok := l.metas.Get(key)
			if ok {
				lst = append(lst, fmt.Sprintf("%s=%v", key, value))
			}
		}
		metasString := strings.Join(lst, " ")

		// Whether to log the caller source.
		var headerString string
		if !l.enableCallerSourceLogging {
			headerString = fmt.Sprintf("[%s %s] %s %s", name, metasString, levelName, now)
		} else {
			headerString = fmt.Sprintf("[%s %s] %s %s %s:%d", name, metasString, levelName, now, filepath, line)
		}

		header := Colored(levelColors[level], headerString)

		if journal.Enabled() {
			logMsg := fmt.Sprintf("%s %s", header, msg)
			journal.Send(logMsg, journalLevel[level], nil)
		} else {
			logMsg := fmt.Sprintf("%s %s\n", header, msg)
			_, err := fmt.Fprintf(l.w, logMsg)
			return err
		}
	}
	return nil
}
