package twosum

func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
	for i := range nums {
		tmp := target - nums[i]
		j, ok := m[tmp]
		if ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return nil
}