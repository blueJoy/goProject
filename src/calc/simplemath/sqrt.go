package simplemath

import "math"

//开方
func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)

}
