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
	f := data.NewMeta(g, 10e-5, 10e-3, 0.0, 0.0)
	fmt.Println(f.Train())
}
