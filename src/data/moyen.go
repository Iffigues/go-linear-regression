package data

import (
	"math"
)

func Moyenne(a []float64) (t float64) {
	size := len(a)
	for _, val := range a {
		t = t + val
	}
	return t / float64(size)
}

func Std(a []float64) (std float64) {
	size := len(a)
	mean := Moyenne(a)
	for j := 0; j < size; j++ {
		std += math.Pow(a[j]-mean, 2)
	}
	return math.Sqrt(std / float64(size))
}

func Normalize(a []float64) (b []float64) {
	m := Moyenne(a)
	s := Std(a)
	for _, val := range a {
		b = append(b, (val-m)/s)
	}
	return
}
