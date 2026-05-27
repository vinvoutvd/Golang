/*
Функция проверяет есть ли в массиве дупликаты по условию:
n[i]==n[j] and abs(i-j)<=k
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	for i:=0; i < len(nums)-1; i++{
		for j:=1; j<=k; j++{
			if i+j<len(nums) && nums[i]==nums[i+j]{
				return true
			}
		}
	}
	return false
}
