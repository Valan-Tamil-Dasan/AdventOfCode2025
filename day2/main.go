package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
   puzzle1() 
   puzzle2() 
}

func getInput() string{
   raw, err := os.ReadFile("input.txt") 

   if(err != nil){
     return "";
   }

   data := string(raw)
   return data
}

type Range struct{
  low int 
  high int
}

func puzzle1(){
  inputString := getInput()
  arr := strings.Split(inputString, ",")
  ans := 0

  rangeArray := make([]Range,0) 

  for _,s := range arr{
    s = strings.TrimSpace(s)
    pair := strings.Split(s,"-") 
    _low, _ := strconv.Atoi(pair[0]) 
    _high, _ := strconv.Atoi(pair[1]) 

    val := Range{
      low : _low,
      high : _high,
    }
    
    rangeArray = append(rangeArray, val)
  }
  
  for i := range 1000000{
    mod := strconv.Itoa(i) + strconv.Itoa(i)
    modInt,_ := strconv.Atoi(mod)
    for _, val := range rangeArray{
      if(modInt <= val.high && modInt >= val.low){
        ans += modInt
      }
    }
  }

  fmt.Println(ans)

}

type StringSet map[int]struct{}

func puzzle2(){
  inputString := getInput()
  arr := strings.Split(inputString, ",")
  ans := 0

  rangeArray := make([]Range,0) 
  set := make(StringSet)

  for _,s := range arr{
    s = strings.TrimSpace(s)
    pair := strings.Split(s,"-") 
    _low, _ := strconv.Atoi(pair[0]) 
    _high, _ := strconv.Atoi(pair[1]) 

    val := Range{
      low : _low,
      high : _high,
    }
    
    rangeArray = append(rangeArray, val)
  }
  
  for i := range 1000000{
    temp := strconv.Itoa(i)
    mod := temp
    for j := 1;j<11;j++{
      mod += temp
      modInt,err := strconv.Atoi(mod)
      if err != nil{
        break
      }
      if(modInt >= 100000000000){
        break
      }

      set[modInt] = struct{}{}
    }


  }


  for key := range set{
    for _, val := range rangeArray{
      if(key <= val.high && key >= val.low){
        ans += key
      }
    }
  }

  fmt.Println(ans)

}
