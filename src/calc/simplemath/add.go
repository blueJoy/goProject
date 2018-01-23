package simplemath

import "fmt"

//枚举定义   无enum字段
const(
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Staurday
	numberOfDays    //不能才calc中使用
)

//go 根据方法和属性的大小写控制访问权限，   小写只能本包访问，大写可以被其他包访问
func Add(a int,b int) int {
	fmt.Println(Staurday)
	return a+b
}
