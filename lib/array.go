package sanzhizhouComponentLib

func RemoveDuplicatesIn64(list *[]int64) *[]int64 {
	if len(*list) == 0 {
		return list
	}

	// 使用 map 记录已存在的元素
	seen := make(map[int64]bool)
	result := make([]int64, 0, len(*list))

	for _, item := range *list {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return &result
}
