package twosum

func twoSum(nums []int, target int) []int {
	//value, index
	var scaned = make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if v, ok := scaned[another]; ok {
			return []int{i, v}
		} else {
			scaned[nums[i]] = i
		}
	}
	return []int{0, 0}
}
