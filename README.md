
# GoTest

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
       t.Assert(true)
    }
    // ------------------------------------------------ //
    // Global and authors perfered:
    import (
        . "github.com/jmervine/GoT"
        "testing"
    )
    func TestFoo(T *testing.T) {
       Go(T).Assert(true)
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
       Go(T).Refute(false)
    }

Note: In assertions, if no message is passed, then a nice message will be
displayed in an effor to be as clear as possible.


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
	
	Go(T).Assert(true)
	Go(T).Assert(true, "should be true")
	
	Go(T).Refute(false)
	Go(T).Refute(false, "should not be true")

#### Assert

```go
func (t *GoT) Assert(args ...interface{})
```
Assert checks for true.

Expects:

    Assert(a bool, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).Assert(true)
	Go(T).Assert(true, "should be true")


#### AssertContains

```go
func (t *GoT) AssertContains(args ...string)
```
AssertContains checks for the existence of the second string inside the first
string.

Expects:

    AssertContains(a string, b string, optional_message string)



#### AssertDeepEqual

```go
func (t *GoT) AssertDeepEqual(args ...interface{})
```
AssertDeepEqual checks for simlarity (see: reflect.DeepEqual).

Expects:

    AssertDeepEqual(a interface{}, b interface{}, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	a1 := []string{"a"}
	a2 := []string{"a"}
	
	Go(T).AssertDeepEqual(a1, a2)
	Go(T).AssertDeepEqual(a1, a2, "should deep equal")


#### AssertEqual

```go
func (t *GoT) AssertEqual(args ...interface{})
```
AssertEqual check for equality.

Expects:

    AssertEqual(a interface{}, b interface{}, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).AssertEqual(1, 1)
	Go(T).AssertEqual(1, 1, "should equal")


#### AssertLength

```go
func (t *GoT) AssertLength(args ...interface{})
```
AssertLength checks for length equal to `n int`. If the type passed cannot be
checked for length and error is logged stating as such.

Expects:

    AssertLength(a interface{}, n int, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).AssertLength("a", 1)
	Go(T).AssertLength("a", 1, "should be length")


#### AssertNil

```go
func (t *GoT) AssertNil(args ...interface{})
```
AssertNil checks for nil.

Expects:

    AssertNil(a interface{}, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).AssertNil(nil)
	Go(T).AssertNil(nil, "should be nil")


#### Refute

```go
func (t *GoT) Refute(args ...interface{})
```
Refute checks for false.

Expects:

    Refute(a bool, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).Refute(false)
	Go(T).Refute(false, "should not be true")


#### RefuteContains

```go
func (t *GoT) RefuteContains(args ...string)
```
RefuteContains checks for the non-existence of the second string inside the
first string.

Expects:

    RefuteContains(a string, b string, optional_message string)



#### RefuteDeepEqual

```go
func (t *GoT) RefuteDeepEqual(args ...interface{})
```
RefuteDeepEqual checks for no simlarity (see: reflect.DeepEqual).

Expects:

    RefuteDeepEqual(a interface{}, b interface{}, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	a1 := []string{"a"}
	a2 := []string{"b"}
	
	Go(T).RefuteDeepEqual(a1, a2)
	Go(T).RefuteDeepEqual(a1, a2, "should not deep equal")


#### RefuteEqual

```go
func (t *GoT) RefuteEqual(args ...interface{})
```
RefuteEqual checks for inequality.

Expects:

    RefuteEqual(a interface{}, b interface{}, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).RefuteEqual(1, 2)
	Go(T).RefuteEqual(1, 2, "should not equal")


#### RefuteLength

```go
func (t *GoT) RefuteLength(args ...interface{})
```
RefuteLength checks for length not equal to `n int`. If the type passed cannot
be checked for length this assertion will pass. (I'm not sure this is the best
way to handle this, feedback is welcome.)

Expects:

    RefuteLength(a interface{}, n int, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).RefuteLength("a", 0)
	Go(T).RefuteLength("a", 0, "should not be length")


#### RefuteNil

```go
func (t *GoT) RefuteNil(args ...interface{})
```
RefuteNil checks for not nil.

Expects:

    RefuteNil(a interface{}, optional_message string)



##### Example:
	// For example only, T comes from:
	//
	//     func TestFoo(T *testing.T)
	//
	T := new(testing.T)
	
	Go(T).RefuteNil(1)
	Go(T).RefuteNil(1, "should not be nil")



