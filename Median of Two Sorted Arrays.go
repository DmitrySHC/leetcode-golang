/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n))  =((
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	first_slice, second_slice := []int{-10, 0, 3, 5, 16, 88, 100}, []int{0, 2, 7, 42, 77, 99}
	median := FindMedianSortedArrays(first_slice, second_slice)
	median2 := findMedianSortedArrays2(first_slice, second_slice)
	fmt.Printf("%f\n", median)
	fmt.Printf("%f\n", median2)

}

func FindMedianSortedArrays(nums1, nums2 []int) (median float64) {
	// slow python style
	var (
		i, j int
		res  []int
	)
	for {
		if len(nums1) == i {
			res = append(res, nums2[j:]...)
			break
		} else if len(nums2) == j {
			res = append(res, nums1[i:]...)
			break
		} else {
			if nums1[i] >= nums2[j] {
				res = append(res, nums2[j])
				j++
			} else {
				res = append(res, nums1[i])
				i++
			}
		}
	}
	if len(res)%2 == 0 {
		median = (float64(res[(len(res)/2)-1]) + float64(res[len(res)/2])) / 2.0
		return
	} else {
		median = float64(res[len(res)/2])
	}
	return
}

func findMedianSortedArrays2(nums1, nums2 []int) (median float64) {
	// with sorting (slow too)
	total := append(nums1, nums2...)
	sort.Ints(total)
	half := len(total) / 2
	if len(total)%2 == 0 {
		median = (float64(total[half-1]) + float64(total[half])) / 2.0
		return
	} else {
		median = float64(total[half])
	}
	return
}
