package distribution

import (
	"math"
	"math/rand"
)

// Gamma returns a random number with gamma distribution.
func Gamma(k, theta float64) float64 {
	if k <= 0 {
		panic("gamma: k must be greater than 0")
	}
	if theta <= 0 {
		panic("gamma: theta must be greater than 0")
	}

	// 0 < k <= 1
	// GS Algorithm
	if k < 1 {
		for {
			// step 1
			u := rand.Float64()
			c := (math.E + k) / math.E
			// step 2
			p := c * u
			if p <= 1 {
				// step 3
				x := math.Pow(p, 1/k)
				v := rand.Float64()
				if v > math.Exp(-x) {
					continue // go to step 1
				}
				return x * theta
			}
			// step 4 (p>1)
			x := -math.Log((c - p) / k)
			if x < 0 {
				continue // go to step 1
			}
			v := rand.Float64()
			if v > math.Pow(x, k-1) {
				continue // go to step 1
			}
			return x * theta
		}
	} else if k == 1 {
		return rand.ExpFloat64() * theta
	} else {
		// k > 1
		// Marsaglia-Tsang Algorithm
		d := k - 1/3.
		c := 1 / math.Sqrt(9*d)
		for {
			x := rand.NormFloat64()
			v := 1 + c*x
			if v <= 0 {
				continue
			}
			v = v * v * v
			u := rand.Float64()
			if u < 1-0.0331*x*x*x*x {
				return d * v * theta
			}
			if math.Log(u) < 0.5*x*x+d*(1-v+math.Log(v)) {
				return d * v * theta
			}
		}
	}
}
