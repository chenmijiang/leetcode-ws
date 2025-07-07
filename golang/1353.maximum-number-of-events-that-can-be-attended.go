/*
 * @lc app=leetcode.cn id=1353 lang=golang
 * @lcpr version=30204
 *
 * [1353] 最多可以参加的会议数目
 *
 * https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/description/
 *
 * algorithms
 * Medium (30.79%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    30.8K
 * Total Submissions: 92.2K
 * Testcase Example:  '[[1,2],[2,3],[3,4]]'
 *
 * 给你一个数组 events，其中 events[i] = [startDayi, endDayi] ，表示会议 i 开始于 startDayi ，结束于
 * endDayi 。
 *
 * 你可以在满足 startDayi <= d <= endDayi 中的任意一天 d 参加会议 i 。在任意一天 d 中只能参加一场会议。
 *
 * 请你返回你可以参加的 最大 会议数目。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：events = [[1,2],[2,3],[3,4]]
 * 输出：3
 * 解释：你可以参加所有的三个会议。
 * 安排会议的一种方案如上图。
 * 第 1 天参加第一个会议。
 * 第 2 天参加第二个会议。
 * 第 3 天参加第三个会议。
 *
 *
 * 示例 2：
 *
 * 输入：events= [[1,2],[2,3],[3,4],[1,2]]
 * 输出：4
 *
 *
 *
 *
 * 提示：​​​​​​
 *
 *
 * 1 <= events.length <= 10^5
 * events[i].length == 2
 * 1 <= startDayi <= endDayi <= 10^5
 *
 *
 */

// @lcpr-template-start
package main

import (
	"container/heap"
	"sort"
)

// @lcpr-template-end
// @lc code=start
func maxEvents(events [][]int) int {
	n := len(events)

	// 按开始时间排序
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	pq := &IntHeap{} // 最小堆，存储结束时间
	heap.Init(pq)

	count := 0
	i := 0

	// 从最早的活动开始时间开始遍历
	d := events[0][0]

	for i < n || pq.Len() > 0 {
		// 将当天可以开始的所有活动的结束时间加入最小堆
		for i < n && events[i][0] == d {
			heap.Push(pq, events[i][1])
			i++
		}

		// 移除已经结束的活动
		for pq.Len() > 0 && (*pq)[0] < d {
			heap.Pop(pq)
		}

		// 参加一个结束时间最早的活动
		if pq.Len() > 0 {
			heap.Pop(pq)
			count++
		}

		d++

		// 优化：如果堆为空，且还有活动，直接跳到下一个活动的开始时间
		if pq.Len() == 0 && i < n {
			d = events[i][0]
		}
	}

	return count
}

// 最小堆实现
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4],[4,4],[2,2],[3,4],[1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[2,3],[3,4],[1,2]]\n
// @lcpr case=end

*/
