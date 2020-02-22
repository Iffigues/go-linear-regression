package data

import "math"

func Coeff_r2(y, y_hat []float64) (r2 float64) {
	mean_y := Moyenne(y)
	error_model := 0.
	variance := 0.
	l := len(y)
	for j := 0; j < l; j++ {
		error_model += math.Pow(y[j]-y_hat[j], 2)
	}
	for j := 0; j < l; j++ {
		variance += math.Pow(y[j]-mean_y, 2)
	}
	r2 = 1. - (error_model / variance)
	return
}
