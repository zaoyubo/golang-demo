package internal

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]bool)
	if m["a"] {
		fmt.Println("have")
	} else {
		fmt.Println("none")
	}
}
