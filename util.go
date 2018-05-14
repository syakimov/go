package main

import (
	"fmt"
	"strings"
)

func printHyrachy(r *Employee, lvl int) {
	if r == nil {
		return
	}

	fmt.Println(strings.Repeat("  ", lvl) + r.Name)

	for _, s := range r.Subordinates {
		printHyrachy(s, lvl+1)
	}
}
