package main
import ("fmt")

func twoSum(nums []int, target int) [2]int {
    n := len(nums)
    arr1 := [2]int{}
    
    for i:=0;i<n;i++ {
        if nums[i]+nums[i+1]==target {
            arr1[0] = i
            arr1[1] = i+1
        }
    }

    return arr1
}
