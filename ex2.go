package check3

func Ex2(list []int) (int, int) {
	var smalloc, temp int
	for i := 0; i <= len(list)-1; i++ {
		smalloc = i
		for j := i + 1; j <= len(list)-1; j++ {
			if list[j] < list[smalloc] {
				smalloc = j
			}
			temp = list[smalloc]
			list[smalloc] = list[i]
			list[i] = temp
		}
	}
	return list[0], list[len(list)-1]
}
