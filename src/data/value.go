package data

func (f Meta) PredictValue(km float64) (price float64) {
	theta0_hat, theta1_hat := f.Rescale(f.Train())
	price = f.EstimatePrice(theta0_hat, theta1_hat, km)
	return
}
