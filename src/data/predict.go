package data

func (dd Meta)Predict (theta0, theta1 float64, X []float64) (y_hat []float64) {
	for _, val := range X {
		y_hat = append(y_hat, dd.EstimatePrice(theta0, theta1, val))
	}
	return
}