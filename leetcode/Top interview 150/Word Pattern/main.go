//https://leetcode.com/problems/word-pattern/?envType=study-plan-v2&envId=top-interview-150

import "strings"

func wordPattern(pattern string, s string) bool {
  // создание списка из строки без пробелов
	arrS := strings.Split(s, " ")
	if len(arrS)!=len(pattern){
		return false}
  //map pattern -> s для каждого символа из pattern проверяет уникальность элемента в arrS 
	PatternToS := map[string]string{}
  //map s -> pattern для каждого элемента из arrS проверяет уникальность элемента в pattern
	SToPattern := map[string]string{}
  //сама проверка
	for i := range(pattern){
		key := string(pattern[i])
		element, ok := PatternToS[key]
		if !ok{
			PatternToS[key] = arrS[i]
		}else if element!=arrS[i]{
			return false
		}
		key = arrS[i]
		element, ok = SToPattern[key]
		if !ok{
			SToPattern[key] = string(pattern[i])
		}else if element!=string(pattern[i]){
			return false
		}
	}
	return true
}
