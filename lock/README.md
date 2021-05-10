

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



## <a name="LogFunc">func</a> [LogFunc](./mutex.go#L36?s=848:896)
``` go
func LogFunc(fn func(m MutexWrapper, op string))
```
Register a global error logger.




## <a name="Mutex">type</a> [Mutex](./mutex.go#L49?s=1209:1477)
``` go
type Mutex struct {
    // contains filtered or unexported fields
}

```
Mutex is the wrapper around sync.Mutex.







### <a name="NewMutex">func</a> [NewMutex](./mutex.go#L60?s=1515:1571)
``` go
func NewMutex(name string, timeout time.Duration) *Mutex
```
NewMutex constructs a new mutex.





### <a name="Mutex.Lock">func</a> (\*Mutex) [Lock](./mutex.go#L79?s=1892:1914)
``` go
func (m *Mutex) Lock()
```
Lock wraps sync.Mutex.Lock().




### <a name="Mutex.Name">func</a> (\*Mutex) [Name](./mutex.go#L69?s=1699:1728)
``` go
func (m *Mutex) Name() string
```
Name implements MutexWrapper.Name().




### <a name="Mutex.Timeout">func</a> (\*Mutex) [Timeout](./mutex.go#L74?s=1796:1835)
``` go
func (m *Mutex) Timeout() time.Duration
```
Timeout implements Mutex.Wrapper.Timeout().




### <a name="Mutex.Unlock">func</a> (\*Mutex) [Unlock](./mutex.go#L99?s=2296:2320)
``` go
func (m *Mutex) Unlock()
```
Unlock wraps sync.Mutex.Unlock().




## <a name="MutexWrapper">type</a> [MutexWrapper](./mutex.go#L41?s=983:1164)
``` go
type MutexWrapper interface {
    // Name returns the name of this mutex wrapper.
    Name() string
    // Timeout returns the timeout value of this mutex wrapper.
    Timeout() time.Duration
}
```
MutexWrapper is an interface implemented by Mutex and RWMutex.










## <a name="RWMutex">type</a> [RWMutex](./mutex.go#L104?s=2386:2658)
``` go
type RWMutex struct {
    // contains filtered or unexported fields
}

```
RWMutex is the wrapper around sync.RWMutex







### <a name="NewRWMutex">func</a> [NewRWMutex](./mutex.go#L115?s=2700:2760)
``` go
func NewRWMutex(name string, timeout time.Duration) *RWMutex
```
NewRWMutex constructs a new rwmutex.





### <a name="RWMutex.Lock">func</a> (\*RWMutex) [Lock](./mutex.go#L134?s=3091:3115)
``` go
func (m *RWMutex) Lock()
```
Lock wraps sync.RWMutex.Lock().




### <a name="RWMutex.Name">func</a> (\*RWMutex) [Name](./mutex.go#L124?s=2892:2923)
``` go
func (m *RWMutex) Name() string
```
Name implements MutexWrapper.Name().




### <a name="RWMutex.RLock">func</a> (\*RWMutex) [RLock](./mutex.go#L159?s=3584:3609)
``` go
func (m *RWMutex) RLock()
```
RLock wraps sync.RWMutex.RLock().




### <a name="RWMutex.RUnlock">func</a> (\*RWMutex) [RUnlock](./mutex.go#L179?s=4000:4027)
``` go
func (m *RWMutex) RUnlock()
```
RUnlock wraps sync.RWMutex.RUnlock().




### <a name="RWMutex.Timeout">func</a> (\*RWMutex) [Timeout](./mutex.go#L129?s=2991:3032)
``` go
func (m *RWMutex) Timeout() time.Duration
```
Timeout implements Mutex.Wrapper.Timeout().




### <a name="RWMutex.Unlock">func</a> (\*RWMutex) [Unlock](./mutex.go#L154?s=3501:3527)
``` go
func (m *RWMutex) Unlock()
```
Unlock wraps sync.RWMutex.Unlock().








