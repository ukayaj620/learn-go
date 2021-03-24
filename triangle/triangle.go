package triangle

import (
	"fmt"
	"strings"
)

func LowerRight(size int) {
	for i := 1; i <= size; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
