package main

import (
	"fmt"
	"os"
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
     return "";
   }
   data := strings.TrimSpace(string(raw))
   return data
}


func getLargest(bank string) int{
  n := len(bank)
  res := 0
  digit2 := int(bank[n-1]) - 48
  largestDigit2 := digit2

  for i := n-2 ; i>=0;i-- {
    digit1 := int(bank[i]) - 48
    curNum := digit1 * 10 + largestDigit2
    res = max(curNum,res)
    largestDigit2 = max(largestDigit2 , digit1)
  }

  return res
}

func puzzle1(inputString string){
  ans := 0 
  arr := strings.Split(inputString, "\n")

  for _, bank := range arr{
    ans += getLargest(bank)
  }

  fmt.Println(ans)
}

func pow(a, b int) int {
  res := 1
  for i := 0; i < b; i++ {
      res *= a
  }
  return res
}

func getLargestTDigit(bank string, t int) int{
  if(t < 2 || t >= len(bank)){
    return 0
  } 
  n := len(bank)

  prev := make([]int , n)
  prev[n-1] = int(bank[n-1]) - 48

  for i := n-2;i>=0;i--{
    digit := int(bank[i]) - 48
    prev[i] = max(prev[i+1] , digit)
  }
  
  for i := range t-1{
    dp := make([]int, n)
    
    temp := 0

    for j := n - i - 2;j<n;j++{
      temp *= 10
      temp += int(bank[j]) - 48
    }

    dp[n-i-2] = temp;

    for j := n-1-(i+2); j>=0;j--{
      cur := (int(bank[j]) - 48) * pow(10 , i + 1) + prev[j+1]
      dp[j] = max(cur , dp[j+1])
    } 

    // fmt.Println(prev)
    prev = dp
    // fmt.Println(dp)
  }
  return prev[0]
}

func puzzle2(inputString string){
  ans := 0
  arr := strings.Split(inputString, "\n")
  t := 12

  for _, bank := range arr{
    ans += getLargestTDigit(bank, t)
  }

  fmt.Println(ans)
  // fmt.Println(getLargestTDigit("818181911112111", 12))
}
