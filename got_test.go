package got

import (
    "testing"
)

func TestAssert(T *testing.T) {
    Go(T).Assert(true, "should assert")
    Go(T).Assert(1 == 1, "should assert")
    Go(T).Assert("a" == "a", "should assert")
}

func TestRefute(T *testing.T) {
    Go(T).Refute(false, "should refute")
    Go(T).Refute(0 == 1, "should refute")
    Go(T).Refute("" == "a", "should refute")
}

func TestAssertEqual(T *testing.T) {
    Go(T).AssertEqual(true, true, "should assert equal")
    Go(T).AssertEqual(1, 1, "should assert equal")
    Go(T).AssertEqual("a", "a", "should assert equal")
}

func TestRefuteEqual(T *testing.T) {
    Go(T).RefuteEqual(false, true, "should refute equal")
    Go(T).RefuteEqual(0, 1, "should refute equal")
    Go(T).RefuteEqual("", "a", "should refute equal")
}

func TestAssertDeepEqual(T *testing.T) {
    Go(T).AssertDeepEqual(1, 1, "should assert deep equal int")
    Go(T).AssertDeepEqual(1.0, 1.0, "should assert deep equal float")

    m1 := map[string]string{"a": "a"}
    m2 := map[string]string{"a": "a"}
    Go(T).AssertDeepEqual(m1, m2, "should assert deep equal map")

    a1 := []string{"a", "b"}
    a2 := []string{"a", "b"}
    Go(T).AssertDeepEqual(a1, a2, "should assert deep equal array")
}

func TestRefuteDeepEqual(T *testing.T) {
    Go(T).RefuteDeepEqual("a", true, "should refute deep equal")
    Go(T).RefuteDeepEqual(1, 1.0, "should refute deep equal")
    Go(T).RefuteDeepEqual(1.0, nil, "should refute deep equal")

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

/* Intentionally Fail */
func TestAllFailureConditions(T *testing.T) {
    t := Go(T)

    t.Assert(false, "Assert: should fail here.")
    t.Refute(true, "Refute: should fail here.")

    t.AssertEqual(1, 2, "AssertEqual: should fail here.")
    t.RefuteEqual(1, 1, "RefuteEqual: should fail here.")

    t.AssertDeepEqual(1, 1.0, "AssertDeepEqual: should fail here.")
    t.RefuteDeepEqual(1, 1, "RefuteDeepEqual: should fail here.")

    t.AssertNil(1, "AssertNil: should fail here.")
    t.RefuteNil(nil, "RefuteNil: should fail here.")

    t.AssertLength("", 1, "AssertLength: should fail here.")
    t.RefuteLength("", 0, "RefuteLength: should fail here.")
}

func ExampleGo() {
    // file: foo_test.go
    // ------------------------------------------------ //
    //
    // Simple:
    //
    // import (
    //     "github.com/jmervine/GoT"
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
    //     "github.com/jmervine/GoT"
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
    //     . "github.com/jmervine/GoT"
    //     "testing"
    // )
    //
    // func TestFoo(T *testing.T) {
    //    Go(T).Assert(true, "should be true")
    //    Go(T).Refute(false, "should be false")
    // }
    //
    // ------------------------------------------------ //
}
