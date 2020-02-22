package main

import (
	"fmt"
	"linear/src/csvs"
	"linear/src/data"
)

func main() {
	a, b, err := csvs.GetCsv("./csv/data.csv")
	if err != nil {
		return
	}
	g, err := data.NewData(a, b)
	if err != nil {
		return
	}
	f := data.NewMeta(g, 10e-8, 10e-3, 0., 0.)
	theta0_hat, theta1_hat := f.Rescale(f.Train())
	y_hat := f.Predict(theta0_hat, theta1_hat, g.KmNatural)
	coeff_r2 := data.Coeff_r2(g.PriceNatural, y_hat)
	fmt.Println(coeff_r2)
}
