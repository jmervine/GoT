package GoT

import (
    "strings"
    "testing"
)

// Using Go "testing" package alone for testing, as it's a bad idea to
// test this using itself.

func TestAssert(t *testing.T) {
    T := new(testing.T) // Stub

    Go(T).Assert(true)
    if T.Failed() {
        t.Error("Assert should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Assert(true, "should assert")
    if T.Failed() {
        t.Error("Assert should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Assert(1 == 1)
    if T.Failed() {
        t.Error("Assert should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Assert("a" == "a")
    if T.Failed() {
        t.Error("Assert should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Assert(false)
    if !T.Failed() {
        t.Error("Assert should have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Assert(1 != 1)
    if !T.Failed() {
        t.Error("Assert should have failed.")
    }
}

func TestRefute(t *testing.T) {
    T := new(testing.T) // Stub

    Go(T).Refute(false, "should refute")
    if T.Failed() {
        t.Error("Refute should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Refute(0 == 1)
    if T.Failed() {
        t.Error("Refute should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Refute("" == "a")
    if T.Failed() {
        t.Error("Refute should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).Refute(true)
    if !T.Failed() {
        t.Error("Refute should have failed.")
    }
}

func TestAssertEqual(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).AssertEqual(true, true, "should assert equal")
    if T.Failed() {
        t.Error("AssertEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertEqual(1, 1)
    if T.Failed() {
        t.Error("AssertEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertEqual("a", "a")
    if T.Failed() {
        t.Error("AssertEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertEqual(1, 2)
    if !T.Failed() {
        t.Error("AssertEqual should have failed.")
    }

}

func TestRefuteEqual(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).RefuteEqual(false, true, "should refute equal")
    if T.Failed() {
        t.Error("RefuteEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteEqual(0, 1)
    if T.Failed() {
        t.Error("RefuteEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteEqual(0, 1)
    if T.Failed() {
        t.Error("RefuteEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteEqual(1, 1)
    if !T.Failed() {
        t.Error("RefuteEqual should have failed.")
    }

}

func TestAssertDeepEqual(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).AssertDeepEqual(1, 1, "should assert deep equal int")
    if T.Failed() {
        t.Error("AssertDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertDeepEqual(1.0, 1.0)
    if T.Failed() {
        t.Error("AssertDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    m1 := map[string]string{"a": "a"}
    m2 := map[string]string{"a": "a"}
    Go(T).AssertDeepEqual(m1, m2, "should assert deep equal map")
    if T.Failed() {
        t.Error("AssertDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    a1 := []string{"a", "b"}
    a2 := []string{"a", "b"}
    Go(T).AssertDeepEqual(a1, a2, "should assert deep equal array")
    if T.Failed() {
        t.Error("AssertDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertDeepEqual(1, 2, "should assert deep equal int")
    if !T.Failed() {
        t.Error("AssertDeepEqual should have failed.")
    }
}

func TestRefuteDeepEqual(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).RefuteDeepEqual("a", true, "should refute deep equal")
    if T.Failed() {
        t.Error("RefuteDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteDeepEqual(1, 1.0)
    if T.Failed() {
        t.Error("RefuteDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteDeepEqual(1.0, nil)
    if T.Failed() {
        t.Error("RefuteDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    m1 := map[string]string{"a": "a"}
    m2 := map[string]string{"a": "b"}
    Go(T).RefuteDeepEqual(m1, m2, "should refute deep equal map")
    if T.Failed() {
        t.Error("RefuteDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    a1 := []string{"a", "b"}
    a2 := []string{"a", "a"}
    Go(T).RefuteDeepEqual(a1, a2, "should refute deep equal array")
    if T.Failed() {
        t.Error("RefuteDeepEqual should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteDeepEqual(1, 1, "should refute deep equal")
    if !T.Failed() {
        t.Error("RefuteDeepEqual should have failed.")
    }
}

func TestAssertNil(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).AssertNil(nil, "should assert nil on nil")
    if T.Failed() {
        t.Error("AssertNil should not have failed.")
    }

    T = new(testing.T) // Stub
    var iface interface{}
    Go(T).AssertNil(iface, "should assert nil on empty interface")
    if T.Failed() {
        t.Error("AssertNil should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertNil("a", "should assert nil on nil")
    if !T.Failed() {
        t.Error("AssertNil should have failed.")
    }
}

func TestRefuteNil(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).RefuteNil(1, "should refute nil on int")
    if T.Failed() {
        t.Error("RefuteNil should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteNil("a", "should refute nil on string")
    if T.Failed() {
        t.Error("RefuteNil should not have failed.")
    }

    T = new(testing.T) // Stub
    arr := []string{"a", "b"}
    Go(T).RefuteNil(arr, "should refute nil on array")
    if T.Failed() {
        t.Error("RefuteNil should not have failed.")
    }

    T = new(testing.T) // Stub
    var iface interface{}
    iface = "a"
    Go(T).RefuteNil(iface, "should refute nil on interface")
    if T.Failed() {
        t.Error("RefuteNil should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteNil(nil, "should refute nil on int")
    if !T.Failed() {
        t.Error("RefuteNil should have failed.")
    }
}

func TestAssertLength(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).AssertLength("asdf", 4, "should assert length on string")
    if T.Failed() {
        t.Error("AssertLength should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertLength([]string{"a", "b"}, 2, "should assert length on array")
    if T.Failed() {
        t.Error("AssertLength should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertLength(map[string]string{"a": "b"}, 1, "should assert length on array")
    if T.Failed() {
        t.Error("AssertLength should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertLength("asdf", 0, "should assert length on string")
    if !T.Failed() {
        t.Error("AssertLength should not have failed.")
    }
}

func TestRefuteLength(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).RefuteLength("asd", 4, "should refute length on string")
    if T.Failed() {
        t.Error("RefuteLength should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteLength([]string{"a"}, 2, "should refute length on array")
    if T.Failed() {
        t.Error("RefuteLength should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteLength(map[string]string{}, 1, "should refute length on array")
    if T.Failed() {
        t.Error("RefuteLength should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteLength("", 0, "should refute length on string")
    if !T.Failed() {
        t.Error("RefuteLength should have failed.")
    }
}

func TestAssertContains(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).AssertContains("asdf", "a")
    if T.Failed() {
        t.Error("AssertContains should not have failed.")
    }

    // testing AssertHas which is an alias of AssertContains here
    T = new(testing.T) // Stub
    Go(T).AssertHas("asdf", "a")
    if T.Failed() {
        t.Error("AssertHas should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertContains([1]int{1}, 1)
    if T.Failed() {
        t.Error("AssertContains should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertContains([]int{1}, 1)
    if T.Failed() {
        t.Error("AssertContains should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).AssertContains("asdf", "q")
    if !T.Failed() {
        t.Error("AssertContains should have failed.")
    }

    // testing AssertHas here
    T = new(testing.T) // Stub
    Go(T).AssertHas("asdf", "q")
    if !T.Failed() {
        t.Error("AssertHas should have failed.")
    }
}

func TestRefuteContains(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).RefuteContains("asdf", "q")
    if T.Failed() {
        t.Error("RefuteContains should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteHas("asdf", "q")
    if T.Failed() {
        t.Error("RefuteHas should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteContains([1]int{1}, 2)
    if T.Failed() {
        t.Error("RefuteContains should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteContains([]int{1}, 2)
    if T.Failed() {
        t.Error("RefuteContains should not have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteContains("asdf", "a")
    if !T.Failed() {
        t.Error("RefuteContains should have failed.")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteHas("asdf", "a")
    if !T.Failed() {
        t.Error("RefuteHas should have failed.")
    }
}

func TestAssertHasKey(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).AssertHasKey(map[string]string{"a": "a"}, "a")
    if T.Failed() {
        t.Error("AssertHasKey should have passed")
    }

    T = new(testing.T) // Stub
    Go(T).AssertHasKey(map[int]string{1: "a"}, 1)
    if T.Failed() {
        t.Error("AssertHasKey should have passed")
    }

    T = new(testing.T) // Stub
    Go(T).AssertHasKey(map[float64]string{1.1: "a"}, 1.1)
    if T.Failed() {
        t.Error("AssertHasKey should have passed")
    }

    T = new(testing.T) // Stub
    Go(T).AssertHasKey(map[string]string{"a": "a"}, "b")
    if !T.Failed() {
        t.Error("AssertHasKey should have failed")
    }

    T = new(testing.T) // Stub
    Go(T).AssertHasKey(map[int]string{1: "a"}, 2)
    if !T.Failed() {
        t.Error("AssertHasKey should have failed")
    }

    T = new(testing.T) // Stub
    Go(T).AssertHasKey(map[float64]string{1.1: "a"}, 1.2)
    if !T.Failed() {
        t.Error("AssertHasKey should have failed")
    }
}

func TestRefuteHasKey(t *testing.T) {
    T := new(testing.T) // Stub
    Go(T).RefuteHasKey(map[string]string{"a": "a"}, "a")
    if !T.Failed() {
        t.Error("RefuteHasKey should have failed")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteHasKey(map[int]string{1: "a"}, 1)
    if !T.Failed() {
        t.Error("RefuteHasKey should have failed")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteHasKey(map[float64]string{1.1: "a"}, 1.1)
    if !T.Failed() {
        t.Error("RefuteHasKey should have failed")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteHasKey(map[string]string{"a": "a"}, "b")
    if T.Failed() {
        t.Error("RefuteHasKey should have passed")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteHasKey(map[int]string{1: "a"}, 2)
    if T.Failed() {
        t.Error("RefuteHasKey should have passed")
    }

    T = new(testing.T) // Stub
    Go(T).RefuteHasKey(map[float64]string{1.1: "a"}, 1.2)
    if T.Failed() {
        t.Error("RefuteHasKey should have passed")
    }
}

func TestAllFailureConditions(T *testing.T) {
    // Uncomment when testing changes to failure messaging.
    //

    //Go(T).Assert(false)
    //Go(T).Assert(false, "Assert: should fail here.")

    //Go(T).Refute(true)
    //Go(T).Refute(true, "Refute: should fail here.")

    //Go(T).AssertEqual(1, 2)
    //Go(T).AssertEqual(1, 2, "AssertEqual: should fail here.")

    //Go(T).RefuteEqual(1, 1)
    //Go(T).RefuteEqual(1, 1, "RefuteEqual: should fail here.")

    //Go(T).AssertDeepEqual(1, 1.0)
    //Go(T).AssertDeepEqual(1, 1.0, "AssertDeepEqual: should fail here.")

    //Go(T).RefuteDeepEqual(1, 1)
    //Go(T).RefuteDeepEqual(1, 1, "RefuteDeepEqual: should fail here.")

    //Go(T).AssertNil(1)
    //Go(T).AssertNil(1, "AssertNil: should fail here.")

    //Go(T).RefuteNil(nil)
    //Go(T).RefuteNil(nil, "RefuteNil: should fail here.")

    //Go(T).AssertLength("", 1)
    //Go(T).AssertLength("", 1, "AssertLength: should fail here.")

    //Go(T).RefuteLength("", 0)
    //Go(T).RefuteLength("", 0, "RefuteLength: should fail here.")

    //Go(T).AssertContains("asdf", "q", "AssertContains: custom message")
    //Go(T).AssertHas("asdf", "q", "AssertHas: custom message")

    //Go(T).AssertContains("asdf", "q")
    //Go(T).AssertContains([1]int{1}, 2)

    //Go(T).RefuteContains("asdf", "a", "AssertContains: custom message")

    //Go(T).RefuteContains("asdf", "a")
    //Go(T).RefuteHas("asdf", "a")
    //Go(T).RefuteContains([1]int{1}, 1)
}

func Test_message(t *testing.T) {
    args := []interface{}{"a", "a"}
    msg := message(3, "default message", args...)

    if msg != "default message" {
        t.Error("expected default message")
    }

    args = append(args, "custom message")
    msg = message(3, "default message", args...)

    if msg != "custom message" {
        t.Error("expected custom message")
    }
}

func Test_equal(t *testing.T) {
    pass, msg := equal("a", "a")

    if !pass {
        t.Error("expected equal to pass")
    }

    if msg != "" {
        t.Error("expected msg to be empty")
    }

    pass, msg = equal("a", "b")
    if pass {
        t.Error("expected equal to fail")
    }
    if msg != "" {
        t.Error("expected msg to be empty")
    }

    pass, msg = equal([]string{"a"}, []string{"a"})
    if pass {
        t.Error("expected equal to fail")
    }
    if !strings.Contains(msg, "runtime error: comparing uncomparable type") {
        t.Error("expected msg to have run time error")
    }
}

func Test_deepEqual(t *testing.T) {

    a1 := []string{"a"}
    a2 := []string{"a"}
    a3 := []int{1}
    a4 := []int{1}

    m1 := map[string]string{"a": "b"}
    m2 := map[string]string{"a": "b"}
    m3 := map[string]int{"a": 1}
    m4 := map[string]int{"a": 1}

    if !deepEqual(a1, a2) {
        t.Error("deepEqual should pass")
    }
    if !deepEqual(a3, a4) {
        t.Error("deepEqual should pass")
    }
    if !deepEqual(m1, m2) {
        t.Error("deepEqual should pass")
    }
    if !deepEqual(m3, m4) {
        t.Error("deepEqual should pass")
    }

    a2 = []string{"b"}
    a4 = []int{2}
    m2 = map[string]string{"a": "c"}
    m4 = map[string]int{"a": 2}

    if deepEqual(a1, a2) {
        t.Error("deepEqual should fail")
    }
    if deepEqual(a3, a4) {
        t.Error("deepEqual should fail")
    }
    if deepEqual(m1, m2) {
        t.Error("deepEqual should fail")
    }
    if deepEqual(m3, m4) {
        t.Error("deepEqual should fail")
    }

    if deepEqual(a1, a3) {
        t.Error("deepEqual should fail")
    }
    if deepEqual(a1, m1) {
        t.Error("deepEqual should fail")
    }
}

func Test_isNil(t *testing.T) {

    var (
        i   interface{}
        a   []string
        m   map[string]string
        s   string
        n   int
        f   float32
        b   bool
    )

    if !isNil(i) {
        t.Error("expected isNil to pass")
    }
    if !isNil(a) {
        t.Error("expected isNil to pass")
    }
    if !isNil(m) {
        t.Error("expected isNil to pass")
    }

    if isNil(s) {
        t.Error("expected isNil to fail")
    }
    if isNil(n) {
        t.Error("expected isNil to fail")
    }
    if isNil(f) {
        t.Error("expected isNil to fail")
    }
    if isNil(b) {
        t.Error("expected isNil to fail")
    }

    // structs
    type Type struct{}
    var tp *Type

    tt := new(Type)

    if isNil(tt) {
        t.Error("expected isNil to fail")
    }
    if !isNil(tp) {
        t.Error("expected isNil to fail")
    }
}

func Test_checkLen(t *testing.T) {

    a1 := []string{"a"}
    a2 := []int{1}

    m1 := map[string]string{"a": "b"}
    m2 := map[string]int{"a": 1}

    s := "a"

    if pass, msg := checkLen(a1, 1); !pass || msg != "" {
        t.Error("expected checkLen to pass without messaging")
    }
    if pass, msg := checkLen(a2, 1); !pass || msg != "" {
        t.Error("expected checkLen to pass without messaging")
    }
    if pass, msg := checkLen(m1, 1); !pass || msg != "" {
        t.Error("expected checkLen to pass without messaging")
    }
    if pass, msg := checkLen(m2, 1); !pass || msg != "" {
        t.Error("expected checkLen to pass without messaging")
    }
    if pass, msg := checkLen(s, 1); !pass || msg != "" {
        t.Error("expected checkLen to pass without messaging")
    }

    if pass, _ := checkLen(a1, 0); pass {
        t.Error("expected checkLen to fail")
    }
    if pass, _ := checkLen(a2, 0); pass {
        t.Error("expected checkLen to fail")
    }
    if pass, _ := checkLen(m1, 0); pass {
        t.Error("expected checkLen to fail")
    }
    if pass, _ := checkLen(m2, 0); pass {
        t.Error("expected checkLen to fail")
    }
    if pass, _ := checkLen(s, 0); pass {
        t.Error("expected checkLen to fail")
    }
}

func Test_hasKey(t *testing.T) {
    m1 := map[string]string{"a": "b"}
    m2 := map[int]string{1: "b"}

    // m1
    if pass, _ := hasKey(m1, "a"); !pass {
        t.Error("expected hasKey to pass")
    }

    if pass, _ := hasKey(m1, "b"); pass {
        t.Error("expected hasKey to fail")
    }

    // m2
    if pass, _ := hasKey(m2, 1); !pass {
        t.Error("expected hasKey to pass")
    }

    if pass, _ := hasKey(m2, 2); pass {
        t.Error("expected hasKey to fail")
    }
    if pass, _ := hasKey(m2, "1"); pass {
        t.Error("expected hasKey to fail")
    }
}

func Test_contains(t *testing.T) {
    s1 := "asdf"
    s2 := "a"
    s3 := "q"

    if pass, _ := contains(s1, s1); !pass {
        t.Error("expected contains to pass")
    }
    if pass, _ := contains(s1, s2); !pass {
        t.Error("expected contains to pass")
    }

    if pass, _ := contains(s1, s3); pass {
        t.Error("expected contains to fail")
    }
    if pass, _ := contains(s2, s1); pass {
        t.Error("expected contains to fail")
    }

    a1 := [1]string{"a"}
    a2 := [1]int{1}
    a3 := []string{"a"}
    a4 := []int{1}

    if pass, _ := contains(a1, "a"); !pass {
        t.Error("expected contains to pass")
    }
    if pass, _ := contains(a2, 1); !pass {
        t.Error("expected contains to pass")
    }
    if pass, _ := contains(a3, "a"); !pass {
        t.Error("expected contains to pass")
    }
    if pass, _ := contains(a4, 1); !pass {
        t.Error("expected contains to pass")
    }

    if pass, _ := contains(a1, "q"); pass {
        t.Error("expected contains to pass")
    }
    if pass, _ := contains(a2, 2); pass {
        t.Error("expected contains to fail")
    }
    if pass, _ := contains(a3, "q"); pass {
        t.Error("expected contains to fail")
    }
    if pass, _ := contains(a4, 2); pass {
        t.Error("expected contains to fail")
    }
}

// Examples
var T = new(testing.T)

func ExampleGoT_Assert() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).Assert(true)
    Go(T).Assert(true, "should be true")
}

func ExampleGoT_Refute() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).Refute(false)
    Go(T).Refute(false, "should not be true")
}

func ExampleGoT_AssertEqual() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).AssertEqual(1, 1)
    Go(T).AssertEqual(1, 1, "should equal")
}

func ExampleGoT_RefuteEqual() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).RefuteEqual(1, 2)
    Go(T).RefuteEqual(1, 2, "should not equal")
}

func ExampleGoT_AssertDeepEqual() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    a1 := []string{"a"}
    a2 := []string{"a"}

    Go(T).AssertDeepEqual(a1, a2)
    Go(T).AssertDeepEqual(a1, a2, "should deep equal")
}

func ExampleGoT_RefuteDeepEqual() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    a1 := []string{"a"}
    a2 := []string{"b"}

    Go(T).RefuteDeepEqual(a1, a2)
    Go(T).RefuteDeepEqual(a1, a2, "should not deep equal")
}

func ExampleGoT_AssertNil() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).AssertNil(nil)
    Go(T).AssertNil(nil, "should be nil")
}

func ExampleGoT_RefuteNil() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).RefuteNil(1)
    Go(T).RefuteNil(1, "should not be nil")
}

func ExampleGoT_AssertLength() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).AssertLength("a", 1)
    Go(T).AssertLength("a", 1, "should be length")
}

func ExampleGoT_RefuteLength() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).RefuteLength("a", 0)
    Go(T).RefuteLength("a", 0, "should not be length")
}

func ExampleGoT_AssertContains() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).AssertContains("asdf", "a")
    Go(T).AssertContains([1]int{1}, 1)
    Go(T).AssertContains([]int{1}, 1)

    // Alias:
    Go(T).AssertHas("asdf", "a")
}

func ExampleGoT_RefuteContains() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).RefuteContains("asdf", "q")
    Go(T).RefuteContains([1]int{1}, 2)
    Go(T).RefuteContains([]int{1}, 2)

    // Alias:
    Go(T).RefuteHas("asdf", "q")
}

func ExampleGo() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).Assert(true)
    Go(T).Assert(true, "should be true")

    Go(T).Refute(false)
    Go(T).Refute(false, "should not be true")
}
