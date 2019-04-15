package mymath

func Sum(x ...int) int {

	total := 0
	for _, value := range x {
		total += value
	}
	return total
}
