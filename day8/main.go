package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
  inputString := getInput()

  arr := []Point{}
  pairArr := []PairPoint{}

  lines := strings.Split(inputString, "\n")
  for _, line := range lines{
    vct := strings.Split(line, ",") 
    i, _ := strconv.Atoi(vct[0])
    j, _ := strconv.Atoi(vct[1])
    k, _ := strconv.Atoi(vct[2])
    arr = append(arr, Point{i,j,k})
  }

  n := len(arr)

  for i := range (n-1){
    for j := i+1; j < n;j++{

      a :=  math.Abs(float64(arr[i].i - arr[j].i))
      b :=  math.Abs(float64(arr[i].j - arr[j].j))
      c :=  math.Abs(float64(arr[i].k - arr[j].k))

      dist := int(a*a + b*b + c*c)
      pairArr = append(pairArr, PairPoint{dist, i ,j})

    }
  }
  
  sort.Slice(pairArr , func (i int,j int) bool  {
    a := pairArr[i] 
    b := pairArr[j] 
    return a.dist < b.dist
  })
  
  puzzle1(arr,pairArr,1000,true)
  puzzle2(arr,pairArr)
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
  k int
}

type PairPoint struct{
  dist int
  first int
  second int
}

func puzzle1(arr []Point ,pairArr []PairPoint, t int,print bool) bool {

  n := len(arr)

  graph := make([][]int, n)

  if(t >= len(pairArr)){
    return false
  }

  for i := range t{
    graph[pairArr[i].first] = append(graph[pairArr[i].first], pairArr[i].second) 
    graph[pairArr[i].second] = append(graph[pairArr[i].second], pairArr[i].first) 
  }

  vis := make([]bool, n)

  res := []int{}

  for i := range n{
    res = append(res, dfs(graph,i,vis))
  }


  sort.Slice(res, func(i, j int) bool {
    return res[i] > res[j]
  })

  if(print){
    fmt.Println(res[0] * res[1] * res[2])
  }
  return res[1] == 0

}

func dfs(graph [][]int,i int, vis []bool) int {
  if(i < 0 || i >= len(graph) || vis[i]){
    return 0
  }

  vis[i] = true
  res := 1
  for _, neigh := range graph[i]{
    res += dfs(graph, neigh, vis)
  }


  return res
}


func puzzle2(arr []Point, pairArr []PairPoint){

  low := 0 
  high := 10000

  for low <= high{

    mid := (low + high) / 2
    res := puzzle1(arr,pairArr,mid,false)

    if(res){
      high = mid - 1
    }else{
      low = mid + 1
    }


  }

  point := pairArr[low-1]

  fmt.Println(arr[point.first].i * arr[point.second].i)

}

