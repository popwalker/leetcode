import java.util.PriorityQueue;

/*
 * @lc app=leetcode id=703 lang=java
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start
class KthLargest {
    /* 
    流式数据中第k大元素解法:
    以小顶堆来实现，java中就是优先级队列
    维护一个size = k的小顶堆
    每个新元素，要先和堆顶元素进行比较
    如果小于堆顶元素，则不处理
    如果大于堆顶元素，则将堆顶元素剔除，并将当前元素插入小顶堆
    返回小顶堆堆顶元素，即为第K大元素
    */
    final PriorityQueue<Integer> q;
    final int k;
    public KthLargest(int k, int[] nums) {
        this.k = k;
        this.q = new PriorityQueue<>(k);

        for (int n: nums){
            add(n);
        }
    }
    
    public int add(int n) {
        if (q.size() < this.k){
            q.offer(n);
        } else if (q.peek() < n){
            q.poll();
            q.offer(n);
        }
        return q.peek();
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest obj = new KthLargest(k, nums);
 * int param_1 = obj.add(val);
 */
// @lc code=end

