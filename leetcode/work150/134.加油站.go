package work150

/*
1. 总油量 vs 总消耗
  如果 sum(gas) < sum(cost)，无解，直接返回 -1
    因为总油不够消耗一圈
2. 从某个点出发油量不够时怎么办？
   设 tank 表示当前油量
   遍历每个加油站：
    tank += gas[i] - cost[i]
    如果 tank < 0，说明从起点到这里油不够
  		那么 不能从之前的起点出发
		下一站作为新的起点，重置 tank = 0
3. 为什么这样贪心能保证唯一解？
	因为如果总油量够，但从某个点油量不足，则这个点之前的任何起点都不可能成功
	所以最后记录的起点就是唯一解
*/

func canCompleteCircuit(gas []int, cost []int) int {
	totalTank, tank := 0, 0
	n := len(gas)
	start := 0
	for i := range n {
		totalTank += gas[i] - cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}
	if totalTank < 0 {
		return -1
	}
	return start
}
