
# GoTest

[![GoDoc](https://godoc.org/gopkg.in/jmervine/GoT.v1?status.png)](https://godoc.org/gopkg.in/jmervine/GoT.v1) [![Build Status](https://travis-ci.org/jmervine/GoT.svg?branch=master)](https://travis-ci.org/jmervine/GoT)

<center>![GoT](http://cdn.mervine.net/GoT.jpg)</center>

#### Simple Assertion wrapper for Go's built in "testing" package.

GoT is designed to be as simple and unintrusive as possible while adding basic Assert and Refute methods to assist in writing clean and clean tests quickly.

##### See [examples](_example) for detailed assertion usage.

## [Documentation](https://godoc.org/gopkg.in/jmervine/GoT.v1)

```go
import "gopkg.in/jmervine/GoT.v1"
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

    // or just the helpers //
    func TestBar(t *testing.T) {
        if GoT.Equal("a", "a") {
            t.Error("expected equality")
        }
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

    var Go = GoT.Go

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
> Go wraps "testing" to apply assertions.

```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).Assert(true)
	Go(T).Assert(true, "should be true")
	
	Go(T).Refute(false)
	Go(T).Refute(false, "should not be true")

```
#### Assert
```go
func (t *GoT) Assert(args ...interface{})
```
> Assert checks for true.

> Expects:

Assert(a bool, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).Assert(true)
	Go(T).Assert(true, "should be true")

```

#### AssertContains
```go
func (t *GoT) AssertContains(args ...interface{})
```
> AssertContains checks for the existence of the second argument inside the
> first string, array or slice.

> Expects:

AssertContains(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).AssertContains("asdf", "a")
	Go(T).AssertContains([1]int{1}, 1)
	Go(T).AssertContains([]int{1}, 1)
	
	// Alias:
	Go(T).AssertHas("asdf", "a")

```

#### AssertDeepEqual
```go
func (t *GoT) AssertDeepEqual(args ...interface{})
```
> AssertDeepEqual checks for simlarity (see: reflect.DeepEqual). Note: I
> recommend custom messaging on this one, otherwise failure output will be
> ambigous.

> Expects:

AssertDeepEqual(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	a1 := []string{"a"}
	a2 := []string{"a"}
	
	Go(T).AssertDeepEqual(a1, a2)
	Go(T).AssertDeepEqual(a1, a2, "should deep equal")

```

#### AssertEqual
```go
func (t *GoT) AssertEqual(args ...interface{})
```
> AssertEqual check for equality.

> Expects:

AssertEqual(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).AssertEqual(1, 1)
	Go(T).AssertEqual(1, 1, "should equal")

```

#### AssertHas
```go
func (t *GoT) AssertHas(args ...interface{})
```
> AssertHas is an alias for AssertContains



#### AssertHasKey
```go
func (t *GoT) AssertHasKey(args ...interface{})
```
> AssertHasKey checks for the existence of the second argument as a key for
> the first argument; supports maps only.

> Expects:

RefuteContains(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).AssertHasKey(map[string]string{"a": "a"}, "a")
	Go(T).AssertHasKey(map[int]string{1: "a"}, 1)

```

#### AssertLength
```go
func (t *GoT) AssertLength(args ...interface{})
```
> AssertLength checks for length equal to `n int`. If the type passed cannot
> be checked for length and error is logged stating as such.

> Expects:

AssertLength(a interface{}, n int, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).AssertLength("a", 1)
	Go(T).AssertLength("a", 1, "should be length")

```

#### AssertNil
```go
func (t *GoT) AssertNil(args ...interface{})
```
> AssertNil checks for nil.

> Expects:

AssertNil(a interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).AssertNil(nil)
	Go(T).AssertNil(nil, "should be nil")

```

#### Refute
```go
func (t *GoT) Refute(args ...interface{})
```
> Refute checks for false.

> Expects:

Refute(a bool, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).Refute(false)
	Go(T).Refute(false, "should not be true")

```

#### RefuteContains
```go
func (t *GoT) RefuteContains(args ...interface{})
```
> RefuteContains checks for the non-existence of the second argument inside
> the first string, array or slice.

> Expects:

RefuteContains(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).RefuteContains("asdf", "q")
	Go(T).RefuteContains([1]int{1}, 2)
	Go(T).RefuteContains([]int{1}, 2)
	
	// Alias:
	Go(T).RefuteHas("asdf", "q")

```

#### RefuteDeepEqual
```go
func (t *GoT) RefuteDeepEqual(args ...interface{})
```
> RefuteDeepEqual checks for no simlarity (see: reflect.DeepEqual). Note: I
> recommend custom messaging on this one, otherwise failure output will be
> ambigous.

> Expects:

RefuteDeepEqual(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	a1 := []string{"a"}
	a2 := []string{"b"}
	
	Go(T).RefuteDeepEqual(a1, a2)
	Go(T).RefuteDeepEqual(a1, a2, "should not deep equal")

```

#### RefuteEqual
```go
func (t *GoT) RefuteEqual(args ...interface{})
```
> RefuteEqual checks for inequality.

> Expects:

RefuteEqual(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).RefuteEqual(1, 2)
	Go(T).RefuteEqual(1, 2, "should not equal")

```

#### RefuteHas
```go
func (t *GoT) RefuteHas(args ...interface{})
```
> RefuteHas is an alias for RefuteContains



#### RefuteHasKey
```go
func (t *GoT) RefuteHasKey(args ...interface{})
```
> RefuteHasKey checks for the existence of the second argument as a key for
> the first argument; supports maps only.

> Expects:

RefuteContains(a, b interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).RefuteHasKey(map[string]string{"a": "a"}, "b")
	Go(T).RefuteHasKey(map[int]string{1: "a"}, 2)

```

#### RefuteLength
```go
func (t *GoT) RefuteLength(args ...interface{})
```
> RefuteLength checks for length not equal to `n int`. If the type passed
> cannot be checked for length this assertion will pass. (I'm not sure this is
> the best way to handle this, feedback is welcome.)

> Expects:

RefuteLength(a interface{}, n int, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).RefuteLength("a", 0)
	Go(T).RefuteLength("a", 0, "should not be length")

```

#### RefuteNil
```go
func (t *GoT) RefuteNil(args ...interface{})
```
> RefuteNil checks for not nil.

> Expects:

RefuteNil(a interface{}, optional_message string)


```go
    // Example:
	// T comes from:
	//
	//     func TestFoo(T *testing.T)
	
	Go(T).RefuteNil(1)
	Go(T).RefuteNil(1, "should not be nil")

```


#### CheckLen
```go
func CheckLen(a interface{}, n int) (bool, string)
```
> CheckLen checks length of the first argument based on the second. If the
> type being checked isn't a valid type for length check, then a message is
> returned stating such and the check fails.
```go
    // Example:
	// Example in testing situation.
	//
	// t comes from:
	//
	//     func TestFoo(t *testing.T)
	var a [1]int
	if pass, err := CheckLen(a, 1); err != "" {
	    t.Error(err)
	} else if !pass {
	    t.Error("should have len of one")
	}
	
	_, err := CheckLen(1, 1)
	fmt.Println(err)
	
	// Output:
	// obtained value type has no length

```
#### Contains
```go
func Contains(a, b interface{}) (check bool, err string)
```
> Contains checks to see if the string, array or slice passed in the first
> argument contains the correctly typed value of the second argument. If the
> type passed is unsupported or some panic occurs while performing the check,
> a message will be returned as a string.
```go
    // Example:
	// Example in testing situation.
	//
	// t comes from:
	//
	//     func TestFoo(t *testing.T)
	var a [1]int
	if pass, err := Contains(a, 1); err != "" {
	    t.Error(err)
	} else if pass {
	    t.Error("should not have one")
	}

```
#### DeepEqual
```go
func DeepEqual(a, b interface{}) bool
```
> DeepEqual calls reflect.DeepEqual, exporting for constancy only.
```go
    // Example:
	// Example in testing situation.
	//
	// t comes from:
	//
	//     func TestFoo(t *testing.T)
	if pass := DeepEqual("a", "a"); !pass {
	    t.Error("a should deep equal a")
	}

```
#### Equal
```go
func Equal(a, b interface{}) bool
```
> Equal is used for checking equality.
```go
    // Example:
	// Example in testing situation.
	//
	// t comes from:
	//
	//     func TestFoo(t *testing.T)
	if !Equal("a", "a") {
	    t.Error("a should equal a")
	}

```
#### HasKey
```go
func HasKey(m, a interface{}) (check bool, err string)
```
> HasKey checks to see if the map passed via the first argument has the
> correctly typed argument of the the second. Panics due to type mismatches
> will be passed back as an error string.
```go
    // Example:
	// Example in testing situation.
	//
	// t comes from:
	//
	//     func TestFoo(t *testing.T)
	m := map[string]string{"a": "a"}
	if pass, err := HasKey(m, "a"); err != "" {
	    t.Error(err)
	} else if !pass {
	    t.Error("map should have key a")
	}

```
#### IsNil
```go
func IsNil(a interface{}) bool
```
> IsNil checks for nil.
```go
    // Example:
	// Example in testing situation.
	//
	// t comes from:
	//
	//     func TestFoo(t *testing.T)
	if pass := IsNil("a"); pass {
	    t.Error("a should not be nil")
	}

```

