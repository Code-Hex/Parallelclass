package main

import (
	"fmt"
	"math/rand"
)

func main() {

	max := 10000000
	a := createAry(max)
	fmt.Println(a)

	ch := make(chan []int)
	pqsort(a, ch)

	fmt.Println(<-ch)
}

func pqsort(ary []int, ch chan []int) {
	if len(ary) < 2 {
		ch <- ary
		return
	}

	ch1 := make(chan []int)
	ch2 := make(chan []int)

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
	go pqsort(ary[:left], ch1)
	go pqsort(ary[left+1:], ch2)

	go func() {
		<-ch1
		<-ch2
		ch <- ary
	}()
}

func createAry(max int) (ary []int) {
	for i := 0; i < max; i++ {
		ary = append(ary, rand.Int()%max)
	}
	return
}
