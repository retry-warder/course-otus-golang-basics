package main

import "fmt"

func main() {
	var size int
	fmt.Println("ЗАДАЙТЕ РАЗМЕР ДОСКИ:")
	fmt.Scanf("%d", &size)
	fmt.Println()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i%2+j%2)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
