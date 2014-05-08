package functional

import (
    "testing"
    . ".."
)

func TestBool(T *testing.T) {
    Go(T).Assert(b1)
    Go(T).Assert(b1, "bool: assert")

    Go(T).Refute(b3)
    Go(T).Refute(b3, "bool: refute")

    Go(T).AssertEqual(b1, b2)
    Go(T).AssertEqual(b1, b2, "bool: assert equal")

    Go(T).RefuteEqual(b1, b3)
    Go(T).RefuteEqual(b1, b3, "bool: refute equal")

    Go(T).AssertDeepEqual(b1, b2)
    Go(T).AssertDeepEqual(b1, b2, "bool: assert deep equal")

    Go(T).RefuteDeepEqual(b1, b3)
    Go(T).RefuteDeepEqual(b1, b3, "bool: refute deep equal")

    Go(T).RefuteNil(b1)
    Go(T).RefuteNil(b1, "bool: refute nil")
}

func TestInt(T *testing.T) {
    Go(T).Assert(n1 == n2, "int: assert")
    Go(T).Refute(n1 == n3, "int: refute")

    Go(T).AssertEqual(n1, n2, "int: assert equal")
    Go(T).RefuteEqual(n1, n3, "int: assert equal")

    Go(T).AssertDeepEqual(n1, n2, "int: assert deep equal")
    Go(T).RefuteDeepEqual(n1, n3, "int: assert deep equal")

    Go(T).RefuteNil(n1, "int: refute nil")
}

func TestFloat(T *testing.T) {
    Go(T).Assert(f1 == f2, "float: assert")
    Go(T).Refute(f1 == f3, "float: refute")

    Go(T).AssertEqual(f1, f2, "float: assert equal")
    Go(T).RefuteEqual(f1, f3, "float: assert equal")

    Go(T).AssertDeepEqual(f1, f2, "float: assert deep equal")
    Go(T).RefuteDeepEqual(f1, f3, "float: assert deep equal")

    Go(T).RefuteNil(f1, "float: refute nil")
}

func TestString(T *testing.T) {
    Go(T).Assert(s1 == s2, "string: assert")
    Go(T).Refute(s1 == s3, "string: refute")

    Go(T).AssertEqual(s1, s2, "string: assert equal")
    Go(T).RefuteEqual(s1, s3, "string: assert equal")

    Go(T).AssertDeepEqual(s1, s2, "string: assert deep equal")
    Go(T).RefuteDeepEqual(s1, s3, "string: assert deep equal")

    Go(T).AssertLength(s1, 4)
    Go(T).AssertLength(s1, 4, "string: assert length")

    Go(T).RefuteLength(s1, 0)
    Go(T).RefuteLength(s1, 0, "string: assert length")

    Go(T).RefuteNil(s1, "string: refute nil")

    Go(T).AssertContains(s1, "as", "string: assert contains")
    Go(T).RefuteContains(s3, "as", "string: assert contains")

    // Alias for Contains
    Go(T).AssertHas(s1, "as", "string: assert contains")
    Go(T).RefuteHas(s3, "as", "string: assert contains")
}

func TestArray(T *testing.T) {
    Go(T).AssertDeepEqual(a1, a2, "array: assert deep equal")
    Go(T).RefuteDeepEqual(a1, a3, "array: refute deep equal")

    Go(T).AssertLength(a1, 4, "array: assert length")
    Go(T).RefuteLength(a1, 0, "array: refute length")

    Go(T).RefuteNil(a1, "array: refute nil")

    Go(T).AssertContains(a1, "a", "array: assert contains")
    Go(T).RefuteContains(a3, "a", "array: assert contains")
}

func TestSlice(T *testing.T) {
    Go(T).AssertDeepEqual(a4, a5, "slice: assert deep equal")
    Go(T).RefuteDeepEqual(a4, a6, "slice: refute deep equal")

    Go(T).AssertLength(a4, 4, "slice: assert length")
    Go(T).RefuteLength(a4, 0, "slice: refute length")

    Go(T).RefuteNil(a4, "slice: refute nil")

    Go(T).AssertContains(a4, "a", "slice: assert contains")
    Go(T).RefuteContains(a6, "a", "slice: assert contains")
}

func TestMap(T *testing.T) {
    Go(T).AssertDeepEqual(m1, m2, "map: assert deep equal")
    Go(T).RefuteDeepEqual(m1, m3, "map: refute deep equal")

    Go(T).AssertLength(m1, 2, "map: assert length")
    Go(T).RefuteLength(m1, 0, "map: refute length")

    Go(T).RefuteNil(m1, "map: refute nil")
}

func TestInterface(T *testing.T) {
    Go(T).AssertEqual(i1, i2, "interface: assert equal")
    Go(T).RefuteEqual(i1, i3, "interface: refute equal")

    Go(T).AssertDeepEqual(i1, i2, "interface: assert deep equal")
    Go(T).RefuteDeepEqual(i1, i3, "interface: refute deep equal")

    Go(T).AssertLength(i1, 4, "interface: assert length")
    Go(T).RefuteLength(i1, 0, "interface: refute length")
    Go(T).RefuteLength(i3, 4, "interface: refute length")
    Go(T).RefuteLength(i4, 4, "interface: refute length")

    Go(T).AssertNil(i4, "interface: assert nil")
    Go(T).RefuteNil(i1, "interface: refute nil")
}

func TestType(T *testing.T) {
    Go(T).AssertEqual(t1, t2, "type: assert equal")
    Go(T).RefuteEqual(t1, t3, "type: refute equal")

    Go(T).AssertDeepEqual(t1, t2, "type: assert deep equal")
    Go(T).RefuteDeepEqual(t1, t3, "type: refute deep equal")

    Go(T).RefuteLength(t1, 0, "type: assert length")
    Go(T).RefuteLength(t4, 4, "type: refute length")

    Go(T).RefuteNil(t1, "type: refute nil")
    Go(T).RefuteNil(t4, "type: refute nil")
}

func TestPointer(T *testing.T) {
    Go(T).AssertEqual(p1.Value, p2.Value, "pointer: assert equal on value")
    Go(T).RefuteEqual(p1.Value, p3.Value, "pointer: refute equal on value")

    Go(T).AssertDeepEqual(p1, p2, "pointer: assert deep equal")
    Go(T).AssertDeepEqual(p5, p6, "pointer: assert deep equal")

    Go(T).RefuteDeepEqual(p1, p3, "pointer: refute deep equal")
    Go(T).RefuteDeepEqual(p5, p7, "pointer: refute deep equal")

    Go(T).RefuteLength(p1, 0, "pointer: refute length")
    Go(T).RefuteLength(p4, 4, "pointer: refute length")
    Go(T).RefuteLength(p5, 4, "pointer: refute length")

    Go(T).AssertNil(p4, "pointer: assert nil")
    Go(T).AssertNil(p8, "pointer: assert nil")

    Go(T).RefuteNil(p1, "pointer: refute nil")
    Go(T).RefuteNil(p5, "pointer: refute nil")
}

