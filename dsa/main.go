package main

import "fmt"


func twoSum(nums []int, target int) []int {
    j:=1
	fmt.Print(j)
    for i:=range(len(nums)){
		for j:=i+1; j<len(nums);j++{
			if nums[i]+ nums[j]==target{
				return []int{i,j}
			}
		}
     fmt.Print(i)
    }

	return []int{0,0}
}





func main(){
nums:=[]int{1,2,3,4,5}
target:=6
twoSum(nums, target)
arr1 := [2]int{2,3}
        arr2 := [...]int{2,3}
        fmt.Println(arr1==arr2)

arr:=make([]int,5)	
fmt.Println(arr)	
arr=append(arr, 1)
fmt.Println(arr)
arr=append(arr, 2)
fmt.Println(arr)

fmt.Println(cap(arr))
fmt.Println(len(arr))

}