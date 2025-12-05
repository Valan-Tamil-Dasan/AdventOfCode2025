package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main(){
  inputString := getInput()
  puzzle1(inputString)
  puzzle2(inputString)
}

func getInput() string{
   raw, err := os.ReadFile("input.txt") 
   if(err != nil){
     fmt.Println("Error reading file")
     return "";
   }
   data := strings.TrimSpace(string(raw))
   return data
}


type FreshRange struct{
  low int
  high int
}

func isFresh(id int,rangesArr []FreshRange)bool{
  for _, freshRange := range rangesArr{
    if(id >= freshRange.low && id <= freshRange.high){
      return true
    }
  }
  return false 
}


func puzzle1(inputString string){

  if len(inputString) == 0 {
    return
  }

  input := strings.Split(inputString, "\n\n") 

  if len(input) < 2 {
    return
  }

  rangesRaw := strings.Split(input[0] , "\n")
  queriesRaw := strings.Split(input[1], "\n")
  
  rangesArr := []FreshRange{}
  queriesArr := []int{}
  
  for _,val := range rangesRaw{
    sl := strings.Split(val,"-")
    low, _ := strconv.Atoi(sl[0])
    high, _ := strconv.Atoi(sl[1])
    rangesArr = append(rangesArr, FreshRange{low,high})
  }

  for _, val := range queriesRaw{
    q, _ := strconv.Atoi(val)
    queriesArr = append(queriesArr, q)
  } 

  ans:= 0

  for _, val := range queriesArr{
    if(isFresh(val,rangesArr)){
      ans++
    }
  }
  
  fmt.Println(ans)

}


func puzzle2(inputString string){
  if len(inputString) == 0 {
    return
  }

  input := strings.Split(inputString, "\n\n") 

  if len(input) < 2 {
    return
  }

  rangesRaw := strings.Split(input[0] , "\n")
  rangesArr := []FreshRange{}
  uniqueRanges := []FreshRange{}
  
  for _,val := range rangesRaw{
    sl := strings.Split(val,"-")
    low, _ := strconv.Atoi(sl[0])
    high, _ := strconv.Atoi(sl[1])
    rangesArr = append(rangesArr, FreshRange{min(low,high),max(low,high)})
  }

  sort.Slice(rangesArr , func(i int, j int) bool{
    return rangesArr[i].low < rangesArr[j].low
  })

  for _, val := range rangesArr{

    if(len(uniqueRanges) == 0){
      uniqueRanges = append(uniqueRanges, val)
      continue
    }
    
    if(val.low <= uniqueRanges[len(uniqueRanges) - 1].high){
      uniqueRanges[len(uniqueRanges) - 1].high = max(val.high , uniqueRanges[len(uniqueRanges) - 1].high)
    }else{
      uniqueRanges = append(uniqueRanges, val)
    }

  } 

  ans := 0

  // fmt.Println(uniqueRanges)
  // fmt.Println(rangesArr)
  for _, val := range uniqueRanges{
    // fmt.Println(val.high - val.low + 1)    
    ans += val.high - val.low + 1
  }

  fmt.Println(ans)

}
