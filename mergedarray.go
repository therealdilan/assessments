package main 

import (
  "sort"
  "fmt"
)

func main() {
  merge([]int{1,2,3,4,5,6,7},4,[]int{7,100,5,6},2)
}


func merge(nums1 []int, m int, nums2 []int, n int)  {
  copy(nums1[:m],nums2[:n])
  sort.Ints(nums1)
  fmt.Println(nums1)
}
