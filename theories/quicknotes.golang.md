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



📦 What is heap allocation?
In programming, memory can be allocated in:

Stack: Temporary, fast, cleaned up when a function exits.

Heap: Longer-lived memory, shared across goroutines, cleaned up by Garbage Collector.

🧹 Yes, heap-allocated memory is tracked by the garbage collector, so using it increases GC load.



🚫 Why not overuse new?
Using new too much can:

Increase memory pressure – the GC has to work harder and run more often.

Make code harder to reason – more pointer usage means more chances for bugs (e.g. nil pointers, unexpected side effects).

Hurt performance – Stack allocations are cheaper and faster than heap allocations.

So:
✅ Use value types (non-pointers) and stack memory where possible.
✅ Use new only when you need heap-allocated pointers.

🧠 What does "safe heap allocation" mean in Go?
When you use new, Go allocates memory on the heap and returns a pointer to it. This is considered “safe” because:

The data will not be destroyed when the function returns.

Go's garbage collector (GC) will manage this memory — it’ll clean it up automatically when it’s no longer needed.









you can compare two pointers in Go. It’s both valid and useful in many cases. Here's how it works and what it really means:

✅ Pointer Comparison in Go
You can compare pointers using the standard comparison operators:

== (equal)

!= (not equal)

go
Copy
Edit
a := 10
b := 10

p1 := &a
p2 := &b
p3 := &a

fmt.Println(p1 == p2) // false: different memory addresses
fmt.Println(p1 == p3) // true: same address
⚠️ Note: Even though a and b have the same value (10), their addresses are different, so p1 == p2 is false.

var m1, m2 map[string]int
fmt.Println(m1 == m2) // ❌ Invalid — will not compile

fmt.Println(&m1 == &m2) // ✅ OK — comparing addresses


Comparison	Meaning
p1 == p2	Do both pointers point to same address?
*p1 == *p2	Do the values they point to match?