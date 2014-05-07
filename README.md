
# Go Test

<center><img src="http://cdn.mervine.net/GoT.jpg" /></center>

#### Simple Assertion wrapper for Go's built in "testing" package.

GoT is designed to be as simple and unintrusive as possible while adding basic Assert and Refute methods to assist in writing clean and clean tests quickly.

##### See [examples](_example) for detailed assertion usage.

## Documentation

```go
import "github.com/jmervine/goperf"
```

### Types

#### GoT

```go
type GoT struct {
    // contains filtered or unexported fields
}
```



#### Go

```go
func Go(T *testing.T) *GoT
```
Core "testing" wrapper to apply assertions.

##### Example:
	// file: foo_test.go
	// ------------------------------------------------ //
	//
	// Simple:
	//
	// import (
	//     "github.com/jmervine/got"
	//     "testing"
	// )
	//
	// func TestFoo(T *testing.T) {
	//    t := got.Go(T)
	//    t.Assert(true, "should be true")
	// }
	//
	// ------------------------------------------------ //
	//
	// Global:
	//
	// import (
	//     "github.com/jmervine/got"
	//     "testing"
	// )
	//
	// var Go := got.Go
	// func TestFoo(T *testing.T) {
	//    Go(T).Assert(true, "should be true")
	//    Go(T).Refute(false, "should be false")
	// }
	//
	// ------------------------------------------------ //
	//
	// Or:
	//
	// import (
	//     . "github.com/jmervine/got"
	//     "testing"
	// )
	//
	// func TestFoo(T *testing.T) {
	//    Go(T).Assert(true, "should be true")
	//    Go(T).Refute(false, "should be false")
	// }
	//
	// ------------------------------------------------ //

#### Assert

```go
func (t *GoT) Assert(a bool, m string)
```
Check for true.



#### AssertDeepEqual

```go
func (t *GoT) AssertDeepEqual(a interface{}, b interface{}, m string)
```
Check for simlarity (see: reflect.DeepEqual).



#### AssertEqual

```go
func (t *GoT) AssertEqual(a interface{}, b interface{}, m string)
```
Check for equality.



#### AssertLength

```go
func (t *GoT) AssertLength(a interface{}, n int, m string)
```
Check length is n.



#### AssertNil

```go
func (t *GoT) AssertNil(a interface{}, m string)
```
Check for nil.



#### Refute

```go
func (t *GoT) Refute(a bool, m string)
```
Check for false.



#### RefuteDeepEqual

```go
func (t *GoT) RefuteDeepEqual(a interface{}, b interface{}, m string)
```
Check for no simlarity (see: reflect.DeepEqual).



#### RefuteEqual

```go
func (t *GoT) RefuteEqual(a interface{}, b interface{}, m string)
```
Check for inequality.



#### RefuteLength

```go
func (t *GoT) RefuteLength(a interface{}, n int, m string)
```
Check length is not n.



#### RefuteNil

```go
func (t *GoT) RefuteNil(a interface{}, m string)
```
Check for not nil.




