package functions

import "fmt"

// 1 1 2 3 5 8 13 21
func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func PrintFibo() int {
	fmt.Print("\n这是斐波那契计算:\n")
	var a = Fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Printf("i => %d\n", a())
	}
	fmt.Print("斐波那契计算 -- END \n")
	return a()
}

func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func PrintSum() {
	a := Adder()
	fmt.Print("\n这是和计算:\n")
	for i := 0; i <= 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d \n", i, a(i))
	}
	fmt.Print("和计算结束 -- END\n")
}
