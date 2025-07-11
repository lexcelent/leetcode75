package main

import "fmt"

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{queue: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	fmt.Println(this.queue)

	left, right := t-3000, t

	var count int
	for i := len(this.queue) - 1; i >= 0; i-- {
		if this.queue[i] >= left && this.queue[i] <= right {
			count++
		} else {
			break
		}
	}

	return count
}

func main() {
	recentCounter := Constructor()
	recentCounter.Ping(1)    // requests = [1], range is [-2999,1], return 1
	recentCounter.Ping(100)  // requests = [1, 100], range is [-2900,100], return 2
	recentCounter.Ping(3001) // requests = [1, 100, 3001], range is [1,3001], return 3
	recentCounter.Ping(3002) // requests = [1, 100, 3001, 3002], range is [2,3002], return 3
}
