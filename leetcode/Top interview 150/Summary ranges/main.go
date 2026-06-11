//https://leetcode.com/problems/summary-ranges/description/?envType=study-plan-v2&envId=top-interview-150

func summaryRanges(nums []int) (ranges []string) {
  //проверка на наличие элементов
	if len(nums)==0{
		return ranges
	}
  //подсчет поочередных чисел
	counter:=0
  //перебор от 0 до N-1
    for i:=0; i<len(nums)-1; i++{
    // если числа отличаются на 1 увеличиваем counter. Шагаем дальше
		if nums[i]+1==nums[i+1]{
			counter+=1
		}else if counter!=0{ // иначе если counter > 0, то записываем полученный range
			ranges = append(ranges, fmt.Sprintf("%d->%d", nums[i-counter], nums[i]))
			counter=0
		}else { //если counter == 0 то есть nums[i]!=nums[i+1], добавляем один жлемент.
			ranges = append(ranges, fmt.Sprint(nums[i]))
			counter=0}
	}
  // после цикла for тоже последний раз вставляем элементы
	if counter==0{
		ranges = append(ranges, fmt.Sprint(nums[len(nums)-1]))
	}else{
		ranges = append(ranges, fmt.Sprintf("%d->%d", nums[len(nums)-counter-1], nums[len(nums)-1]))
	}
	return ranges
}
