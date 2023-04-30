package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {

	if n < 1 {
		return 0, errors.New("n cannot be less than 1")
	}
	if n == 1 {
		return 0, nil
	}

	steps := 0
	for ; n > 1; steps++ {

		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}

	return steps, nil
}
