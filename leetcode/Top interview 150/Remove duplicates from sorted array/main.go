func removeDuplicates(nums []int) int {
	i, counter, swap := 0, 1, 1
	for swap < len(nums){
		if nums[swap]>nums[i]{
			nums[i+1] = nums[swap]
			i++
			counter+=1
		}
		swap++
	}
    return counter
}
