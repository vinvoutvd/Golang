/*
сдвиг односвязного списка вправо на k элементов:
k=2 [1->2->3->nil]
3->1->2->nil
3->2->1->nil - result linked list
*/
func rotateRight(head *ListNode, k int) *ListNode {
	var lastNode *ListNode
  //ссылка на начало linked list
	var headNode *ListNode = head
  // проверка на достаточность элементов и сдвига (если сдвиг равен 0 ничего сдвигать не требуется)
	if head==nil || head.Next==nil || k==0{
		return head
	}
	for k>=0{
    // переменная для подсчёта количества элементов
		ln := 2
    //проверяем  будет ли следующий после head иметь поле Next==nil
		for head.Next.Next!=nil{
			head = head.Next
			ln++
		}
    //определяем на сколько шагов в целом нужно сдвинуть list (остаток от деления) -> чтобы не считать очень долго если k будет большим
		k%=ln
    //опять проверка на 0
		if k==0{
			return headNode
		}
    // присваем на предпоследний элемент
		lastNode = head.Next
    // связываем конечный элемент с началом
		lastNode.Next = headNode
    // обнуляем предпоследний элемент
		head.Next = nil
    // ссылка на новое начало
		headNode = lastNode
    // список head - используется для итерации по linked list
		head = lastNode
		k--
	}
	return lastNode
}
