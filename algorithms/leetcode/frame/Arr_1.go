package main

// 303. 区域和检索 - 数组不可变
type NumArray struct {
	preSum  []int
	numsLen int
}

func Constructor(nums []int) NumArray {
	// 没有元素，直接返回
	if len(nums) == 0 {
		return NumArray{
			preSum:  []int{},
			numsLen: 0,
		}
	}
	// 构建前缀和
	nA := NumArray{
		preSum:  make([]int, len(nums)+1),
		numsLen: len(nums) + 1,
	}
	for i := 1; i < len(nums)+1; i++ {
		nA.preSum[i] = nA.preSum[i-1] + nums[i-1]
	}
	return nA
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
