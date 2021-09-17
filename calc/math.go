package calc

import "errors"

func Add(number ...int) (int, error) {
	s := 0
	if len(number) < 2 {
		return 0, errors.New("numbers provided are less than two")
	}
	for _, n := range number {
		s += n
	}
	return s, nil
}
