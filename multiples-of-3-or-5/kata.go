package kata

func Multiple3And5(number int) int {
	n := 0

	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			n += i
		}
	}

	return n
}
