package main

import "fmt"

func main() {
	//变量
	var name string = "尼古拉斯赵四"
	var age  int = 16
	var f1 float32 =9.9
	var f2 float64 =9.999
	classes := 17
	var arr1 [10]int
	arr1[0]=1
	arr1[1]=2
	 arr2 := [3]string{"a","b","c"}

	/** slice 动态数组 */
	var arr3 [] int
	fmt.Println("hello,go")
	fmt.Println("name="+name)
	fmt.Println("age=",age)
	fmt.Println("f1=",f1)
	fmt.Println("f2=",f2*2)
	fmt.Println("classes=",classes)
	fmt.Println("调用函数值=",add1(classes))
	fmt.Println("调用函数值,传地址=",add2(&classes))
	a,b := getMultiAge(age)
	fmt.Println("获取连个返回值=",a,b)
	fmt.Println(arr1)
	fmt.Println(arr2)
	printDefer()
	arr3[0]=7

}

/**
  单个返回值
 */
func add1(age int) (newAge int)  {
	return age+1
}
/**
多个返回值
 */
func getMultiAge(age int) (first int,second int)  {
	return age+1,age*2
}

/**
  传变量的地址
 */
func add2(age *int) int {
	return *age+2
}

func printDefer()  {
	for i:=0 ;i<5; i++{
		defer fmt.Println(i)
	}
}