package main

import (
	"fmt"
	// "math"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.0}
	sort.Float64s(temperatures)
	fmt.Println(temperatures)
	ma := make(map[int64][]float64)
	max := create_max(temperatures[0])
	var values []float64
	for i := range temperatures {
		if int64(temperatures[i]) >= max {
			ma[max] = values
			max = create_max(temperatures[i])
			values = make([]float64, 0)
		}
		values = append(values, float64(temperatures[i]))
	}
	if len(values) != 0 {
		ma[max] = values
	}
	fmt.Println(ma)
}

func create_max(temperature float64) int64 {
	temperature_int := int64(temperature / 10.0)
	if temperature_int >= 0 {
		return (temperature_int + 1) * 10
	} else {
		return temperature_int * 10
	}
}
