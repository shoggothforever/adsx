package config

const KMaxHeight uint = 12

// 可排序的变量类型的泛型列表
type Generic interface {
	Ordered
}
type Number interface {
	Integer | Float
}
type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}
