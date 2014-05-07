package functional

import (
    "testing"
    t ".."
)

func TestBool(T *testing.T) {
    t.Go(T).Assert(b1, "bool: assert")
    t.Go(T).Refute(b3, "bool: refute")

    t.Go(T).AssertEqual(b1, b2, "bool: assert equal")
    t.Go(T).RefuteEqual(b1, b3, "bool: refute equal")

    t.Go(T).AssertDeepEqual(b1, b2, "bool: assert deep equal")
    t.Go(T).RefuteDeepEqual(b1, b3, "bool: refute deep equal")

    t.Go(T).RefuteNil(b1, "bool: refute nil")
}

func TestInt(T *testing.T) {
    t.Go(T).Assert(n1 == n2, "int: assert")
    t.Go(T).Refute(n1 == n3, "int: refute")

    t.Go(T).AssertEqual(n1, n2, "int: assert equal")
    t.Go(T).RefuteEqual(n1, n3, "int: assert equal")

    t.Go(T).AssertDeepEqual(n1, n2, "int: assert deep equal")
    t.Go(T).RefuteDeepEqual(n1, n3, "int: assert deep equal")

    t.Go(T).RefuteNil(n1, "int: refute nil")
}

func TestFloat(T *testing.T) {
    t.Go(T).Assert(f1 == f2, "float: assert")
    t.Go(T).Refute(f1 == f3, "float: refute")

    t.Go(T).AssertEqual(f1, f2, "float: assert equal")
    t.Go(T).RefuteEqual(f1, f3, "float: assert equal")

    t.Go(T).AssertDeepEqual(f1, f2, "float: assert deep equal")
    t.Go(T).RefuteDeepEqual(f1, f3, "float: assert deep equal")

    t.Go(T).RefuteNil(f1, "float: refute nil")
}

func TestString(T *testing.T) {
    t.Go(T).Assert(s1 == s2, "string: assert")
    t.Go(T).Refute(s1 == s3, "string: refute")

    t.Go(T).AssertEqual(s1, s2, "string: assert equal")
    t.Go(T).RefuteEqual(s1, s3, "string: assert equal")

    t.Go(T).AssertDeepEqual(s1, s2, "string: assert deep equal")
    t.Go(T).RefuteDeepEqual(s1, s3, "string: assert deep equal")

    t.Go(T).AssertLength(s1, 4, "string: assert length")
    t.Go(T).RefuteLength(s1, 0, "string: assert length")

    t.Go(T).RefuteNil(s1, "string: refute nil")
}

func TestArray(T *testing.T) {
    t.Go(T).AssertDeepEqual(a1, a2, "array: assert deep equal")
    t.Go(T).RefuteDeepEqual(a1, a3, "array: refute deep equal")

    t.Go(T).AssertLength(a1, 4, "array: assert length")
    t.Go(T).RefuteLength(a1, 0, "array: refute length")

    t.Go(T).RefuteNil(a1, "array: refute nil")
}

func TestMap(T *testing.T) {
    t.Go(T).AssertDeepEqual(m1, m2, "map: assert deep equal")
    t.Go(T).RefuteDeepEqual(m1, m3, "map: refute deep equal")

    t.Go(T).AssertLength(m1, 2, "map: assert length")
    t.Go(T).RefuteLength(m1, 0, "map: refute length")

    t.Go(T).RefuteNil(m1, "map: refute nil")
}

func TestInterface(T *testing.T) {
    t.Go(T).AssertEqual(i1, i2, "interface: assert equal")
    t.Go(T).RefuteEqual(i1, i3, "interface: refute equal")

    t.Go(T).AssertDeepEqual(i1, i2, "interface: assert deep equal")
    t.Go(T).RefuteDeepEqual(i1, i3, "interface: refute deep equal")

    t.Go(T).AssertLength(i1, 4, "interface: assert length")
    t.Go(T).RefuteLength(i1, 0, "interface: refute length")
    t.Go(T).RefuteLength(i3, 4, "interface: refute length")
    t.Go(T).RefuteLength(i4, 4, "interface: refute length")

    t.Go(T).AssertNil(i4, "interface: assert nil")
    t.Go(T).RefuteNil(i1, "interface: refute nil")
}

func TestType(T *testing.T) {
    t.Go(T).AssertEqual(t1, t2, "type: assert equal")
    t.Go(T).RefuteEqual(t1, t3, "type: refute equal")

    t.Go(T).AssertDeepEqual(t1, t2, "type: assert deep equal")
    t.Go(T).RefuteDeepEqual(t1, t3, "type: refute deep equal")

    t.Go(T).RefuteLength(t1, 0, "type: assert length")
    t.Go(T).RefuteLength(t4, 4, "type: refute length")

    t.Go(T).RefuteNil(t1, "type: refute nil")
    t.Go(T).RefuteNil(t4, "type: refute nil")
}

func TestPointer(T *testing.T) {
    t.Go(T).AssertEqual(p1.Value, p2.Value, "pointer: assert equal on value")
    t.Go(T).RefuteEqual(p1.Value, p3.Value, "pointer: refute equal on value")

    t.Go(T).AssertDeepEqual(p1, p2, "pointer: assert deep equal")
    t.Go(T).AssertDeepEqual(p5, p6, "pointer: assert deep equal")

    t.Go(T).RefuteDeepEqual(p1, p3, "pointer: refute deep equal")
    t.Go(T).RefuteDeepEqual(p5, p7, "pointer: refute deep equal")

    t.Go(T).RefuteLength(p1, 0, "pointer: refute length")
    t.Go(T).RefuteLength(p4, 4, "pointer: refute length")
    t.Go(T).RefuteLength(p5, 4, "pointer: refute length")

    t.Go(T).AssertNil(p4, "pointer: assert nil")
    t.Go(T).AssertNil(p8, "pointer: assert nil")

    t.Go(T).RefuteNil(p1, "pointer: refute nil")
    t.Go(T).RefuteNil(p5, "pointer: refute nil")
}

