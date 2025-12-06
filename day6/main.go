package main

import (
	"fmt"
	"os"
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

func pow(a, b int) int {
  res := 1
  for i := 0; i < b; i++ {
      res *= a
  }
  return res
}


func puzzle1(inputString string){

  if len(inputString) == 0 {
    return
  }
  

  ans := 0
  
  lines := strings.Split(inputString, "\n")
  grid := make([][]int, len(lines))

  n := len(grid)

  for i, line := range lines{
    fields := strings.Fields(line)
    grid[i] = make([]int, len(fields))

    if(i == n-1){
      for j, f := range fields{
        if(f == "*"){
          grid[i][j] = 1
        }else{
          grid[i][j] = 0
        }
      }
      continue
    }

    for j, f := range fields {
        num, _ := strconv.Atoi(f)
        grid[i][j] = num
    }
  }

  m := len(grid[0])

  for j := range m{
    ans += calcRow(grid, j)
  }


  fmt.Println(ans)

}

func calcRow(grid [][]int, j int) int {
  n := len(grid)
  if(grid[n-1][j] == 1){
    return mulCol(grid, j);
  } 

  return addCol(grid, j);
}

func mulCol(grid [][]int, j int) int {
  ans := 1 
  n := len(grid)

  for i := range n{
    ans *= grid[i][j]
  } 

  return ans
}

func addCol(grid [][]int, j int) int {
  ans := 0 
  n := len(grid)
  for i := range n{
    ans += grid[i][j]
  } 

  return ans
}

func puzzle2(inputString string){
  if len(inputString) == 0 {
    return
  }

  ans := 0
  lines := strings.Split(inputString, "\n")
  
  n := len(lines) 
  m := len(lines[0]) 

  ops := strings.Fields( lines[n-1] )
  opPointer := len(ops) - 1

  var res int

  res = setRes(ops[opPointer])

  for j := m-1;j>=-1;j--{

    if(j < 0){
      ans += res
      continue
    }

    if isColBlank(lines, j){
      opPointer--;
      ans += res
      res = setRes(ops[opPointer])
      continue
    }

    cur := getNum(lines, j)
    if(ops[opPointer] == "*"){
      res *= cur
    }else{
      res += cur
    }

  }

  fmt.Println(ans)


}

func isColBlank(lines []string, j int)bool{
  n := len(lines)  
  for i := range n-1{
    if(lines[i][j] != ' '){
      return false
    }
  }
  return true
}

func getNum(lines []string, j int) int {
  n := len(lines)  
  cur := 0

  for i := range n-1{
    if(lines[i][j] != ' '){
      cur *= 10
      cur += int(lines[i][j]) - 48
    }
  }

  return cur
}

func setRes(op string) int{
  if(op == "*"){
    return 1
  }
  return 0
}
