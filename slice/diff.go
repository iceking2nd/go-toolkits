package slice

func Diff[T comparable](a []T, b []T) ([]T, []T, []T) {
	aMap := make(map[T]bool)
	bMap := make(map[T]bool)
	var (
		common []T
		aOnly  []T
		bOnly  []T
	)

	// 将a和b中的元素分别存储在map中
	for _, item := range a {
		aMap[item] = true
	}
	for _, item := range b {
		bMap[item] = true
	}

	// 遍历a和b中的元素，将其分类到aOnly、bOnly和common三个切片中
	for item := range aMap {
		if !bMap[item] {
			aOnly = append(aOnly, item)
		} else {
			common = append(common, item)
		}
	}
	for item := range bMap {
		if !aMap[item] {
			bOnly = append(bOnly, item)
		}
	}

	return aOnly, bOnly, common
}
