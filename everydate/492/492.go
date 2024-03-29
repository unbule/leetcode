package main

import "math"

func main() {

}

func constructRectangle(area int) []int {
	for i := int(math.Sqrt(float64(area))); i >= 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{}
}
