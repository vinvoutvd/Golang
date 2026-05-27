/*
проверка есть ли в связном списке цикл:
Создаются два указателя: один идёт на один шаг впереди другого, есть ли вдруг эти указатели равны то цикл есть, нет в обратном случае
*/
func hasCycle(head *ListNode) bool {
	a := head
	b := head
	for b!=nil && b.Next!=nil{
		a = a.Next
		b = b.Next.Next
		if a == b{
			return true
		}
	}
    return false
}
