package two_sum

func twoSum(nums []int, target int) []int {

	for i, num1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-num1 == nums[j] {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], nums[i]}
		}
		m[nums[i]] = i
	}
	return nil
}
