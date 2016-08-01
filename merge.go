package main

import (
  "fmt"
)

// My function name initial is uppercase.

func P_merge_sort(x []int, c chan []int) {
  if len(x) < 2 {
    c <- x
    return
  }

  c1 := make(chan []int, 1)
  c2 := make(chan []int, 1)
  mid := len(x) / 2

  go P_merge_sort(x[:mid], c1) //split 0 ~ mid
  go P_merge_sort(x[mid:], c2)//split mid ~ END of x[]

  go func() {
    c <- Merge(<-c1, <-c2)  //block untill "P_Merge_sort" and "Merge()" is over.
  }()
}


func Merge(left []int, right []int) []int {
  result := make([]int,len(left)+len(right))
  var i, j int;

  for i < len(left) && j < len(right) {
    if left[i] <= right[j] {
      result[i+j] = left[i]
      i++
    } else {
      result[i+j] = right[j]
      j++
    }
  }

  for i < len(left) {
    result[i+j] = left[i]
    i++
  }

  for j < len(right) {
    result[i+j] = right[j]
    j++
  }

  return result
}


func main () {
  ori_data := []int{4,8,3,9,7,6,3,9,0,996,86,759,47,15,18,787,365,418915,48,9} //random data
  c := make(chan []int)

  go P_merge_sort(ori_data, c)

  fmt.Println(<-c) //block untill the "P_merge_sort" is over.

}
