package main

import (
	"math"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkQuick(b *testing.B) {
	max := 3000000
	a := createAry(max)
	qsort(a)
}

func BenchmarkDispersion(b *testing.B) {
	max := 3000000
	a := createAry(max)
	Dispersion(a)
}

/*
func BenchmarkParallel(b *testing.B) {
	max := 30000000
	a := createAry(max)
	Parallel(a)
}
*/

func Dispersion(a []int) []int {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)

	ch := make(chan []int)
	go dispersion(a, int(math.Log2(float64(cpu))), ch)
	return <-ch

}

func dispersion(a []int, depth int, ch chan []int) {
	mid := len(a) / 2
	ch1 := make(chan []int)
	ch2 := make(chan []int)
	if depth > 0 {
		go dispersion(a[:mid], depth-2, ch1)
		go dispersion(a[mid:], depth-2, ch2)
		go func() {
			ch <- merge(<-ch1, <-ch2)
		}()
		return
	}

	go func(mid int) {
		qsort(a[mid:])
		ch1 <- a[mid:]
	}(mid)
	go func(mid int) {
		qsort(a[:mid])
		ch2 <- a[:mid]
	}(mid)

	go func() {
		ch <- merge(<-ch1, <-ch2)
	}()
}

func merge(left, right []int) []int {
	var x int
	start := 0
	result := make([]int, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if right[0] > left[0] {
			x, left = left[0], left[1:]
		} else {
			x, right = right[0], right[1:]
		}
		result[start] = x
		start++
	}

	for _, v := range left {
		result[start] = v
		start++
	}

	for _, v := range right {
		result[start] = v
		start++
	}

	return result
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
	c := make(chan []int)
	quickSortAsync(a, c)
	<-c
	close(c)
	/*
		var wg sync.WaitGroup
		wg.Add(1)
		go pqsort(a, &wg)
		wg.Done()
	*/
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
	wg.Wait()
}

func quickSortAsync(ary []int, c chan []int) {
	if len(ary) < 2 {
		c <- ary
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

	c1 := make(chan []int, 1)
	c2 := make(chan []int, 1)
	go quickSortAsync(ary[:left], c1)
	go quickSortAsync(ary[left+1:], c2)
	go func() {
		<-c1
		<-c2
		close(c1)
		close(c2)
		c <- ary
	}()
}

func createAry(max int) []int {
	ary := make([]int, max)
	for i := 0; i < max; i++ {
		ary[i] = rand.Int() % max
	}
	return ary
}
