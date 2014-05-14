/*

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


Note: In assertions, if no message is passed, then a nice message will be displayed
in an effor to be as clear as possible.

*/
package GoT

import (
    "fmt"
    "path"
    "reflect"
    "runtime"
    "strings"
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
//
// Expects:
//
//   Assert(a bool, optional_message string)
//
func (t *GoT) Assert(args ...interface{}) {
    msg := message(2, "Assert expected true, got false", args...)

    if !args[0].(bool) {
        t.error(msg)
    }
}

// Refute checks for false.
//
// Expects:
//
//   Refute(a bool, optional_message string)
//
func (t *GoT) Refute(args ...interface{}) {
    msg := message(2, "Refute expected false, got true", args...)

    if args[0].(bool) {
        t.error(msg)
    }
}

// AssertEqual check for equality.
//
// Expects:
//
//   AssertEqual(a, b interface{}, optional_message string)
//
func (t *GoT) AssertEqual(args ...interface{}) {
    msg := message(3, fmt.Sprintf("AssertEqual expected %v == %v", args...), args...)

    if !Equal(args[0], args[1]) {
        t.error(msg)
    }
}

// RefuteEqual checks for inequality.
//
// Expects:
//
//   RefuteEqual(a, b interface{}, optional_message string)
//
func (t *GoT) RefuteEqual(args ...interface{}) {
    msg := message(3, fmt.Sprintf("RefuteEqual expected %v != %v", args...), args...)

    if Equal(args[0], args[1]) {
        t.error(msg)
    }
}

// AssertDeepEqual checks for simlarity (see: reflect.DeepEqual).
// Note: I recommend custom messaging on this one, otherwise failure
// output will be ambigous.
//
// Expects:
//
//   AssertDeepEqual(a, b interface{}, optional_message string)
//
func (t *GoT) AssertDeepEqual(args ...interface{}) {
    msg := message(3, "AssertDeepEqual expected deep equal check to succeed", args...)

    if !DeepEqual(args[0], args[1]) {
        t.error(msg)
    }
}

// RefuteDeepEqual checks for no simlarity (see: reflect.DeepEqual).
// Note: I recommend custom messaging on this one, otherwise failure
// output will be ambigous.
//
// Expects:
//
//   RefuteDeepEqual(a, b interface{}, optional_message string)
//
func (t *GoT) RefuteDeepEqual(args ...interface{}) {
    msg := message(3, "RefuteDeepEqual expected deep equal check to fail", args...)

    if DeepEqual(args[0], args[1]) {
        t.error(msg)
    }
}

// AssertNil checks for nil.
//
// Expects:
//
//   AssertNil(a interface{}, optional_message string)
//
func (t *GoT) AssertNil(args ...interface{}) {
    msg := message(2, fmt.Sprintf("AssertNil expected %v to be nil", args[0]), args...)

    if !IsNil(args[0]) {
        t.error(msg)
    }
}

// RefuteNil checks for not nil.
//
// Expects:
//
//   RefuteNil(a interface{}, optional_message string)
//
func (t *GoT) RefuteNil(args ...interface{}) {
    msg := message(2, fmt.Sprintf("RefuteNil expected %v not to be nil", args[0]), args...)

    if IsNil(args[0]) {
        t.error(msg)
    }
}

// AssertLength checks for length equal to `n int`. If the type passed
// cannot be checked for length and error is logged stating as such.
//
// Expects:
//
//   AssertLength(a interface{}, n int, optional_message string)
//
func (t *GoT) AssertLength(args ...interface{}) {
    // message helper doesn't cut it for this one
    var msg string
    if len(args) == 3 && args[2].(string) != "" {
        msg = args[2].(string)
    }

    if pass, err := CheckLen(args[0], args[1].(int)); !pass && err != "" {
        t.error(err)
    } else if !pass {
        if msg == "" {
            l := reflect.ValueOf(args[0]).Len()
            msg = fmt.Sprintf("AssertLength expected length of %d, got %d", args[1].(int), l)
        }
        t.error(msg)
    }
}

// RefuteLength checks for length not equal to `n int`. If the type passed
// cannot be checked for length this assertion will pass. (I'm not sure
// this is the best way to handle this, feedback is welcome.)
//
// Expects:
//
//   RefuteLength(a interface{}, n int, optional_message string)
//
func (t *GoT) RefuteLength(args ...interface{}) {
    msg := message(3, fmt.Sprintf("RefuteLength expected length to not be %d", args[1]), args...)

    if pass, _ := CheckLen(args[0], args[1].(int)); pass {
        t.error(msg)
    }
}

// AssertContains checks for the existence of the second argument inside
// the first string, array or slice.
//
// Expects:
//
//   AssertContains(a, b interface{}, optional_message string)
//
func (t *GoT) AssertContains(args ...interface{}) {
    t.assertContains("AssertContains", args...)
}

// AssertHas is an alias for AssertContains
func (t *GoT) AssertHas(args ...interface{}) {
    t.assertContains("AssertHas", args...)
}

// RefuteContains checks for the non-existence of the second argument
// inside the first string, array or slice.
//
// Expects:
//
//   RefuteContains(a, b interface{}, optional_message string)
//
func (t *GoT) RefuteContains(args ...interface{}) {
    t.refuteContains("RefuteContains", args...)
}

// RefuteHas is an alias for RefuteContains
func (t *GoT) RefuteHas(args ...interface{}) {
    t.refuteContains("RefuteHas", args...)
}

// AssertHasKey checks for the existence of the second argument
// as a key for the first argument; supports maps only.
//
// Expects:
//
//   RefuteContains(a, b interface{}, optional_message string)
//
func (t *GoT) AssertHasKey(args ...interface{}) {
    msg := message(3, fmt.Sprintf("AssertHasKey expected key %v", args[1]), args...)

    if pass, err := HasKey(args[0], args[1]); !pass && err != "" {
        t.error(err)
    } else if !pass {
        t.error(msg)
    }
}

// RefuteHasKey checks for the existence of the second argument
// as a key for the first argument; supports maps only.
//
// Expects:
//
//   RefuteContains(a, b interface{}, optional_message string)
//
func (t *GoT) RefuteHasKey(args ...interface{}) {
    msg := message(3, fmt.Sprintf("RefuteHasKey did not expect key %v", args[1]), args...)

    if pass, _ := HasKey(args[0], args[1]); pass {
        t.error(msg)
    }
}

// TODO:
//
// For map's, array's, strings and maybe structs.
// - AssertHasVal
// - RefuteHasVal

/**
 * Exporting Helpers for manual use.
 ***************************************************/

// Equal is used for checking equality.
func Equal(a, b interface{}) bool {
    if DeepEqual(a, b) {
        return true
    }

    if reflect.ValueOf(a) == reflect.ValueOf(b) {
        return true
    }

    if fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b) {
        return true
    }

    return false
}

// DeepEqual calls reflect.DeepEqual, exporting for constancy only.
func DeepEqual(a, b interface{}) bool {
    return reflect.DeepEqual(a, b)
}

// IsNil checks for nil.
func IsNil(a interface{}) bool {
    if a == nil {
        return true
    }
    switch value := reflect.ValueOf(a); value.Kind() {
    case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
        return value.IsNil()
    }
    return false
}

// CheckLen checks length of the first argument based on the second. If the type
// being checked isn't a valid type for length check, then a message is returned
// stating such and the check fails.
func CheckLen(a interface{}, n int) (bool, string) {
    value := reflect.ValueOf(a)
    switch value.Kind() {
    case reflect.Map, reflect.Array, reflect.Slice, reflect.Chan, reflect.String:
        return value.Len() == n, ""
    }
    return false, "obtained value type has no length"
}

// Contains checks to see if the string, array or slice passed in the first
// argument contains the correctly typed value of the second argument. If the type
// passed is unsupported or some panic occurs while performing the check, a message
// will be returned as a string.
func Contains(a, b interface{}) (check bool, err string) {
    defer func() {
        if catch := recover(); catch != nil {
            check = false
            err = fmt.Sprint(catch)
        }
    }()

    haystack := reflect.ValueOf(a)
    needle := reflect.ValueOf(b)

    switch haystack.Kind() {
    case reflect.String:
        return strings.Contains(a.(string), b.(string)), ""
    case reflect.Array, reflect.Slice:
        for i := 0; i < haystack.Len(); i++ {
            // convert to two comparable types
            if haystack.Index(i).Interface() == needle.Interface() {
                return true, ""
            }
        }
    default:
        return false, fmt.Sprintf("assertion error: Contains doesn't support %v", haystack.Kind())
    }

    return false, ""
}

// HasKey checks to see if the map passed via the first argument has the
// correctly typed argument of the the second. Panics due to type mismatches
// will be passed back as an error string.
func HasKey(m, a interface{}) (check bool, err string) {
    defer func() {
        if catch := recover(); catch != nil {
            check = false
            err = fmt.Sprint(catch)
        }
    }()

    return reflect.ValueOf(m).MapIndex(reflect.ValueOf(a)).IsValid(), ""
}

/**
 * Helpers
 ***************************************************/

// message is a helper which tries to be a catch all for
// all assertion messaging, centralizing the login contained
// within to a single location
func message(n int, m string, args ...interface{}) string {
    if len(args) == n && args[n-1].(string) != "" {
        return args[n-1].(string)
    }
    return m
}

// error is a helper wrapping the "testing" packages internal
// error logger to include the actual file name and line number
// where the assetion failed
//
// Expects:
//
//   t.error(message string, optional_call_depth int)
//
func (t *GoT) error(args ...interface{}) {
    m := args[0].(string)

    depth := 2
    if len(args) == 2 {
        depth = args[1].(int)
    }

    var err string
    if _, file, line, ok := runtime.Caller(depth); ok {
        err = fmt.Sprintf("> %s:%d: %s", path.Base(file), line, m)
    } else {
        err = fmt.Sprintf("> ???:-1: %s", m)
    }
    t.t.Error(err)
}

// assertContains is the real assertion, allowing for aliasing.
func (t *GoT) assertContains(name string, args ...interface{}) {
    msg := message(3, fmt.Sprintf(name+" expected \"%v\" to contain \"%v\"", args...), args...)
    if pass, err := Contains(args[0], args[1]); !pass && err != "" {
        t.error(err, 3)
    } else if !pass {
        t.error(msg, 3)
    }
}

// refuteContains is the real assertion, allowing for aliasing.
func (t *GoT) refuteContains(name string, args ...interface{}) {
    msg := message(3, fmt.Sprintf(name+" expected \"%v\" to not contain \"%v\"", args...), args...)
    if pass, _ := Contains(args[0], args[1]); pass {
        t.error(msg, 3)
    }
}
