package main

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] {
			start = mid
		} else {
			end = mid
		}
	}
	if arr[start-1] < arr[start] && arr[start] > arr[start+1] {
		return start
	} else {
		return end
	}
}
