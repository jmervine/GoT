package got

// Simple Assertion wrapper for Go's built in "testing" package.

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

type GoT struct {
    t *testing.T
}

// Core "testing" wrapper to apply assertions.
func Go(T *testing.T) *GoT {
    return &GoT{t: T}
}

/**
 * Assertions
 ***************************************************/

// Check for true.
func (t *GoT) Assert(a bool, m string) {
    if !a {
        t.error(m)
    }
}

// Check for false.
func (t *GoT) Refute(a bool, m string) {
    if a {
        t.error(m)
    }
}

// Check for equality.
func (t *GoT) AssertEqual(a interface{}, b interface{}, m string) {
    if pass, err := equal(a, b); err != "" {
        t.error(err)
    } else if !pass {
        t.error(m)
    }
}

// Check for inequality.
func (t *GoT) RefuteEqual(a interface{}, b interface{}, m string) {
    if pass, err := equal(a, b); err != "" {
        t.error(err)
    } else if pass {
        t.error(m)
    }
}

// Check for simlarity (see: reflect.DeepEqual).
func (t *GoT) AssertDeepEqual(a interface{}, b interface{}, m string) {
    if !deepEqual(a, b) {
        t.error(m)
    }
}

// Check for no simlarity (see: reflect.DeepEqual).
func (t *GoT) RefuteDeepEqual(a interface{}, b interface{}, m string) {
    if deepEqual(a, b) {
        t.error(m)
    }
}

// Check for nil.
func (t *GoT) AssertNil(a interface{}, m string) {
    if !isNil(a) {
        t.error(m)
    }
}

// Check for not nil.
func (t *GoT) RefuteNil(a interface{}, m string) {
    if isNil(a) {
        t.error(m)
    }
}

// Check length is n.
func (t *GoT) AssertLength(a interface{}, n int, m string) {
    if pass, err := checkLen(a, n); !pass && err != "" {
        t.error(err)
    } else if !pass {
        t.error(m)
    }
}

// Check length is not n.
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
