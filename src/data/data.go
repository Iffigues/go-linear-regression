package data

import (
	"errors"
	"fmt"
	"math"
)

type Data struct {
	Km    []float64
	Price []float64
	Size  int
}

const (
	BitsPerWord = 32 << (^uint(0) >> 63)
	MaxInt      = 1<<(BitsPerWord-1) - 1
	MinInt      = -MaxInt - 1
)

type Meta struct {
	Theta0_init float64
	Theta1_init float64
	Epsilon     float64
	Seuil       float64
	L           Data
}

func NewData(a, b []float64) (g Data, err error) {
	if len(a) == 0 || len(b) == 0 || len(a) != len(b) {
		err = errors.New("not good things")
		return
	}
	g.Km = Normalize(a)
	fmt.Println(g.Km)
	g.Price = Normalize(b)
	fmt.Println(g.Price)
	g.Size = len(a)
	return
}

func NewMeta(data Data, seuil, epsilon, t0, t1 float64) (r Meta) {
	r.L = data
	r.Seuil = seuil
	r.Epsilon = epsilon
	r.Theta0_init = t0
	r.Theta1_init = t1
	return
}

func (dd Meta) EstimatePrice(theta0, theta1, x float64) (d float64) {
	return theta0 + (theta1 * x)
}

func (d Meta) Grad_theta0(theta0, theta1 float64) (res float64) {
	res = 0
	for i := 0; i <= d.L.Size-1; i++ {
		res = res + (d.EstimatePrice(theta0, theta1, d.L.Km[i]) - d.L.Price[i])
	}
	return d.Epsilon * (1.00 / float64(d.L.Size)) * res
}

func (d Meta) Grad_theta1(theta0, theta1 float64) (res float64) {
	res = 0
	for i := 0; i <= d.L.Size-1; i++ {
		res = res + (d.EstimatePrice(theta0, theta1, d.L.Km[i])-d.L.Price[i])*d.L.Km[i]
	}
	return d.Epsilon * (1.00 / float64(d.L.Size)) * res
}

func (d Meta) Train() (theta0, theta1 float64) {
	theta0 = d.Theta0_init
	theta1 = d.Theta1_init
	theta0_next := theta0 - d.Grad_theta0(theta0, theta1)
	theta1_next := theta1 - d.Grad_theta1(theta0, theta1)
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		diff0 := math.Abs(theta0_next - theta0)
		diff1 := math.Abs(theta1_next - theta1)
		fmt.Println(diff0, diff1, d.Seuil)
		if (diff0 <= d.Seuil) && (diff1 <= d.Seuil) {
			return theta0, theta1
		}
		theta0 = theta0_next
		theta1 = theta1_next
		theta0_next = theta0 - d.Grad_theta0(theta0, theta1)
		theta1_next = theta1 - d.Grad_theta1(theta0, theta1)
	}
	return
}
