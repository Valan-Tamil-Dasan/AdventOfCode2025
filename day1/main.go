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

func puzzle2(){
  ans := 0
  cur := 50

  inputString := getInput()
  lines := strings.Fields(inputString)

  for _, s := range lines{
    dir := s[0:1]  
    numStr := s[1:]

    toAdd, _ := strconv.Atoi(numStr)

    if (dir == "L" ){
      temp := cur
      temp -= toAdd

      if(temp <= 0){
        ans += ((-1 * temp) / 100)
        if(cur != 0){
          ans++
        }
      }

      toAdd = 100 - (toAdd % 100)

    }else{
      temp := cur
      temp += toAdd 

      ans += (temp / 100)
    }

    cur += toAdd
    cur %= 100

  }
  fmt.Println(ans)
}


func puzzle1(){
  inputString := getInput()

  lines := strings.Fields(inputString)

  ans := 0
  cur := 50

    

  for _ , s := range lines {

    dir := s[0:1]
    numStr := s[1:]
    
    toAdd, _ := strconv.Atoi(numStr)

    if(dir == "L"){
      toAdd = 100 - (toAdd % 100)
    }

    cur += toAdd
    cur %= 100

    if (cur == 0){
      ans++;
    }

  }

  fmt.Println(ans)
}

func getInput() string{
  rawData, err := os.ReadFile("input.txt")

  if err != nil {
    fmt.Println("Error occured" , err)
    return ""
  }

  data := string(rawData)
  return data
}
