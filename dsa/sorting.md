Bubble sort:

1. In the first pass largest element come at the right end
?example
1, 3, 4, 2, 5
1, 3, 2, 4, 5

thats why the innerloop runs till n-i-1
 n-i because in every pas last item will be sorted alreadt
 n-i-1 becaus , we need one more item to compare the last

example n can be compared with n+1


<!-- bubble sort python code -->

def bubblesort(arr):
    if len(arr)<=1:
         return arr

    for i in range(n):

      for j in range(0, n-i-1):
           if arr[j]>arr[j+1]:
               arr[j],arr[j+1]=arr[j+1],arr[j] 



<!--  selection sort -->

In selection sort  minimum element will be at left

Idea: finding min lement and swaping with the left item.

def selection_sort(arr):
    if len(arr)<=1:
        return arr

    for i in range (n):
      min=i

      for j in range(n)
       if arr[i]<arr[j]:
            i=j

       arr[i], arr[min]=arr[min],arr[i]



<!-- merge sort -->

def merge_sort():
  
  m=len(arr)//2

  left= merge_sort(arr[:m])
  right=merge_sort(arr[m:])

  return merge(left,right)


def merge(left,right):
    result=[]
    i=j=0

    while i< len(left) and j< len(right):
          
          if left[i]<right[j]:
             result.append(left[i])
             i=i+1
     
           else:
              result.append(right[j])
              j=j+1
    result.extend(left[i:])  
    result.extend(right[j:])

    return result         



def quick_sort(arr):

  n=len(arr)
  pivot=arr[0]

  left=[x for x in arr[1:] if x<pivot]
  right=[x for x in arr[1:] if x>pivot]
  
  return quick_sort(left) +[pivot]+ quick_sort(right)

def quick_sort(arr):
    if len(arr) <= 1:
        return arr
    
    pivot = arr[0]
    left = [x for x in arr[1:] if x < pivot]
    right = [x for x in arr[1:] if x >= pivot]

    return quick_sort(left) + [pivot] + quick_sort(right)

# Example
print(quick_sort([3, 6, 1, 5, 2]))


package main

import (
	"fmt"
)

// MergeSort recursively splits and merges
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge both halves
	return merge(left, right)
}

// Merging two sorted arrays
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Compare elements from left and right, and append smaller one
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{5, 3, 8, 1, 2}
	sorted := mergeSort(arr)
	fmt.Println("Merge Sort:", sorted)
}



package main

import (
	"fmt"
)

func cyclicSort(nums []int) {
	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			// Swap to correct index
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}
}

func main() {
	arr := []int{3, 1, 5, 4, 2}
	cyclicSort(arr)
	fmt.Println("Cyclic Sort:", arr)
}


package main

import (
    "fmt"
)

func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[0]
    left := []int{}
    right := []int{}

    for _, val := range arr[1:] {
        if val < pivot {
            left = append(left, val)
        } else {
            right = append(right, val)
        }
    }

    return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
    nums := []int{3, 6, 1, 5, 2}
    sorted := quickSort(nums)
    fmt.Println(sorted)
}


package main

import (
	"fmt"
)

package main

import (
	"fmt"
)

// quickSort sorts the array using the Hoare partition scheme
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	s := low
	e := high
	m := s + (e-s)/2
	pivot := arr[m]

	for s <= e {
		// Move left pointer until we find a value >= pivot
		for arr[s] < pivot {
			s++
		}
		// Move right pointer until we find a value <= pivot
		for arr[e] > pivot {
			e--
		}
		// Swap values and move pointers inward
		if s <= e {
			arr[s], arr[e] = arr[e], arr[s]
			s++
			e--
		}
	}

	// Recursively sort the two partitions
	quickSort(arr, low, e)
	quickSort(arr, s, high)
}

func main() {
	arr := []int{9, 4, 7, 6, 3, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}




func main() {
	arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Quick Sort:", arr)
}
