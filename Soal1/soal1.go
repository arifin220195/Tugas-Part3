package main

import "fmt"

func ganjilOrGenap(n int) string {
	// write your code
	if n%2 == 0 {
		return fmt.Sprintf("%d adalah bilangan genap", n)

	} else {
		return fmt.Sprintf("%d adalah bilangan ganjil", n)

	}

}

func main() {
	fmt.Println(ganjilOrGenap(2))  // 2 adalah bilangan genap
	fmt.Println(ganjilOrGenap(5))  // 5 adalah bilangan ganjil
	fmt.Println(ganjilOrGenap(11)) // 11 adalah bilangan ganjil
	fmt.Println(ganjilOrGenap(12)) // 12 adalah bilangan genap
}
