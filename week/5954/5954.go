package main

import "fmt"

func main() {
	plants := []int{2, 2, 3, 3}
	fmt.Println(minimumRefill(plants, 5, 5))
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	counta, countb := 0, 0
	cura, curb := capacityA, capacityB
	plantb := make([]int, len(plants))
	for j := len(plants) - 1; j >= 0; j-- {
		plantb[len(plants)-1-j] = plants[j]
	}
	for i := 0; i < len(plants)/2; i++ {
		if cura < plants[i] {
			cura = capacityA
			counta++
		}
		cura -= plants[i]
	}
	for j := 0; j < len(plantb)/2; j++ {
		if curb < plantb[j] {
			curb = capacityB
			countb++
		}
		curb -= plantb[j]
	}
	if len(plants)%2 == 0 {
		return counta + countb
	} else {
		if cura >= curb {
			if cura < plants[len(plants)/2] {
				return counta + countb + 1
			} else {
				return counta + countb
			}
		} else {
			if curb < plants[len(plants)/2] {
				return counta + countb + 1
			} else {
				return counta + countb
			}
		}
	}
}
