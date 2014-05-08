/*

Simple Assertion wrapper for Go's built in "testing" package.

GoT is designed to be as simple and unintrusive as possible while adding basic Assert and Refute methods to assist in writing clean and clean tests quickly.

Setup Examples:

    // ------------------------------------------------ //
    // Simple:

    import (
        "github.com/jmervine/GoT"
        "testing"
    )

    func TestFoo(T *testing.T) {
       t := got.Go(T)
       t.Assert(true, "should be true")
    }

    // ------------------------------------------------ //
    // Global:
    //

    import (
        "github.com/jmervine/GoT"
        "testing"
    )

    var Go := got.Go
    func TestFoo(T *testing.T) {
       Go(T).Assert(true, "should be true")
       Go(T).Refute(false, "should be false")
    }

    // ------------------------------------------------ //
    // Authors Perfered:

    import (
        . "github.com/jmervine/GoT"
        "testing"
    )

    func TestFoo(T *testing.T) {
       Go(T).Assert(true, "should be true")
       Go(T).Refute(false, "should be false")
    }

*/
package got

import (
    "fmt"
    "path"
    "reflect"
    "runtime"
    "testing"
)

/**
 * Setup
 ***************************************************/

// GoT is container for Go's build in testing package.
type GoT struct {
    t *testing.T
}

// Go wraps "testing" to apply assertions.
func Go(T *testing.T) *GoT {
    return &GoT{t: T}
}

/**
 * Assertions
 ***************************************************/

// Assert checks for true.
func (t *GoT) Assert(a bool, m string) {
    if !a {
        t.error(m)
    }
}

// Refute checks for false.
func (t *GoT) Refute(a bool, m string) {
    if a {
        t.error(m)
    }
}

// AssertEqual check for equality.
func (t *GoT) AssertEqual(a interface{}, b interface{}, m string) {
    if pass, err := equal(a, b); err != "" {
        t.error(err)
    } else if !pass {
        t.error(m)
    }
}

// RefuteEqual checks for inequality.
func (t *GoT) RefuteEqual(a interface{}, b interface{}, m string) {
    if pass, err := equal(a, b); err != "" {
        t.error(err)
    } else if pass {
        t.error(m)
    }
}

// AssertDeepEqual checks for simlarity (see: reflect.DeepEqual).
func (t *GoT) AssertDeepEqual(a interface{}, b interface{}, m string) {
    if !deepEqual(a, b) {
        t.error(m)
    }
}

// RefuteDeepEqual checks for no simlarity (see: reflect.DeepEqual).
func (t *GoT) RefuteDeepEqual(a interface{}, b interface{}, m string) {
    if deepEqual(a, b) {
        t.error(m)
    }
}

// AssertNil checks for nil.
func (t *GoT) AssertNil(a interface{}, m string) {
    if !isNil(a) {
        t.error(m)
    }
}

// RefuteNil checks for not nil.
func (t *GoT) RefuteNil(a interface{}, m string) {
    if isNil(a) {
        t.error(m)
    }
}

// AssertLength checks for length equal to `n int`. If the type passed
// cannot be checked for length and error is logged stating as such.
func (t *GoT) AssertLength(a interface{}, n int, m string) {
    if pass, err := checkLen(a, n); !pass && err != "" {
        t.error(err)
    } else if !pass {
        t.error(m)
    }
}

// RefuteLength checks for length not equal to `n int`. If the type passed
// cannot be checked for length this assertion will pass. (I'm not sure
// this is the best way to handle this, feedback is welcome.)
func (t *GoT) RefuteLength(a interface{}, n int, m string) {
    if pass, _ := checkLen(a, n); pass {
        t.error(m)
    }
}

/**
 * Helpers
 ***************************************************/

func (t *GoT) error(m string) {
    var err string
    if _, file, line, ok := runtime.Caller(2); ok {
        err = fmt.Sprintf("> %s:%d: %s", path.Base(file), line, m)
    } else {
        err = fmt.Sprintf("> ???:-1: %s", m)
    }
    t.t.Error(err)
}

func equal(a interface{}, b interface{}) (r bool, e string) {
    defer func() {
        if catch := recover(); catch != nil {
            r = false
            e = fmt.Sprint(catch)
        }
    }()

    return a == b, ""
}

func deepEqual(a interface{}, b interface{}) bool {
    return reflect.DeepEqual(a, b)
}

func isNil(a interface{}) bool {
    if a == nil {
        return true
    }
    switch value := reflect.ValueOf(a); value.Kind() {
    case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
        return value.IsNil()
    }
    return false
}

func checkLen(a interface{}, n int) (pass bool, err string) {
    value := reflect.ValueOf(a)
    switch value.Kind() {
    case reflect.Map, reflect.Array, reflect.Slice, reflect.Chan, reflect.String:
        return value.Len() == n, ""
    default:
        return false, "obtained value type has no length"
    }
}
