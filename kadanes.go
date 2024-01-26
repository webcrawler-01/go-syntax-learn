func maxSubArray(nums []int) int {
    var maxSum int=nums[0];
    var currentSum int = nums[0] // Initialize with the first element

    for i := 1; i < len(nums); i++ {
        if currentSum+nums[i] >= nums[i] {
            currentSum = currentSum + nums[i]
        } else {
            currentSum = nums[i]
        }

        maxSum = int(math.Max(float64(currentSum), float64(maxSum)))
    }

    return maxSum
}
