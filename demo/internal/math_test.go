package internal

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	var a int = 10
	var b int = 3
	var c int = 6
	fmt.Println(a / b) // 3
	fmt.Println(a / c) // 1
}
