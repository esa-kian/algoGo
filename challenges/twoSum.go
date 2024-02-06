// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.


func twoSum(nums []int, target int) []int {
    var result []int
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                result = append(result, i, j)
            }
        }
    }
    return result
}
