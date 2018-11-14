package distribution

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamma(t *testing.T) {
	const N = 5000000

	for _, tt := range []struct {
		k, theta float64
	}{
		{0.5, 2.0},
		{1.0, 2.0},
		{1.5, 2.0},
	} {
		var mu, sigma2 float64
		for i := 0; i < N; i++ {
			x := Gamma(tt.k, tt.theta)
			mu += x
			sigma2 += x * x
		}
		mu = mu / N
		sigma2 = sigma2/N - mu*mu

		assert.InDelta(t, tt.k*tt.theta, mu, 0.01, "k=%f, theta=%f", tt.k, tt.theta)
		assert.InDelta(t, tt.k*math.Pow(tt.theta, 2), sigma2, 0.01, "k=%f, theta=%f", tt.k, tt.theta)
	}
}
