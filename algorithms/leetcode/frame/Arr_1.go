package main

// NumArray 303. 区域和检索 - 数组不可变
type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{
			[]int{},
		}
	}

	preSum := make([]int, len(nums)+1)

	for i := 1; i < len(nums)+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	return NumArray{
		preSum: preSum,
	}
}

func (numArray *NumArray) SumRange(left int, right int) int {
	return numArray.preSum[right+1] - numArray.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
