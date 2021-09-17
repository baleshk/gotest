package calc

func Add(number ...int) int {
	s := 0
	for _, n := range number {
		s += n
	}
	return s
}
