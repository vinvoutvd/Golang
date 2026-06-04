//решено через алгоритм "cкользящее окно" (sliding window)
// на каждом элементе смотрим меньше ли оно предыдущего если да то обновляем минимум,
// если нет то проверяем больше ли разница между элементом и минимумом текущего profit, если да то обновляем profit
func maxProfit(prices []int) int {
    mn := prices[0]
	  profit := 0
    for i:=1; i<len(prices); i++{
		switch{
		case prices[i] < mn: mn = prices[i]
		case prices[i]-mn > profit: profit = prices[i]-mn
		}
    }
    return profit
}
