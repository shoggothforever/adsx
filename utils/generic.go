package utils

import (
	"math/rand"
	"skiplist/config"
)

func Gt[T config.Generic](v1, v2 T) bool {

	return v1 > v2
}
func Eq[T config.Generic](v1, v2 T) bool {

	return v1 == v2
}
func Lq[T config.Generic](v1, v2 T) bool {

	return v1 < v2
}
func RandomLevel() uint {
	level := (uint)(1)
	for rand.Int()%4 == 0 && level < config.KMaxHeight {
		level++
	}
	return level
}
