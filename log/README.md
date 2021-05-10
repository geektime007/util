

# log
`import "github.com/geektime007/util/log"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package log implements leveled logging.

Example


```go
package main

import (
	"github.com/geektime007/util/log"
)

var logger = log.Get("ExampleName")

func main() {
	logger.SetLevel(log.INFO)
	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Warnf("This is a number %v", 1)
}
```



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Colored(color string, text string) string](#Colored)
* [type Logger](#Logger)
  * [func Get(name string) *Logger](#Get)
  * [func (l *Logger) Debug(a ...interface{}) error](#Logger.Debug)
  * [func (l *Logger) Debugf(format string, a ...interface{}) error](#Logger.Debugf)
  * [func (l *Logger) Disable()](#Logger.Disable)
  * [func (l *Logger) DisableCallerSourceLogging()](#Logger.DisableCallerSourceLogging)
  * [func (l *Logger) Enable()](#Logger.Enable)
  * [func (l *Logger) Error(a ...interface{}) error](#Logger.Error)
  * [func (l *Logger) Errorf(format string, a ...interface{}) error](#Logger.Errorf)
  * [func (l *Logger) Fatal(a ...interface{})](#Logger.Fatal)
  * [func (l *Logger) Fatalf(format string, a ...interface{})](#Logger.Fatalf)
  * [func (l *Logger) Info(a ...interface{}) error](#Logger.Info)
  * [func (l *Logger) Infof(format string, a ...interface{}) error](#Logger.Infof)
  * [func (l *Logger) SetCallerDepth(callerDepth int)](#Logger.SetCallerDepth)
  * [func (l *Logger) SetColored(b bool)](#Logger.SetColored)
  * [func (l *Logger) SetLevel(level int)](#Logger.SetLevel)
  * [func (l *Logger) SetMeta(name string, value interface{})](#Logger.SetMeta)
  * [func (l *Logger) SetNameLength(n int)](#Logger.SetNameLength)
  * [func (l *Logger) SetWriter(w io.Writer)](#Logger.SetWriter)
  * [func (l *Logger) Warn(a ...interface{}) error](#Logger.Warn)
  * [func (l *Logger) Warnf(format string, a ...interface{}) error](#Logger.Warnf)


#### <a name="pkg-files">Package files</a>
[log.go](/src/github.com/geektime007/util/log/log.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    DEBUG int = iota
    INFO
    WARN
    ERROR
)
```
Level

``` go
const (
    DefaultCallerDepth = 2
)
```
default




## <a name="Colored">func</a> [Colored](./log.go#L213?s=4649:4695)
``` go
func Colored(color string, text string) string
```
Colored returns text in color.




## <a name="Logger">type</a> [Logger](./log.go#L61?s=947:1279)
``` go
type Logger struct {
    // contains filtered or unexported fields
}

```
Logger abstraction.







### <a name="Get">func</a> [Get](./log.go#L74?s=1310:1339)
``` go
func Get(name string) *Logger
```
New creates a new Logger.





### <a name="Logger.Debug">func</a> (\*Logger) [Debug](./log.go#L161?s=3169:3215)
``` go
func (l *Logger) Debug(a ...interface{}) error
```
Debug logs message with level DEBUG.




### <a name="Logger.Debugf">func</a> (\*Logger) [Debugf](./log.go#L187?s=3832:3894)
``` go
func (l *Logger) Debugf(format string, a ...interface{}) error
```
Debugf formats and logs message with level DEBUG.




### <a name="Logger.Disable">func</a> (\*Logger) [Disable](./log.go#L151?s=3006:3032)
``` go
func (l *Logger) Disable()
```
Disable the logging.




### <a name="Logger.DisableCallerSourceLogging">func</a> (\*Logger) [DisableCallerSourceLogging](./log.go#L141?s=2753:2798)
``` go
func (l *Logger) DisableCallerSourceLogging()
```
DisableCallerSourceLogging disables the logging for caller source.
This sets to true by default.




### <a name="Logger.Enable">func</a> (\*Logger) [Enable](./log.go#L156?s=3080:3105)
``` go
func (l *Logger) Enable()
```
Enable the logging.




### <a name="Logger.Error">func</a> (\*Logger) [Error](./log.go#L176?s=3554:3600)
``` go
func (l *Logger) Error(a ...interface{}) error
```
Error logs message with level ERROR.




### <a name="Logger.Errorf">func</a> (\*Logger) [Errorf](./log.go#L202?s=4331:4393)
``` go
func (l *Logger) Errorf(format string, a ...interface{}) error
```
Errorf formats and logs message with level ERROR.




### <a name="Logger.Fatal">func</a> (\*Logger) [Fatal](./log.go#L181?s=3689:3729)
``` go
func (l *Logger) Fatal(a ...interface{})
```
Fatal and logs message with level FATAL.




### <a name="Logger.Fatalf">func</a> (\*Logger) [Fatalf](./log.go#L207?s=4500:4556)
``` go
func (l *Logger) Fatalf(format string, a ...interface{})
```
Fatalf formats and logs message with level FATAL.




### <a name="Logger.Info">func</a> (\*Logger) [Info](./log.go#L166?s=3298:3343)
``` go
func (l *Logger) Info(a ...interface{}) error
```
Info logs message with level INFO.




### <a name="Logger.Infof">func</a> (\*Logger) [Infof](./log.go#L192?s=3999:4060)
``` go
func (l *Logger) Infof(format string, a ...interface{}) error
```
Infof formats and logs message with level INFO.




### <a name="Logger.SetCallerDepth">func</a> (\*Logger) [SetCallerDepth](./log.go#L135?s=2567:2615)
``` go
func (l *Logger) SetCallerDepth(callerDepth int)
```
SetCallerDepth sets the caller depth for this logger.




### <a name="Logger.SetColored">func</a> (\*Logger) [SetColored](./log.go#L117?s=2158:2193)
``` go
func (l *Logger) SetColored(b bool)
```
SetColored sets the color enability.




### <a name="Logger.SetLevel">func</a> (\*Logger) [SetLevel](./log.go#L122?s=2250:2286)
``` go
func (l *Logger) SetLevel(level int)
```
SetLevel sets the logging level.




### <a name="Logger.SetMeta">func</a> (\*Logger) [SetMeta](./log.go#L146?s=2894:2950)
``` go
func (l *Logger) SetMeta(name string, value interface{})
```
Add logging meta adds logging meta to the logger.




### <a name="Logger.SetNameLength">func</a> (\*Logger) [SetNameLength](./log.go#L132?s=2450:2487)
``` go
func (l *Logger) SetNameLength(n int)
```
SetNameLength sets the name length.




### <a name="Logger.SetWriter">func</a> (\*Logger) [SetWriter](./log.go#L127?s=2357:2396)
``` go
func (l *Logger) SetWriter(w io.Writer)
```
SetWriter sets the writer.




### <a name="Logger.Warn">func</a> (\*Logger) [Warn](./log.go#L171?s=3425:3470)
``` go
func (l *Logger) Warn(a ...interface{}) error
```
Warn logs message with level WARN.




### <a name="Logger.Warnf">func</a> (\*Logger) [Warnf](./log.go#L197?s=4164:4225)
``` go
func (l *Logger) Warnf(format string, a ...interface{}) error
```
Warnf formats and logs message with level WARN.








