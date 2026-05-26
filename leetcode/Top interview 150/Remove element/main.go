func removeElement(nums []int, val int) int {
    begin, end := 0, len(nums)-1
    for begin <= end {
        if nums[begin] == val {
            nums[begin], nums[end] = nums[end], nums[begin]
            end--
        } else {
            begin++
        }
    }
    return begin
}
