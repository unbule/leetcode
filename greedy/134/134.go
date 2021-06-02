package main

import "fmt"

func main(){
	gas := []int{5,1,2,3,4}
	cost := []int{4,4,1,5,1}
	res := canCompleteCircuit(gas,cost)
	fmt.Println(res)

}

func canCompleteCircuit(gas []int, cost []int) int {
	SumGas := 0
	j := 0
	for i:=0; i<len(gas); i++{
		beginGas := chooseBegin(i, gas, cost)
		if beginGas == -1{
			continue
		}
		SumGas = 0
		SumGas = SumGas + gas[beginGas] - cost[beginGas]
		//fmt.Println("for wai",SumGas,i)
		for j = beginGas+1; j%len(gas) != beginGas; j++{
			SumGas = SumGas + gas[j%len(gas)] - cost[j%len(gas)]
			//fmt.Println(SumGas)
			if SumGas < 0{
				break
			}
		}
		if j%len(gas) == beginGas{
			return j%len(gas)
		}
	} 
	return -1
}

func chooseBegin(i int, gas, cost []int) int {
	if gas[i] >= cost[i]{
		return i
	}
	return -1
}