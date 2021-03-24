package triangle

import (
	"fmt"
	"strings"
)

func LowerRight(size int) {
	fmt.Println("\nLower Right")
	for i := 1; i <= size; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

func UpperRight(size int) {
	fmt.Println("\nUpper Right")
	for i := size; i > 0; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}

func LowerLeft(size int) {
	fmt.Println("\nLower Left")
	for i := 1; i <= size; i++ {
		fmt.Println(strings.Repeat(" ", size-i) + strings.Repeat("*", i))
	}
}

func UpperLeft(size int) {
	fmt.Println("\nUpper Left")
	for i := size; i > 0; i-- {
		fmt.Println(strings.Repeat(" ", size-i) + strings.Repeat("*", i))
	}
}
