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
// time O(n^2)
// space O(1)

func twoSumMap(nums []int, target int) []int {
  table := make(map[int]int)
  indices := make([]int, 2)
  for i := 0; i < len(nums); i++ {
    table[nums[i]] = i
  }
  for i := 0; i < len(nums); i++ {
    complement := target - nums[i]
    if _, ok := table[complement]; ok {
      if table[complement] != i {
        indices[0] = i
        indices[1] = table[complement]
        return indices
      }
    }
  }
  return indices
}
// time O(n)
// space O(1)
