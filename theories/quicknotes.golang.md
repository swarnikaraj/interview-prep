Sort in ascending order in place
	sort.Ints(nums)

    sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // Reverse order
	})
 sort in copied slice

 sort_num=append([]int{},nums...)   
 nums := []int{-1, 0, 1, 2, -1, -4}

	// Create a copy of the original slice
	sortedNums := append([]int{}, nums...) // Clone the slice

	// Sort the copy
	sort.Ints(sortedNums)


    / Function that takes multiple integers
func sum(nums ...int) {
	fmt.Println(nums)
}

func main() {
	arr := []int{1, 2, 3}
	sum(arr...) // Spread-like behavior, passing elements of arr as separate arguments
}