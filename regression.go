package main

import "fmt"

// 配列を受け取って最小二乗法で回帰問題を解く
func regression(arr_x []float64, arr_y []float64) (float64, float64) {
	var a = 0.0
	var b = 0.0
	n := len(arr_x)
	var sum_x = 0.0
	var sum_y = 0.0
	var sum_xy = 0.0
	var sum_x2 = 0.0
	for i := 0; i < n; i++ {
		sum_x += arr_x[i]
		sum_y += arr_y[i]
		sum_xy += arr_x[i] * arr_y[i]
		sum_x2 += arr_x[i] * arr_x[i]
	}
	nf := float64(n)
	a = (nf * sum_xy - sum_x * sum_y) / (nf * sum_x2 - sum_x * sum_x)
	b = (sum_y - a * sum_x) / nf
	return a, b
}

func main() {
	x := []float64 {1, 2, 3}
	y := []float64 {1, 2, 3}
	var a, b = regression(x, y)
	fmt.Printf("a = %f, b = %f\n", a, b)
}