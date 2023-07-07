package main

import "fmt"

func main() {
	// Masukkan batas atas untuk mencari kelipatan 7
	batas := 100

	fmt.Println("Kelipatan 7 antara 1 dan", batas, ":")

	for i := 1; i <= batas; i++ {
		if i%7 == 0 {
			fmt.Print(" ", i)
		}
	}
}
