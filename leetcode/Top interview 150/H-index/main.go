//https://leetcode.com/problems/h-index/?envType=study-plan-v2&envId=top-interview-150
//индекс Хёрша - https://ru.wikipedia.org/wiki/H-index

import "slices"

//нужно найти максимальное число h для которого соблюдается условие: количество элементов больших или равных числу h больше или равно h
// [3,0,6,1,5] число Хёрша будет равно 3 потому что есть три числа 3,6,5 которые больше или равны трём и их количество = 3 
func hIndex(citations []int) int {
  // сортируем список, так найти количество элемнтов больших h будет гораздо быстрее [i:]
    slices.Sort(citations)
    H := 0
	for h:=0; h <= len(citations); h++{
		for i := range(citations){
      // проверка на 0 и само условие 
			if citations[i]>=h && citations[i]!=0 && len(citations[i:])>=h{
				H = h
        break
			}
		}
	}
	return H
}
