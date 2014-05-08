package GoT

import (
    "strings"
    "testing"
)

func TestAssert(T *testing.T) {
    Go(T).Assert(true, "should assert")
    Go(T).Assert(1 == 1)
    Go(T).Assert("a" == "a")
}

func TestRefute(T *testing.T) {
    Go(T).Refute(false, "should refute")
    Go(T).Refute(0 == 1)
    Go(T).Refute("" == "a")
}

func TestAssertEqual(T *testing.T) {
    Go(T).AssertEqual(true, true, "should assert equal")
    Go(T).AssertEqual(1, 1)
    Go(T).AssertEqual("a", "a")
}

func TestRefuteEqual(T *testing.T) {
    Go(T).RefuteEqual(false, true, "should refute equal")
    Go(T).RefuteEqual(0, 1)
    Go(T).RefuteEqual("", "a")
}

func TestAssertDeepEqual(T *testing.T) {
    Go(T).AssertDeepEqual(1, 1, "should assert deep equal int")
    Go(T).AssertDeepEqual(1.0, 1.0)

    m1 := map[string]string{"a": "a"}
    m2 := map[string]string{"a": "a"}
    Go(T).AssertDeepEqual(m1, m2, "should assert deep equal map")

    a1 := []string{"a", "b"}
    a2 := []string{"a", "b"}
    Go(T).AssertDeepEqual(a1, a2, "should assert deep equal array")
}

func TestRefuteDeepEqual(T *testing.T) {
    Go(T).RefuteDeepEqual("a", true, "should refute deep equal")
    Go(T).RefuteDeepEqual(1, 1.0)
    Go(T).RefuteDeepEqual(1.0, nil)

    m1 := map[string]string{"a": "a"}
    m2 := map[string]string{"a": "b"}
    Go(T).RefuteDeepEqual(m1, m2, "should refute deep equal map")

    a1 := []string{"a", "b"}
    a2 := []string{"a", "a"}
    Go(T).RefuteDeepEqual(a1, a2, "should refute deep equal array")
}

func TestAssertNil(T *testing.T) {
    Go(T).AssertNil(nil, "should assert nil on nil")

    var iface interface{}
    Go(T).AssertNil(iface, "should assert nil on empty interface")
}

func TestRefuteNil(T *testing.T) {
    Go(T).RefuteNil(1, "should refute nil on int")
    Go(T).RefuteNil("a", "should refute nil on string")

    arr := []string{"a", "b"}
    Go(T).RefuteNil(arr, "should refute nil on array")

    var iface interface{}
    iface = "a"
    Go(T).RefuteNil(iface, "should refute nil on interface")
}

func TestAssertLength(T *testing.T) {
    Go(T).AssertLength("asdf", 4, "should assert length on string")
    Go(T).AssertLength([]string{"a", "b"}, 2, "should assert length on array")
    Go(T).AssertLength(map[string]string{"a": "b"}, 1, "should assert length on array")
}

func TestRefuteLength(T *testing.T) {
    Go(T).RefuteLength("asd", 4, "should refute length on string")
    Go(T).RefuteLength([]string{"a"}, 2, "should refute length on array")
    Go(T).RefuteLength(map[string]string{}, 1, "should refute length on array")
}

func TestAssertContains(T *testing.T) {
    Go(T).AssertContains("asdf", "a")
    Go(T).AssertContains([1]int{1}, 1)
    Go(T).AssertContains([]int{1}, 1)
}

func TestRefuteContains(T *testing.T) {
    Go(T).RefuteContains("asdf", "q")
    Go(T).RefuteContains([1]int{1}, 2)
    Go(T).RefuteContains([]int{1}, 2)
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
}

func Test_equal(T *testing.T) {
    t := Go(T)

    pass, msg := equal("a", "a")

    // While using GoT to test this, I'm not using the assetion
    // that contins the helper being tested.
    t.Assert(pass, "equal should return true")
    t.Assert(msg == "", "message should be empty")

    pass, msg = equal("a", "b")
    t.Refute(pass, "equal should return false")
    t.Assert(msg == "", "message should be empty")

    pass, msg = equal([]string{"a"}, []string{"a"})
    t.Refute(pass, "equal should return false")
    t.Assert(strings.Contains(msg, "runtime error: comparing uncomparable type"), "should contain runtime error")
}

func Test_deepEqual(T *testing.T) {
    t := Go(T)

    // While using GoT to test this, I'm not using the assetion
    // that contins the helper being tested.
    a1 := []string{"a"}
    a2 := []string{"a"}
    a3 := []int{1}
    a4 := []int{1}

    m1 := map[string]string{"a": "b"}
    m2 := map[string]string{"a": "b"}
    m3 := map[string]int{"a": 1}
    m4 := map[string]int{"a": 1}

    t.Assert(deepEqual(a1, a2), "deep equal should return true")
    t.Assert(deepEqual(a3, a4), "deep equal should return true")
    t.Assert(deepEqual(m1, m2), "deep equal should return true")
    t.Assert(deepEqual(m3, m4), "deep equal should return true")

    a2 = []string{"b"}
    a4 = []int{2}
    m2 = map[string]string{"a": "c"}
    m4 = map[string]int{"a": 2}

    t.Refute(deepEqual(a1, a2), "deep equal should return false")
    t.Refute(deepEqual(a3, a4), "deep equal should return false")
    t.Refute(deepEqual(m1, m2), "deep equal should return false")
    t.Refute(deepEqual(m3, m4), "deep equal should return false")

    t.Refute(deepEqual(a1, a3), "deep equal should return false")
    t.Refute(deepEqual(a1, m1), "deep equal should return false")
}

func Test_isNil(T *testing.T) {
    t := Go(T)

    // While using GoT to test this, I'm not using the assetion
    // that contins the helper being tested.
    var (
        i   interface{}
        a   []string
        m   map[string]string
        s   string
        n   int
        f   float32
        b   bool
    )

    t.Assert(isNil(i))
    t.Assert(isNil(a))
    t.Assert(isNil(m))
    t.Assert(isNil(m))

    t.Refute(isNil(s))
    t.Refute(isNil(n))
    t.Refute(isNil(f))
    t.Refute(isNil(b))

    // structs
    type Type struct{}
    var tp *Type

    tt := new(Type)

    t.Refute(isNil(tt))
    t.Assert(isNil(tp))
}

func Test_checkLen(T *testing.T) {
    t := Go(T)

    // While using GoT to test this, I'm not using the assetion
    // that contins the helper being tested.
    a1 := []string{"a"}
    a2 := []int{1}

    m1 := map[string]string{"a": "b"}
    m2 := map[string]int{"a": 1}

    s := "a"

    t.Assert(checkLen(a1, 1))
    t.Assert(checkLen(a2, 1))
    t.Assert(checkLen(m1, 1))
    t.Assert(checkLen(m2, 1))
    t.Assert(checkLen(s, 1))

    t.Refute(checkLen(a1, 0))
    t.Refute(checkLen(a2, 0))
    t.Refute(checkLen(m1, 0))
    t.Refute(checkLen(m2, 0))
    t.Refute(checkLen(s, 0))
}

func Test_contains(T *testing.T) {
    t := Go(T)

    // While using GoT to test this, I'm not using the assetion
    // that contins the helper being tested.
    s1 := "asdf"
    s2 := "a"
    s3 := "q"

    t.Assert(contains(s1, s1))
    t.Assert(contains(s1, s2))

    t.Refute(contains(s1, s3))
    t.Refute(contains(s2, s1))

    a1 := [1]string{"a"}
    a2 := [1]int{1}
    a3 := []string{"a"}
    a4 := []int{1}

    //m1 := map[string]string{"a": "b"}
    //m2 := map[string]int{"a": 1}

    t.Assert(contains(a1, "a"))
    t.Assert(contains(a2, 1))
    t.Assert(contains(a3, "a"))
    t.Assert(contains(a4, 1))

    t.Refute(contains(a1, "q"))
    t.Refute(contains(a2, 2))
    t.Refute(contains(a3, "q"))
    t.Refute(contains(a4, 2))
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
}

func ExampleGoT_RefuteContains() {
    // T comes from:
    //
    //     func TestFoo(T *testing.T)

    Go(T).RefuteContains("asdf", "q")
    Go(T).RefuteContains([1]int{1}, 2)
    Go(T).RefuteContains([]int{1}, 2)
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
