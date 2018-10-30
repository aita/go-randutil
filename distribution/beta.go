package distribution

// Beta returns a random number with beta distribution.
func Beta(alpha, beta float64) float64 {
	if alpha <= 0 {
		panic("beta: alpha must be greater than 0")
	}
	if beta <= 0 {
		panic("beta: beta must be greater than 0")
	}

	x := Gamma(alpha, 1.0)
	y := Gamma(beta, 1.0)
	return x / (x + y)
}
