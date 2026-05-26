//https://www.codewars.com/kata/5902bc7aba39542b4a00003d/train/go

package main

import (
	"fmt"
	"strings"
)

func WhoEatsWho(zoo string) []string {
  animals := strings.Split(zoo, ",")
  output := make([]string, 0)
  output = append(output, zoo)
  eating_list := map[string][]string{
    "antelope":{"grass"},
    "big-fish":{"little-fish"},
    "bug":{"leaves"},
    "bear":{"big-fish", "bug", "chicken", "cow", "leaves", "sheep"},
    "chicken":{"bug"},
    "cow":{"grass"},
    "fox":{"chicken", "sheep"},
    "giraffe":{"leaves"},
    "lion":{"antelope", "lion"},
    "panda":{"leaves"},
    "sheep":{"grass"},
  }
  exit:=1
  for exit!=2{
    for i, animal := range(animals){
      switch {
      case i>0 && Contains(eating_list[animal], animals[i-1]):
        output = append(output, fmt.Sprintf("%s eats %s", animal, animals[i-1]))
        animals = append(animals[:i-1], animals[i:]...)
        exit = 0
      case i<len(animals)-1 && Contains(eating_list[animal], animals[i+1]):
        output = append(output, fmt.Sprintf("%s eats %s", animal, animals[i+1]))
        animals = append(animals[:i+1],animals[i+2:]... )
        exit=0
      }
      if exit==0{
        break
      }
    }
    exit++
  }
  output = append(output, strings.Join(animals, ","))
  return output
}

func Contains(data []string, line string) bool{
  for _, element := range(data){
    if element==line{
      return true
    }
  }
  return false
}
