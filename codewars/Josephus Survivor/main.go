//https://www.codewars.com/kata/555624b601231dc7a400017a/train/go
package main

import "fmt"

func JosephusSurvivor(n, k int) int {
	circle := make([]int, n)
	for i := range(circle){
		circle[i] = i+1
	}
	i := -1
	j := 0
	for len(circle)!=1{
		i++
		j++
		i%=len(circle)
		if j == k{
			j=0
			circle = append(circle[:i], circle[i+1:]...)
			i--
		}
	}
	return circle[0]
}
