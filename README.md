
# Go Test

[![GoDoc](https://godoc.org/github.com/jmervine/GoT?status.png)](https://godoc.org/github.com/jmervine/GoT)

<center>![GoT](http://cdn.mervine.net/GoT.jpg)</center>

#### Simple Assertion wrapper for Go's built in "testing" package.

GoT is designed to be as simple and unintrusive as possible while adding basic Assert and Refute methods to assist in writing clean and clean tests quickly.

##### See [examples](_example) for detailed assertion usage.

## [Documentation](https://godoc.org/github.com/jmervine/GoT)

```go
import "github.com/jmervine/GoT"
```

Package GoT is a simple assertion wrapper for Go's built in "testing" package,
which is designed to be as simple and unintrusive as possible while adding basic
Assert and Refute methods to assist in writing clean and clean tests quickly.

Setup Examples:

    // ------------------------------------------------ //
    // Simple:
    import (
        "github.com/jmervine/GoT"
        "testing"
    )
    func TestFoo(T *testing.T) {
       t := GoT.Go(T)
       t.Assert(true, "should be true")
    }
    // ------------------------------------------------ //
    // Global and authors perfered:
    import (
        . "github.com/jmervine/GoT"
        "testing"
    )
    func TestFoo(T *testing.T) {
       Go(T).Assert(true, "should be true")
       Go(T).Refute(false, "should be false")
    }
    // ------------------------------------------------ //
    // Global (but ugly):
    //
    import (
        "github.com/jmervine/GoT"
        "testing"
    )
    var Go (func (*testing.T) *GoT.GoT) = GoT.Go
    func TestFoo(T *testing.T) {
       Go(T).Assert(true, "should be true")
       Go(T).Refute(false, "should be false")
    }


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
Go wraps "testing" to apply assertions.

##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).Assert(true, "should be true")
	Go(T).Refute(false, "should not be true")

#### Assert

```go
func (t *GoT) Assert(a bool, m string)
```
Assert checks for true.



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).Assert(true, "should be true")


#### AssertDeepEqual

```go
func (t *GoT) AssertDeepEqual(a interface{}, b interface{}, m string)
```
AssertDeepEqual checks for simlarity (see: reflect.DeepEqual).



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	a1 := []string{"a"}
	a2 := []string{"a"}
	
	Go(T).AssertDeepEqual(a1, a2, "should deep equal")


#### AssertEqual

```go
func (t *GoT) AssertEqual(a interface{}, b interface{}, m string)
```
AssertEqual check for equality.



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).AssertEqual(1, 1, "should equal")


#### AssertLength

```go
func (t *GoT) AssertLength(a interface{}, n int, m string)
```
AssertLength checks for length equal to `n int`. If the type passed cannot be
checked for length and error is logged stating as such.



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).AssertLength("a", 1, "should be length")


#### AssertNil

```go
func (t *GoT) AssertNil(a interface{}, m string)
```
AssertNil checks for nil.



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).AssertNil(nil, "should be nil")


#### Refute

```go
func (t *GoT) Refute(a bool, m string)
```
Refute checks for false.



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).Refute(false, "should not be true")


#### RefuteDeepEqual

```go
func (t *GoT) RefuteDeepEqual(a interface{}, b interface{}, m string)
```
RefuteDeepEqual checks for no simlarity (see: reflect.DeepEqual).



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	a1 := []string{"a"}
	a2 := []string{"b"}
	
	Go(T).RefuteDeepEqual(a1, a2, "should not deep equal")


#### RefuteEqual

```go
func (t *GoT) RefuteEqual(a interface{}, b interface{}, m string)
```
RefuteEqual checks for inequality.



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).RefuteEqual(1, 2, "should not equal")


#### RefuteLength

```go
func (t *GoT) RefuteLength(a interface{}, n int, m string)
```
RefuteLength checks for length not equal to `n int`. If the type passed cannot
be checked for length this assertion will pass. (I'm not sure this is the best
way to handle this, feedback is welcome.)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).RefuteLength("a", 0, "should not be length")


#### RefuteNil

```go
func (t *GoT) RefuteNil(a interface{}, m string)
```
RefuteNil checks for not nil.



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).RefuteNil(1, "should not be nil")



