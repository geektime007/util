

# lock
`import "github.com/geektime007/util/lock"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package lock wraps the sync.Mutex and sync.RWMutex to log messages on
deadlocks.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func LogFunc(fn func(m MutexWrapper, op string))](#LogFunc)
* [type Mutex](#Mutex)
  * [func NewMutex(name string, timeout time.Duration) *Mutex](#NewMutex)
  * [func (m *Mutex) Lock()](#Mutex.Lock)
  * [func (m *Mutex) Name() string](#Mutex.Name)
  * [func (m *Mutex) Timeout() time.Duration](#Mutex.Timeout)
  * [func (m *Mutex) Unlock()](#Mutex.Unlock)
* [type MutexWrapper](#MutexWrapper)
* [type RWMutex](#RWMutex)
  * [func NewRWMutex(name string, timeout time.Duration) *RWMutex](#NewRWMutex)
  * [func (m *RWMutex) Lock()](#RWMutex.Lock)
  * [func (m *RWMutex) Name() string](#RWMutex.Name)
  * [func (m *RWMutex) RLock()](#RWMutex.RLock)
  * [func (m *RWMutex) RUnlock()](#RWMutex.RUnlock)
  * [func (m *RWMutex) Timeout() time.Duration](#RWMutex.Timeout)
  * [func (m *RWMutex) Unlock()](#RWMutex.Unlock)


#### <a name="pkg-files">Package files</a>
[mutex.go](/src/github.com/geektime007/util/lock/mutex.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    OpMutexLock    = "Mutex.Lock"
    OpRWMutexLock  = "RWMutex.Lock"
    OpRWMutexRLock = "RWMutex.RLock"
)
```
Constants of operations.


## <a name="pkg-variables">Variables</a>
``` go
var DefaultLogFunc = func(m MutexWrapper, op string) {
    logger.Errorf("Timeout! Wait %dms to acquire lock %s, blocked operation is %s, callers stack:\n%s",
        int(m.Timeout().Nanoseconds()/1000000), m.Name(), op, debug.Stack())
}
```
Default log function



## <a name="LogFunc">func</a> [LogFunc](.//target/mutex.go#L35?s=796:844)
``` go
func LogFunc(fn func(m MutexWrapper, op string))
```
Register a global error logger.




## <a name="Mutex">type</a> [Mutex](.//target/mutex.go#L48?s=1157:1425)
``` go
type Mutex struct {
    // contains filtered or unexported fields
}

```
Mutex is the wrapper around sync.Mutex.







### <a name="NewMutex">func</a> [NewMutex](.//target/mutex.go#L59?s=1463:1519)
``` go
func NewMutex(name string, timeout time.Duration) *Mutex
```
NewMutex constructs a new mutex.





### <a name="Mutex.Lock">func</a> (\*Mutex) [Lock](.//target/mutex.go#L78?s=1840:1862)
``` go
func (m *Mutex) Lock()
```
Lock wraps sync.Mutex.Lock().




### <a name="Mutex.Name">func</a> (\*Mutex) [Name](.//target/mutex.go#L68?s=1647:1676)
``` go
func (m *Mutex) Name() string
```
Name implements MutexWrapper.Name().




### <a name="Mutex.Timeout">func</a> (\*Mutex) [Timeout](.//target/mutex.go#L73?s=1744:1783)
``` go
func (m *Mutex) Timeout() time.Duration
```
Timeout implements Mutex.Wrapper.Timeout().




### <a name="Mutex.Unlock">func</a> (\*Mutex) [Unlock](.//target/mutex.go#L98?s=2244:2268)
``` go
func (m *Mutex) Unlock()
```
Unlock wraps sync.Mutex.Unlock().




## <a name="MutexWrapper">type</a> [MutexWrapper](.//target/mutex.go#L40?s=931:1112)
``` go
type MutexWrapper interface {
    // Name returns the name of this mutex wrapper.
    Name() string
    // Timeout returns the timeout value of this mutex wrapper.
    Timeout() time.Duration
}
```
MutexWrapper is an interface implemented by Mutex and RWMutex.










## <a name="RWMutex">type</a> [RWMutex](.//target/mutex.go#L103?s=2334:2606)
``` go
type RWMutex struct {
    // contains filtered or unexported fields
}

```
RWMutex is the wrapper around sync.RWMutex







### <a name="NewRWMutex">func</a> [NewRWMutex](.//target/mutex.go#L114?s=2648:2708)
``` go
func NewRWMutex(name string, timeout time.Duration) *RWMutex
```
NewRWMutex constructs a new rwmutex.





### <a name="RWMutex.Lock">func</a> (\*RWMutex) [Lock](.//target/mutex.go#L133?s=3039:3063)
``` go
func (m *RWMutex) Lock()
```
Lock wraps sync.RWMutex.Lock().




### <a name="RWMutex.Name">func</a> (\*RWMutex) [Name](.//target/mutex.go#L123?s=2840:2871)
``` go
func (m *RWMutex) Name() string
```
Name implements MutexWrapper.Name().




### <a name="RWMutex.RLock">func</a> (\*RWMutex) [RLock](.//target/mutex.go#L158?s=3532:3557)
``` go
func (m *RWMutex) RLock()
```
RLock wraps sync.RWMutex.RLock().




### <a name="RWMutex.RUnlock">func</a> (\*RWMutex) [RUnlock](.//target/mutex.go#L178?s=3948:3975)
``` go
func (m *RWMutex) RUnlock()
```
RUnlock wraps sync.RWMutex.RUnlock().




### <a name="RWMutex.Timeout">func</a> (\*RWMutex) [Timeout](.//target/mutex.go#L128?s=2939:2980)
``` go
func (m *RWMutex) Timeout() time.Duration
```
Timeout implements Mutex.Wrapper.Timeout().




### <a name="RWMutex.Unlock">func</a> (\*RWMutex) [Unlock](.//target/mutex.go#L153?s=3449:3475)
``` go
func (m *RWMutex) Unlock()
```
Unlock wraps sync.RWMutex.Unlock().








