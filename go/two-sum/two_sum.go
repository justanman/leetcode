package two_sum

func twoSum(nums []int, target int) []int {
	indices := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indices[0] = i
				indices[1] = j
			}
		}
	}
	return indices
}
