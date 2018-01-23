最近开始学习GoLang,对于学习进行一下总结：

###一、基本语法
####1. 变量
##### 声明：
```
var a1 int
var a2 string
var a3 [10]int  // 数组
var a4 []int // 数组切片
var a5 struct {    //定义数据结构
f int
}
var a6 *int // 指针
var a7 map[string]int  // map，key为string类型，value为int类型
var a8 func(a int) int
//多个变量
var m1,m2 string
var (
v1 int
v2 string
)
```
##### 初始化：
```
var a1 int = 10
var a2 = 10
a3 := 10

//数组
array := [5]int{1,2,3,4,5}
//切片   创建元素个为5，初始值为0，预留10个元素的存储空间
slice := make([]int,5,10)
//map  interface可以是任何类型
map := make(map[string]interface{})

//go支持多重赋值，避免通过中间变量或者^的方式
i,j = j,i
```
匿名变量：
go语言支持多重返回值，如果存在不需要的返回值，可以通过_来忽略
```
func GetInfo()(name string,age int,description string){
return "Job",22
}
//只获取age的值
_,age,_ := GetInfo()
```
#####常量：
```
//单一常量
const ROOT = "root"

//go中没有枚举的类型，可以通过常量的方式实现。主要iota默认初始为0，每次出现iota加1，出现cost 则再次初始化为0
const(
Sunday = iota        //如果后续的赋值一样，可以省略
Monday
Wednesday
Thursday
Friday
Saturday
days            //由于访问权限问题，小写其他包不能访问
)
```
#####类型：
基本类型：
 布尔类型： bool
  整型： int8 、 byte 、 int16 、 int 、 uint 、 uintptr 等
  浮点类型： float32 、 float64
  复数类型： complex64 、 complex128
  字符串： string
  字符类型： rune
  错误类型： error
复合类型：
  指针（pointer）
  数组（array）
  切片（slice）
  字典（map）
  通道（chan）
  结构体（struct）
  接口（interface）

#####流程控制
```
//if   > 1,= 0, < -1    主要演示if,肯定有更优写法
func compare(x,y int) int{
if x > y {
return 1
}else if x ==y {
return 0
}else {
return -1
}
}

//switch
switch i {
case 0:
fmt.Print("0")
case 1:
fmt.Print("1")
fallthrouth       //会继续执行紧跟着的下一个case
case 2, 3, 4:
fmt.Print(2,3,4)
default:
fmt.Print("default')
}

//for 循环  也可以用break和continue控制循环
for{
fmt.Print("hhahahah")     //不带参数为无线循环
}

//普通循环5次
for i := 0 ; i < 5; i ++{
fmt.Print("i = "+i)
}

//类似迭代循环，第一个参数为下标，第二个为值
array := [] int{1,2,3}
for i,v := range array{
fmt.Print("index = ",i," value = ",value)
}

//goto 跳转【很少使用】
func example{
i := 0
HERE:    //做一个标记
fmt.Println(i)
i++
if i < 10 {
goto HERE    //跳转到标记处
}
}

```
#####函数
```
//基本函数
func Add(x,y int) int{
}

//不定参数   必须是最后一个参数,任意不定参数可以用...interface{}表示
func myfunc(header *Header,param ...string){
}

//匿名函数
f := func(x,y) int{
return x*y
}

//闭包
i := 20
a := func()(func()) {
var i int = 10   //这个变量只能闭包中访问
return func() {
fmt.Printf("i, j: %d, %d\n", i, j)
}
}()
```
#####错误处理
error 接口方式
```
func Add(x,y int) int,error{
}
sum,err :=Add(1,2)
if err != nil{
//错误处理
}
```
defer 方式
对于释放资源类的做法，go中使用defer关键字。代码块中可以存在多个defer，执行方式类似栈，先声明的后执行。即使出现异常也会执行

panic()和recover()关键字
panic会终止程序流程，而recover()会终止错误流程。可以配合保证程序运行中处理异常