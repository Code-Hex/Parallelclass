package main

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"sync"
)

func main() {
	max := 10000000
	a := createAry(max)
	fmt.Println(a)

	var wg sync.WaitGroup
	pqsort(a, &wg)
	wg.Wait()

	fmt.Println(a)
}

func pqsort(ary []int, wg *sync.WaitGroup) {
	if len(ary) < 2 {
		return
	}

	left, right := 0, len(ary)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(ary)

	// Move the pivot to the right
	ary[pivotIndex], ary[right] = ary[right], ary[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range ary {
		if ary[i] < ary[right] {
			ary[i], ary[left] = ary[left], ary[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	ary[left], ary[right] = ary[right], ary[left]

	// Go down the rabbit hole
	wg.Add(2)
	go func() {
		pqsort(ary[:left], wg)
		wg.Done()
	}()
	go func() {
		pqsort(ary[left+1:], wg)
		wg.Done()
	}()
}

func createAry(max int) (ary []int) {
	for i := 0; i < max; i++ {
		ary = append(ary, rand.Int()%max)
	}
	return
}
