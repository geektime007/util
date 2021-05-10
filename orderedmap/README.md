

# orderedmap
`import "github.com/geektime007/util/orderedmap"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package orderedmap implements an ordered map.

Example


```go
m := orderedmap.NewOrderedMap()

// Set key with value.
m.Set("key", 1)

// Get value by key.
val, ok := m.Get("key")

// Get keys in insertion order.
for _, key := range m.Keys() {
	val, ok := m.Get(key)
}
```



## <a name="pkg-index">Index</a>
* [type OrderedIntMap](#OrderedIntMap)
  * [func NewOrderedIntMap() *OrderedIntMap](#NewOrderedIntMap)
  * [func (m *OrderedIntMap) Get(key int) (val interface{}, ok bool)](#OrderedIntMap.Get)
  * [func (m *OrderedIntMap) Keys() []int](#OrderedIntMap.Keys)
  * [func (m *OrderedIntMap) Len() int](#OrderedIntMap.Len)
  * [func (m *OrderedIntMap) Pop(key int) (val interface{}, ok bool)](#OrderedIntMap.Pop)
  * [func (m *OrderedIntMap) Set(key int, val interface{})](#OrderedIntMap.Set)
  * [func (m *OrderedIntMap) Values() []interface{}](#OrderedIntMap.Values)
* [type OrderedMap](#OrderedMap)
  * [func NewOrderedMap() *OrderedMap](#NewOrderedMap)
  * [func (m *OrderedMap) Get(key string) (val interface{}, ok bool)](#OrderedMap.Get)
  * [func (m *OrderedMap) Keys() []string](#OrderedMap.Keys)
  * [func (m *OrderedMap) Len() int](#OrderedMap.Len)
  * [func (m *OrderedMap) Pop(key string) (val interface{}, ok bool)](#OrderedMap.Pop)
  * [func (m *OrderedMap) Set(key string, val interface{})](#OrderedMap.Set)
  * [func (m *OrderedMap) Values() []interface{}](#OrderedMap.Values)
* [type OrderedUintMap](#OrderedUintMap)
  * [func NewOrderedUintMap() *OrderedUintMap](#NewOrderedUintMap)
  * [func (m *OrderedUintMap) Get(key uint) (val interface{}, ok bool)](#OrderedUintMap.Get)
  * [func (m *OrderedUintMap) Keys() []uint](#OrderedUintMap.Keys)
  * [func (m *OrderedUintMap) Len() int](#OrderedUintMap.Len)
  * [func (m *OrderedUintMap) Pop(key uint) (val interface{}, ok bool)](#OrderedUintMap.Pop)
  * [func (m *OrderedUintMap) Set(key uint, val interface{})](#OrderedUintMap.Set)
  * [func (m *OrderedUintMap) Values() []interface{}](#OrderedUintMap.Values)


#### <a name="pkg-files">Package files</a>
[orderedmap.go](/src/github.com/geektime007/util/orderedmap/orderedmap.go)






## <a name="OrderedIntMap">type</a> [OrderedIntMap](./orderedmap.go#L80?s=1678:1745)
``` go
type OrderedIntMap struct {
    // contains filtered or unexported fields
}

```
OrderedIntMap is an ordered map with integers as keys.







### <a name="NewOrderedIntMap">func</a> [NewOrderedIntMap](./orderedmap.go#L86?s=1796:1834)
``` go
func NewOrderedIntMap() *OrderedIntMap
```
NewOrderedIntMap returns a new OrderedIntMap.





### <a name="OrderedIntMap.Get">func</a> (\*OrderedIntMap) [Get](./orderedmap.go#L115?s=2435:2498)
``` go
func (m *OrderedIntMap) Get(key int) (val interface{}, ok bool)
```
Get a value by key.




### <a name="OrderedIntMap.Keys">func</a> (\*OrderedIntMap) [Keys](./orderedmap.go#L94?s=1971:2007)
``` go
func (m *OrderedIntMap) Keys() []int
```
Keys returns the keys in insertion order.




### <a name="OrderedIntMap.Len">func</a> (\*OrderedIntMap) [Len](./orderedmap.go#L134?s=2817:2850)
``` go
func (m *OrderedIntMap) Len() int
```
Len returns the length of this map.




### <a name="OrderedIntMap.Pop">func</a> (\*OrderedIntMap) [Pop](./orderedmap.go#L121?s=2555:2618)
``` go
func (m *OrderedIntMap) Pop(key int) (val interface{}, ok bool)
```
Pop a value by key.




### <a name="OrderedIntMap.Set">func</a> (\*OrderedIntMap) [Set](./orderedmap.go#L106?s=2274:2327)
``` go
func (m *OrderedIntMap) Set(key int, val interface{})
```
Set a key into this map.




### <a name="OrderedIntMap.Values">func</a> (\*OrderedIntMap) [Values](./orderedmap.go#L97?s=2076:2122)
``` go
func (m *OrderedIntMap) Values() []interface{}
```
Values returns the values in insertion order.




## <a name="OrderedMap">type</a> [OrderedMap](./orderedmap.go#L23?s=438:508)
``` go
type OrderedMap struct {
    // contains filtered or unexported fields
}

```
OrderedMap is an ordered map with strings as keys.







### <a name="NewOrderedMap">func</a> [NewOrderedMap](./orderedmap.go#L29?s=553:585)
``` go
func NewOrderedMap() *OrderedMap
```
NewOrderedMap returns a new OrderedMap.





### <a name="OrderedMap.Get">func</a> (\*OrderedMap) [Get](./orderedmap.go#L58?s=1186:1249)
``` go
func (m *OrderedMap) Get(key string) (val interface{}, ok bool)
```
Get a value by key.




### <a name="OrderedMap.Keys">func</a> (\*OrderedMap) [Keys](./orderedmap.go#L37?s=725:761)
``` go
func (m *OrderedMap) Keys() []string
```
Keys returns the keys in insertion order.




### <a name="OrderedMap.Len">func</a> (\*OrderedMap) [Len](./orderedmap.go#L77?s=1568:1598)
``` go
func (m *OrderedMap) Len() int
```
Len returns the length of this map.




### <a name="OrderedMap.Pop">func</a> (\*OrderedMap) [Pop](./orderedmap.go#L64?s=1306:1369)
``` go
func (m *OrderedMap) Pop(key string) (val interface{}, ok bool)
```
Pop a value by key.




### <a name="OrderedMap.Set">func</a> (\*OrderedMap) [Set](./orderedmap.go#L49?s=1025:1078)
``` go
func (m *OrderedMap) Set(key string, val interface{})
```
Set a key into this map.




### <a name="OrderedMap.Values">func</a> (\*OrderedMap) [Values](./orderedmap.go#L40?s=830:873)
``` go
func (m *OrderedMap) Values() []interface{}
```
Values returns the values in insertion order.




## <a name="OrderedUintMap">type</a> [OrderedUintMap](./orderedmap.go#L137?s=2940:3010)
``` go
type OrderedUintMap struct {
    // contains filtered or unexported fields
}

```
OrderedUintMap is an ordered map with unsigned integers as keys.







### <a name="NewOrderedUintMap">func</a> [NewOrderedUintMap](./orderedmap.go#L143?s=3063:3103)
``` go
func NewOrderedUintMap() *OrderedUintMap
```
NewOrderedUintMap returns a new OrderedUintMap.





### <a name="OrderedUintMap.Get">func</a> (\*OrderedUintMap) [Get](./orderedmap.go#L172?s=3712:3777)
``` go
func (m *OrderedUintMap) Get(key uint) (val interface{}, ok bool)
```
Get a value by key.




### <a name="OrderedUintMap.Keys">func</a> (\*OrderedUintMap) [Keys](./orderedmap.go#L151?s=3243:3281)
``` go
func (m *OrderedUintMap) Keys() []uint
```
Keys returns the keys in insertion order.




### <a name="OrderedUintMap.Len">func</a> (\*OrderedUintMap) [Len](./orderedmap.go#L191?s=4098:4132)
``` go
func (m *OrderedUintMap) Len() int
```
Len returns the length of this map.




### <a name="OrderedUintMap.Pop">func</a> (\*OrderedUintMap) [Pop](./orderedmap.go#L178?s=3834:3899)
``` go
func (m *OrderedUintMap) Pop(key uint) (val interface{}, ok bool)
```
Pop a value by key.




### <a name="OrderedUintMap.Set">func</a> (\*OrderedUintMap) [Set](./orderedmap.go#L163?s=3549:3604)
``` go
func (m *OrderedUintMap) Set(key uint, val interface{})
```
Set a key into this map.




### <a name="OrderedUintMap.Values">func</a> (\*OrderedUintMap) [Values](./orderedmap.go#L154?s=3350:3397)
``` go
func (m *OrderedUintMap) Values() []interface{}
```
Values returns the values in insertion order.








