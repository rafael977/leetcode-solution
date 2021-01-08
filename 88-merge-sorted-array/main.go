package main

import "fmt"

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	for i, v := range nums2 {
// 		j := m + i
// 		nums1[j] = v
// 		for j > 0 {
// 			if nums1[j] > nums1[j-1] {
// 				break
// 			}
// 			nums1[j] = nums1[j] ^ nums1[j-1]
// 			nums1[j-1] = nums1[j] ^ nums1[j-1]
// 			nums1[j] = nums1[j] ^ nums1[j-1]
// 			j--
// 		}
// 	}
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	p, p1, p2 := n+m-1, m-1, n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

func main() {
	nums1, nums2 := []int{1, 2, 3, 0, 0, 0, 0}, []int{1, 2, 5, 6}
	merge(nums1, 3, nums2, 4)
	fmt.Printf("%v\n", nums1)

	nums1, nums2 = []int{1}, []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Printf("%v\n", nums1)
}
