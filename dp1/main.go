package main

// Есть и другое решение, без массива )
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}

func minCostClimbingStairs(cost []int) int {
	// 746. Min Cost Climbing Stairs
	// Переделать самостоятельно!!!

	a, b := 0, 0
	for i := 1; i < len(cost); i++ {
		a, b = b, min(a+cost[i-1], b+cost[i])
	}

	return b
}

func main() {
	tribonacci(7)

	minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
}
