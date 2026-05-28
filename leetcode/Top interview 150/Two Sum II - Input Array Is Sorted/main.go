/*
ищет индексы тех элементов в списке которые дают в сумме target
(писок упорядочен по возрастанию и есть хотя бы одно решение)
*/
func twoSum(numbers []int, target int) []int {
	for i:=0; i<len(numbers)-1;i++{
		for j:=i+1; j<=len(numbers)-1;j++{
			if numbers[i]+numbers[j]==target{
				return []int{i+1, j+1}
			}
		}
	}
	return nil
}
