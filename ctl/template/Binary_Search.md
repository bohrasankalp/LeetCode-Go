---
title: 2.11 Binary Search
type: docs
weight: 11
---

# Binary Search

- The classic way of writing binary search. Three points to note:
1. Loop exit condition, note that low <= high, not low < high.
2. The value of mid, mid := low + (high-low)>>1
3. Update of low and high. low = mid + 1, high = mid -1.

```go
func binarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
```

- A variant of binary search. There are 4 basic variants:
1. Find the first element equal to target, time complexity O(logn)
2. Find the last element equal to target, time complexity O(logn)
3. Find the first element greater than or equal to target, time complexity O(logn)
4. Find the last element less than or equal to target, time complexity O(logn)

```go
// Binary search for the first element equal to target, time complexity O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { //Find the first element equal to target
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// Binary search for the last element equal to target, time complexity O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // finds the last element equal to target
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// Binary search for the first element greater than or equal to target, time complexity O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // finds the first element greater than or equal to target
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// Binary search for the last element less than or equal to target, time complexity O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // finds the last element less than or equal to target
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

- Use binary search in basically sorted arrays. The classic solution can be solved, and the variant writing method can also be written. The common question type is to find the peak in the peak array and find the dividing point in the rotated ordered array. Question 33, Question 81, Question 153, Question 154, Question 162, Question 852

```go
func peakIndexInMountainArray(A []int) int {
	low, high := 0, len(A)-1
	for low < high {
		mid := low + (high-low)>>1
		// If mid is large, there is a peak on the left, high = m, if mid + 1 is large, there is a peak on the right, low = mid + 1
		if A[mid] > A[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
```

- max-min Maximum minimization problem. Find the maximum value if the minimum condition is met. Question 410, Question 875, Question 1011, Question 1283.


{{.AvailableTagTable}}