package main

import (
	"fmt"
	"slices"

	"github.com/gitchander/gobal3/utils/iterate"
)

func main() {
	n := 7

	fmt.Println("Forward iterate:")
	for f := iterate.NewForward(n); f.Next(); {
		fmt.Println(f.Index())
	}
	fmt.Println()

	fmt.Println("Backward iterate:")
	for b := iterate.NewBackward(n); b.Next(); {
		fmt.Println(b.Index())
	}
	fmt.Println()

	fmt.Println("golang slices.Backward:")
	s := []string{"hello", "world", "my", "friend"}
	for i, x := range slices.Backward(s) {
		fmt.Println(i, x)
	}
}
