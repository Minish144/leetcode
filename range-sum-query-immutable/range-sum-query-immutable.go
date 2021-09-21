type NumArray struct {
    array []int
}


func Constructor(nums []int) NumArray {
    return NumArray{
        array: nums,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    sum := 0
    for i := left; i <= right; i++ {
        sum += this.array[i]
    }
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */