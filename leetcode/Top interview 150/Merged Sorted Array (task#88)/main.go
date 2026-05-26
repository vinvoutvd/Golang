func merge(nums1 []int, m int, nums2 []int, n int){
	i := m-1
	j := n-1
	index := n+m-1
	for i>=0 && j>=0{
		a := nums1[i]
		b := nums2[j]
		switch{
		case a>=b:
			nums1[index]=a
			i--
		case a<b:
			nums1[index]=b
			j--
		}
		index--
	}
	for j>=0{
		nums1[index]=nums2[j]
		j--
		index--
	}
}
