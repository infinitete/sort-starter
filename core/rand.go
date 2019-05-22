package core

import (
	"math/rand"
	"sync"
)

type A chan []int
type B chan []A

func Rand(chans int, size int) B {
	ch := make(B)
	numbers := make([]A, chans)

	var wg sync.WaitGroup

	for i := 0; i < chans; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			var nsc A
			nsc = randSingle(size)
			numbers[x] = nsc
		}(i)
	}

	go func() {
		wg.Wait()
		ch <- numbers
		defer close(ch)
	}()

	return ch
}

func randSingle(size int) A {
	ch := make(A, size)

	go func() {
		numbers := make([]int, size)
		for i := 0; i < size; i++ {
			numbers[i] = rand.Int()
		}
		ch <- numbers
		defer close(ch)
	}()

	return ch
}
