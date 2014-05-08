package functional

var (
    b1 , b2 , b3      bool
    n1 , n2 , n3      int
    f1 , f2 , f3      float64
    s1 , s2 , s3      string
    m1 , m2 , m3      map[string]string

    i1 , i2 , i3 , i4 interface{}
    t1 , t2 , t3 , t4 Type
    p1 , p2 , p3 , p4 *Type
    p5 , p6 , p7 , p8 *[]int

    // array
    a1 , a2 , a3      [4]string

    // slice
    a4 , a5 , a6      []string
)

type Type struct {
    Value int
}

func init() {
    b1 = true
    b2 = true
    b3 = false

    n1 = 1
    n2 = 1
    n3 = 2

    f1 = 1.0
    f2 = 1.0
    f3 = 2.0

    s1 = "asdf"
    s2 = "asdf"
    s3 = "qwer"

    a1 = [4]string{ "a", "s", "d", "f" }
    a2 = [4]string{ "a", "s", "d", "f" }
    a3 = [4]string{ "q", "w", "e", "f" }

    a4 = []string{ "a", "s", "d", "f" }
    a5 = []string{ "a", "s", "d", "f" }
    a6 = []string{ "q", "w", "e", "f" }

    m1 = map[string]string{ "a": "s", "d": "f" }
    m2 = map[string]string{ "a": "s", "d": "f" }
    m3 = map[string]string{ "q": "w", "e": "r" }

    i1 = "asdf"
    i2 = "asdf"
    i3 = 1234

    t1 = Type{ Value: 1 }
    t2 = Type{ Value: 1 }
    t3 = Type{ Value: 0 }

    p1 = &Type{ Value: 1 }
    p2 = &Type{ Value: 1 }
    p3 = &Type{ Value: 0 }

    p5 = &[]int{ 1,2,3,4 }
    p6 = &[]int{ 1,2,3,4 }
    p7 = &[]int{ 5,6,7,8 }
}
