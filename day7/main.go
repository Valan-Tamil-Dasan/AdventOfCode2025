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
     fmt.Println("Error reading file")
     return "";
   }
   data := strings.TrimSpace(string(raw))
   return data
}

type Point struct{
  i int
  j int
}

func puzzle1(inputString string){

  if len(inputString) == 0 {
    return
  }

  lines := strings.Fields(inputString)
  
  n:=len(lines)
  m:=len(lines[0])

  vis := make([][]bool,n)

  for i := range n{
    vis[i] = make([]bool,m)
  }
  
  q := make([]Point,0)

  for i := range n{
    for j := range m{
      if lines[i][j] == 'S'{
        q = append(q, Point{i,j})
        break
      }
    }
  }

  ans := 0

  for(len(q) != 0){
    node := q[0]
    q = q[1:]
  
    if(node.j < 0 || node.j >= m || node.i < 0 || node.i >= n) {
      continue
    }

    if(vis[node.i][node.j]){
      continue
    }

    
    if(lines[node.i][node.j] == '^'){
      if(vis[node.i][node.j] == false){
        ans++;
        q = append(q, Point{node.i,node.j-1})
        q = append(q, Point{node.i,node.j+1})
      }
    }else{
        q = append(q, Point{node.i+1,node.j})
    }

    vis[node.i][node.j] = true

  }

  fmt.Println(ans)
  
}

func puzzle2(inputString string){
  if len(inputString) == 0 {
    return
  }

  lines := strings.Fields(inputString)
  
  n:=len(lines)
  m:=len(lines[0])

  dp := make([][]int,n)

  for i := range n{
    dp[i] = make([]int,m)
  }
  
  for i := range n{
    for j := range m{
      if lines[i][j] == 'S'{
        dp[i][j] = 1
        break
      }
    }
  }

  for i := range n{
    for j := range m{
      if lines[i][j] == '^'{
        dp[i][j] = 0
      }else{
        if(i != 0){
          dp[i][j] = dp[i-1][j]
          if(j < m-1 && lines[i][j+1] == '^'){
            dp[i][j] += dp[i-1][j+1] 
          }
          if(j > 0 && lines[i][j-1] == '^'){
            dp[i][j] += dp[i-1][j-1] 
          }
        }
      }
    }
  }

  ans := 0
  for j := range m{
    ans += dp[n-1][j] 
  }
  fmt.Println(ans)

}

