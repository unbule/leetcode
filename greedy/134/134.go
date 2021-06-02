/**
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	gas := []int{5, 1, 2, 3, 4}
	cost := []int{4, 4, 1, 5, 1}
	res := canCompleteCircuit(gas, cost)
	fmt.Println(res)

}

//暴力破解
func canCompleteCircuit(gas []int, cost []int) int {
	SumGas := 0
	j := 0
	for i := 0; i < len(gas); i++ {
		beginGas := chooseBegin(i, gas, cost)
		if beginGas == -1 {
			continue
		}
		SumGas = 0
		SumGas = SumGas + gas[beginGas] - cost[beginGas]
		//fmt.Println("for wai",SumGas,i)
		for j = beginGas + 1; j%len(gas) != beginGas; j++ {
			SumGas = SumGas + gas[j%len(gas)] - cost[j%len(gas)]
			//fmt.Println(SumGas)
			if SumGas < 0 {
				break
			}
		}
		if j%len(gas) == beginGas {
			return j % len(gas)
		}
	}
	return -1
}

func chooseBegin(i int, gas, cost []int) int {
	if gas[i] >= cost[i] {
		return i
	}
	return -1
}

//图表法
func canCompleteCircuit2(gas []int, cost []int) int {
	min := math.MaxInt8
	gasIndex := 0
	//costIndex := 0
	sumd := 0
	for i := 0; i < len(gas); i++ {
		sumd = sumd + gas[i] - cost[i]
		if sumd < min {
			min = sumd
			gasIndex = i
		}
	}
	if sumd >= 0 {
		return (gasIndex + 1) % len(gas)
	}
	return -1
}
