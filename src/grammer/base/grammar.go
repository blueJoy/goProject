package base

import (
	"fmt"
	"os"
	"io"
	"expvar"
	"strconv"
	"log"
)


/*
	总结：Go和C一样，类型都是基于值传递，想要修改变量的值，只能传递指针 *

		Go 中b=&a，赋值的是引用，*a 指针也是引用

		java中基本数据类型是值传递【不能直接修改值，只能修改后返回】，对象属于引用传递【可以直接修改值】

 */


//定义常量   iota每次const初始为0，每次调用+1，相同表达式可以省略
const  (
	a = 1 + iota
	b
	c
)

func arraySlice (){

	arr := [10] int {1,2,3,4,5,6,7,8,9,10}

	//切片 5代表前5个元素，不是下标
	mySlice := arr[:5]

	//创建切片
	myS1 := make([] int ,5)

	//追加元素
	myS1 := append(myS1,1,2,3,4)

	//追加切片
	myS1 := append(myS1,mySlice...)

	//内容复制，以较少的函数复制
	copy(myS1,mySlice)

	fmt.Println("array is :")
	for _,v := range arr {
		fmt.Print(v)
	}

	fmt.Println("slice is :")
	for _,v := range mySlice {
		fmt.Print(v)
	}

}

func mapDemo()  {

	//定义数据结构
	type PersonInfo struct {

		ID string
		Name string
		Address string
	}

	//创建容量为100的map
	var personDB map[string] PersonInfo = make(map[string] PersonInfo,100)

	//添加数据
	personDB["12345"] = PersonInfo{"12345","Tom","Room 203"}
	personDB["1"] = PersonInfo{"1","Jack","Room 101"}

	//查找数据
	person, ok := personDB["1234"]

	// ok是一个bool类型，返回true表示找到对应的数据
	if ok {
		fmt.Println("Found person",person.Name)
	}else{
		fmt.Println("Did not found person with ID 1234")
	}


	//删除元素,传入的key不能为nil
	delete(personDB,"1234")
}

//args ...int 估计参数类型只能为int  args ...interface{} 可以传递任意类型的数据
func multiArgs(args ...interface{}){
	for _,arg := range args {
		//获取参数类型
		switch arg.(type) {
		case int:
			fmt.Println(arg,"is an int value.")
			//fallthrough 可以跳转执行下一个case,但是不能用在type上
			//fallthrough
		case string:
			fmt.Println(arg,"is an string value.")
		case int64:
			fmt.Println(arg,"is an int64 value.")
		default:
			fmt.Print(arg,"is an unkown type.")
		}
	}
}

//匿名函数
func anonymousFunc(){

	//匿名函数
	f := func(x,y int) int{
		return x +y
	}

	//闭包
	var j int = 5

	//返回值为函数
	a := func() (func()) {
		var i int = 10
		return func() {
			fmt.Errorf("i,j: %d, %d\n",i,j)
		}
	}()  //末尾小括号不能少

	a()

	j *= 2

	a()

	/*
	执行结果为：
		i,j: 10, 5
		i,j: 10, 10
	 */
}

//defer
func copyFile(dst, src string) (w int64,err error){
	//打开文件句柄
	srcFile,err := os.Open(src)
	if err != nil {
		return
	}

	//defer 做清理释放功能，后面也可以是匿名函数，defer语句按照先进后厨原则执行。
	defer srcFile.Close()

	dstFile,err := os.Create(dst)
	if err != nil {
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile,srcFile)
}

//  ---------------   对象内容  -------------------

//定义一个矩形
type Rect struct {
	x,y float64
	width,height float64
}

//定义计算矩形面积的方法
func (r *Rect) Area() float64{
	return r.width * r.height
}

//创建对象的四种方式
func createObj(){

	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0,0,100,200}   //显式初始化
	rect4 := &Rect{width:100,height:200}  //显式初始化

	//计算面积     未进行显式初始化，默认初始化为0值
	rect1.Area()
	rect2.Area()
	rect3.Area()
	rect4.Area()
}


// -------------   匿名组合  --------------------


type Base struct {
	Name string
}
func (base *Base) Foo(){}
func (base *Base) Bar(){}

//Foo 继承了Base,并改写了Bar()方法
type Foo struct {
	Base  //继承base
	Age int
	*log.Logger   //这样在foo类型的所有成员方法中可以直接调用log.Logger提供的方法
	//Name string  如果组合类型和被组合类型都包含Name,成员。所有都访问的是Foo的Name变量，Base的Name相当于被隐藏了。
}
func (foo *Foo) Bar(){
	foo.Base.Bar()
	fmt.Println("我改写了")
	//调用log.Logger的Println方法
	foo.Println("log-->我改写了")
}

