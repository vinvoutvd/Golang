//https://leetcode.com/problems/isomorphic-strings/?envType=study-plan-v2&envId=top-interview-150

// решил сам максимально неэффективно
package main

import (
	"slices"
	"strings"
)
//создал два сета последовательных неповторяющихся элементов
//потом объеденил их в map и заменил строку s через map
func isIsomorphic(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
    setS := []string{}
    setT := []string{}
    for index := range(s){
        if !slices.Contains(setS, string(s[index])){
            setS = append(setS, string(s[index]))
        }
        if !slices.Contains(setT, string(t[index])){
            setT = append(setT, string(t[index]))
        }
    }
    if len(setS)!=len(setT){
        return false
    }
    pairs := []string{}
    for index := range(setS){
        pairs = append(pairs, setS[index], setT[index])
    }
    replacer := strings.NewReplacer(pairs...)
	s = replacer.Replace(s)
    return s==t
}
