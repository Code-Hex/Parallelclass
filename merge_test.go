package main

import (
  "math/rand"
  "testing"
)

// My function name initial is uppercase.


var testInt = createAry(3000000)

func P_MergeSort_old(x []int, c chan []int) {
  if len(x) < 2 {
    c <- x
    return
  }

  c1 := make(chan []int, 1)
  c2 := make(chan []int, 1)
  mid := len(x) / 2

  go P_MergeSort_old(x[:mid], c1) //split 0 ~ mid
  go P_MergeSort_old(x[mid:], c2)//split mid ~ END of x[]

  go func() {
    c <- Merge(<-c1, <-c2)  //block untill "P_MergeSort" and "Merge()" is over.
  }()
}

func P_MergeSort(x []int, c chan []int) {
  if len(x) < 2 {
    c <- x
    return
  }

  if len(x) < 10000 {
    c <- MergeSort(x)
    return
  }

  c1 := make(chan []int, 1)
  c2 := make(chan []int, 1)
  mid := len(x) / 2

  go P_MergeSort(x[:mid], c1) //split 0 ~ mid
  go P_MergeSort(x[mid:], c2)//split mid ~ END of x[]

  go func() {
    c <- Merge(<-c1, <-c2)  //block untill "P_MergeSort" and "Merge()" is over.
  }()
}


func Merge(left,right []int)(ret []int) {
    ret = []int{}
    for len(left) > 0 && len(right) > 0 {
        var x int
        if right[0] > left[0] {
            x, left = left[0],left[1:]
        } else {
            x, right = right[0], right[1:]
        }
        ret = append(ret,x)
    }

    ret = append(ret, left...)
    ret = append(ret, right...)
    return
}

func sort(left, right []int)(ret []int) {
    if len(left) > 1 {
        l, r := split(left)
        left = sort(l,r)
    }
    if len(right) > 1 {
        l, r := split(right)
        right = sort(l,r)
    }

    ret = Merge(left, right)
    return
}

func split(values []int)(left, right []int) {
    left = values[:len(values)/2]
    right = values[len(values)/2:]
    return
}

func MergeSort(values []int)(ret []int) {
    left, right := split(values)
    ret = sort(left, right)
    return
}





func createAry(max int) (ary []int) {
  for i := 0; i < max; i++ {
    ary = append(ary, rand.Int()%max)
  }
  return
}



func BenchmarkP_MergeSort(b *testing.B) {
  c := make(chan []int, 1)

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    P_MergeSort(testInt,c)
    <- c
  }
}

func BenchmarkMergeSort(b *testing.B){

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    MergeSort(testInt)
  }
}

func BenchmarkP_MergeSort_old(b *testing.B){
  c := make(chan []int, 1)

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    P_MergeSort_old(testInt,c)
    <- c
  }
}




//func main () {
  

/*
  ori_data := []int{4,8,3,9,7,6,3,9,0,996,86,759,47,15,18,787,365,418915,48,9} //random data
  c := make(chan []int)

  go P_MergeSort(ori_data, c)

  fmt.Println(<-c) //block untill the "P_MergeSort" is over.
*/


