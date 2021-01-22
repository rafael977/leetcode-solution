package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n == 0 {
		return 0
	}
	k1 := (n + 1) / 2
	k2 := (n + 2) / 2
	return float64(findKthMin(nums1, nums2, k1)+findKthMin(nums1, nums2, k2)) / 2
}

func findKthMin(num1, num2 []int, k int) int {
	if len(num1) == 0 {
		return num2[k-1]
	}
	if len(num2) == 0 {
		return num1[k-1]
	}
	if k == 1 {
		if num1[0] < num2[0] {
			return num1[0]
		} else {
			return num2[0]
		}
	}
	k2 := k/2 - 1
	i1, i2 := k2, k2
	if len(num1)-1 < i1 {
		i1 = len(num1) - 1
	}
	if len(num2)-1 < i2 {
		i2 = len(num2) - 1
	}

	if num1[i1] < num2[i2] {
		return findKthMin(num1[(i1+1):], num2, k-i1-1)
	} else {
		return findKthMin(num1, num2[(i2+1):], k-i2-1)
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))

	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))

	fmt.Println(findMedianSortedArrays([]int{1}, []int{}))

	fmt.Println(findMedianSortedArrays([]int{}, []int{4}))
}
