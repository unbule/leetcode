package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	mapp := make(map[string]int)
	n := len(nums)
	sort.Ints(nums)
	//fmt.Println(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			slic := nums[j+1:]
			index := sort.SearchInts(slic, 0-(nums[i]+nums[j]))
			if index >= len(slic) {
				continue
			}
			if (nums[i]+nums[j])+slic[index] == 0 {
				cn := make([]int, 3)
				//fmt.Println("numssss")
				cn[0] = nums[i]
				cn[1] = nums[j]
				cn[2] = slic[index]
				res = append(res, cn)
			}
		}
	}
	ans := [][]int{}
	str := ""
	for x := 0; x < len(res); x++ {
		str = ""
		for _, num := range res[x] {
			str += strconv.Itoa(num)
		}
		if mapp[str] == 1 {
			continue
		}
		ans = append(ans, res[x])
		mapp[str] = 1
	}
	return ans
}

func threeSum1(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		c := n - 1
		target := -1 * nums[a]
		for b := a + 1; b < n; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for b < c && nums[b]+nums[c] > target {
				c--
			}

			if b == c {
				break
			}

			if nums[b]+nums[c] == target {
				res = append(res, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return res
}
