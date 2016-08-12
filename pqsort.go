package main

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkDispersion(b *testing.B) {
	max := 30000000
	a := createAry(max)
	Dispersion(a)
}

func BenchmarkParallel(b *testing.B) {
	max := 30000000
	a := createAry(max)
	Parallel(a)
}

func Dispersion(a []int) {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)
	length := len(a) / cpu

	var wg sync.WaitGroup

	wg.Add(cpu)
	for i := 0; i < cpu; i++ {
		go func(i int) {
			qsort(a[i*length : (i+1)*length])
			wg.Done()
		}(i)
	}
}

func qsort(ary []int) {
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
	qsort(ary[:left])
	qsort(ary[left+1:])
}

func Parallel(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan []int)
	go pqsort(a, ch)
	<-ch
}

func pqsort(ary []int, ch chan []int) {
	if len(ary) < 2 {
		ch <- []int{}
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
	ch1 := make(chan []int)
	ch2 := make(chan []int)
	go pqsort(ary[:left], ch1)
	go pqsort(ary[left+1:], ch2)
	<-ch1
	<-ch2
	close(ch1)
	close(ch2)
	ch <- ary
}

func createAry(max int) []int {
	ary := make([]int, max)
	for i := 0; i < max; i++ {
		ary[i] = rand.Int() % max
	}
	return ary
}
