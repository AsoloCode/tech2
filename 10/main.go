package main

import (
	"fmt"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temperatureGroups := make(map[int][]float64)

	for _, temp := range temperatures {
		groupKey := int(temp/10) * 10
		temperatureGroups[groupKey] = append(temperatureGroups[groupKey], temp)
	}

	for key, group := range temperatureGroups {
		fmt.Printf("%d: %v\n", key, group)
	}
}
