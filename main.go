package main

import (
	"fmt"
	"github.com/infinitete/sort-starter/core"
)

func main() {
	ch := core.Rand(100, 100000)

	for _, s := range <-ch {
		for c := range s {
			for _, v := range c {
				fmt.Println(v)
			}
		}
	}
}
