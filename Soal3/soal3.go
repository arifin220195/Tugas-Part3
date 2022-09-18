package main

import "fmt"

func nilaiUjian(n int) string {
	// write your code
	switch {
	case n >= 85:
		return "A"
	case n >= 65:
		return "B"
	case n >= 35:
		return "C"
	default:
		return "Belajar lagi"

	}
}

func main() {
	fmt.Println(nilaiUjian(50)) // C
	fmt.Println(nilaiUjian(85)) // A
	fmt.Println(nilaiUjian(65)) // B
	fmt.Println(nilaiUjian(35)) // C
}
