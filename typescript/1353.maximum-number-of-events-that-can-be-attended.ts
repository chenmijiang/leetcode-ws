/*
 * @lc app=leetcode.cn id=1353 lang=typescript
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

// @lcpr-template-end
// @lc code=start
function maxEvents(events: number[][]): number {
  const n = events.length;
  // 排序
  events.sort((a, b) => a[0] - b[0]);
  const minHeap = new MinHeap();
  let ans = 0,
    i = 0,
    d = events[0][0];

  while (i < n || minHeap.size() > 0) {
    while (i < n && events[i][0] === d) {
      minHeap.insert(events[i][1]);
      i++;
    }

    // 移除过期会议
    while (minHeap.size() > 0 && minHeap.peek() < d) {
      minHeap.extract();
    }

    // 参加一个结束时间最早的活动
    if (minHeap.size() > 0) {
      minHeap.extract();
      ans++;
    }

    // 更新天数
    if (minHeap.size() === 0 && i < n) {
      d = events[i][0];
    } else {
      d++;
    }
  }

  return ans;
}

// 最小堆
class MinHeap {
  private heap: number[];
  constructor() {
    this.heap = [];
  }

  size(): number {
    return this.heap.length;
  }

  peek(): number {
    return this.heap[0];
  }

  insert(val: number): void {
    this.heap.push(val);
    this.siftUp(this.heap.length - 1);
  }

  extract(): number | undefined {
    if (this.heap.length === 0) return undefined;
    
    const min = this.heap[0];
    const last = this.heap.pop()!;
    
    if (this.heap.length > 0) {
      this.heap[0] = last;
      this.siftDown(0);
    }
    
    return min;
  }

  private siftUp(index: number): void {
    let parent = Math.floor((index - 1) / 2);
    
    while (index > 0 && this.heap[parent] > this.heap[index]) {
      [this.heap[parent], this.heap[index]] = [this.heap[index], this.heap[parent]];
      index = parent;
      parent = Math.floor((index - 1) / 2);
    }
  }

  private siftDown(index: number): void {
    let smallest = index;
    const left = 2 * index + 1;
    const right = 2 * index + 2;
    const length = this.heap.length;
    
    if (left < length && this.heap[left] < this.heap[smallest]) {
      smallest = left;
    }
    
    if (right < length && this.heap[right] < this.heap[smallest]) {
      smallest = right;
    }
    
    if (smallest !== index) {
      [this.heap[index], this.heap[smallest]] = [this.heap[smallest], this.heap[index]];
      this.siftDown(smallest);
    }
  }
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2],[2,3],[3,4]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[2,3],[3,4],[1,2]]\n
// @lcpr case=end

 */
