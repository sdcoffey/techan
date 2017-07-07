package talib4g

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Pow(i, j int) int {
	p := 1
	for j > 0 {
		if j&1 != 0 {
			p *= i
		}
		j >>= 1
		i *= i
	}

	return p
}

func Abs(b int) int {
	if b < 0 {
		return -b
	}

	return b
}
