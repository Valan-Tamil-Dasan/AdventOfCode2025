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

func isValid(i int, j int,grid [][]byte) bool{
  count := 0
  row := []int {1,1,1,-1,-1,-1,0,0}
  col := []int {-1,0,1,-1,0,1,1,-1}

  for k := range 8{
    ni := i + row[k]
    nj := j + col[k]
    
    if(ni < 0 || nj < 0 || ni >= len(grid) || nj >= len(grid[0])){
      continue
    }

    if(grid[ni][nj] == '@'){
      count++
    }

  }

  return count < 4
  
}

func puzzle1(inputString string){
  arr := strings.Split(inputString,"\n")

  grid := make([][]byte,len(arr))  

  for i := range arr{
    grid[i] = []byte(arr[i])
  }

  ans := 0

  for i, row := range grid{
    for j, cell := range row{
      if(cell == '@'){
        if isValid(i,j,grid){
          ans++
        }
      } 
    }
  }

  fmt.Println(ans)

}

type Pair struct{
  i int
  j int
}

func pushAdjacent(i int,j int,grid [][]byte,q []Pair) []Pair {
  row := []int {1,1,1,-1,-1,-1,0,0}
  col := []int {-1,0,1,-1,0,1,1,-1}

  for k := range 8{
    ni := i + row[k]
    nj := j + col[k]
    
    if(ni < 0 || nj < 0 || ni >= len(grid) || nj >= len(grid[0])){
      continue
    }

    if(grid[ni][nj] == '@'){
      if(isValid(ni,nj,grid)){
        q = append(q, Pair{ni,nj})
      }
    }

  }

  return q
}

func puzzle2(inputString string){
  arr := strings.Split(inputString,"\n")
  grid := make([][]byte,len(arr))  
  ans := 0

  for i := range arr{
    grid[i] = []byte(arr[i])
  }
  
  q := []Pair{}

  for i, row := range grid{
    for j, cell := range row{
      if(cell == '@'){
        if isValid(i,j,grid) {
          q = append(q,Pair{i,j})
        }
      } 
    }
  }

  for len(q) > 0{
    node := q[0]
    q = q[1:]

    if(grid[node.i][node.j] != '@'){
      continue
    }
     
    grid[node.i][node.j] = '.'
    ans ++
    
    q = pushAdjacent(node.i,node.j,grid,q)

  }
  
  fmt.Println(ans)

}
