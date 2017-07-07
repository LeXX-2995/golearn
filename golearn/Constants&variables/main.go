package main

import "fmt"

const p = "death and taxes"
const(
	_ = iota // 0
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (1*20)
	GB = 1 << (iota * 10) // 1 << (1*30)
	TB = 1 << (iota * 10) // 1 << (1*40)
)

func main(){
	const q  = 452
	fmt.Println("p - ", p)
	fmt.Println("q- ",q)
	print()

}


func print (){
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}