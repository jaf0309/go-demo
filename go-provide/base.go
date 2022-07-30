package main

import "fmt"

func rangeNums() {

}

func main() {
	var a = [...]int{1, 2, 3}
	fmt.Println(a[0])
	for k, v := range a {
		fmt.Println("k:", k)
		fmt.Println("v:", v)
	}

}
