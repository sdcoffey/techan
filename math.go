package techan

// Min returns the smaller integer of the two integers passed in
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// Max returns the larger of the two integers passed in
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// Pow returns the first integer to the power of the second integer
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

// Abs returns the absolute value of the passed-in integer
func Abs(b int) int {
	if b < 0 {
		return -b
	}

	return b
}
