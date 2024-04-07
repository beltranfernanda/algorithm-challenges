package challenge1

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var x int
	output := []int{0, -1}
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		x = target - v
		value, exist := m[x]
		if exist == true && value != i {
			output[0] = i
			output[1] = value
			return output
		}
	}
	return output
}

// Most efficient solution
func twoSum(nums []int, target int) []int {
	// value to index
	m := make(map[int]int)

	for i, n := range nums {
		if index, exists := m[target-n]; exists && index != i {
			return []int{i, index}
		}

		m[n] = i
	}
	return []int{}
}
