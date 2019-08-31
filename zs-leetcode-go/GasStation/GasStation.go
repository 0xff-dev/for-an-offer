package GasStation

// 能走完一圈肯定是钱是正数, 都傻了。。。。
func canCompleteCircuit(gas []int, cost []int) int {
	// road := len(gas)
	// for index := 0; index < road; index++ {
	// 	if cost[index] > gas[index] {
	// 		continue
	// 	}
	// 	start, allGas := (index+1)%road, gas[index]
	// 	for ; start != index; start = (start + 1) % road {
	// 		if allGas < cost[(start-1+road)%road] {
	// 			break
	// 		}
	// 		allGas += (-cost[(start-1+road)%road] + gas[start])
	// 	}
	// 	if start == index && allGas >= cost[(index-1+road)%road] {
	// 		return index
	// 	}
	// }
	// return -1
	start, total, tank := 0, 0, 0
	for index := 0; index < len(gas); index++ {
		tank += (gas[index] - cost[index])
		if tank < 0 {
			start = index + 1
			total += tank
			tank = 0
		}
	}
	if tank+total < 0 {
		return -1
	}
	return start
}
