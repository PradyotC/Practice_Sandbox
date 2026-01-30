package main

import "fmt"

func main(){
  
  //variable declaration
  var colorArray []string
  
  countryCodeMap := make(map[string]string)
  directionSet := make(map[string]struct{})
  var matrix [][]int
  
  //Adding Data to variable
  colorArray = append(colorArray,"red")
  colorArray = append(colorArray,"blue")
  colorArray = append(colorArray, "green")
  
  for idx, val := range(colorArray) {
    fmt.Printf("Color is %s and Index is %d\n", val, idx)
  }
  
  for i:=len(colorArray)-1; i>=0; i-- {
    fmt.Printf("Color is %s and Index is %d\n", colorArray[i], i)
  }
  
  for _, r := range(colorArray[0]) {
    fmt.Printf("Actual rune is %c but manipulated rune is %c\n", r, r+1)
  }
  
  countryCodeMap["USA"] = "us"
  countryCodeMap["China"] = "cn"
  countryCodeMap["Japan"] = "jp"
  countryCodeMap["Myanmar"] = "bu"
  
  for k, v := range countryCodeMap {
    if k=="Myanmar" {
      fmt.Printf("Old Code of Myanmar is %s\n", v)
      countryCodeMap["Myanmar"] = "mm"
      fmt.Printf("New Code of Myanmar is %s\n", "mm")
    }
    fmt.Printf("Country Code of %s is %s\n",k,v)
  }
  
  _,isPresent := countryCodeMap["USA"]
  if isPresent{
    fmt.Println("Country Code of USA is Present")
  } else {
    fmt.Println("Country Code of USA is Absent")
  }
  delete(countryCodeMap,"China")
  
  directionSet["North"] = struct{}{}
  directionSet["South"] = struct{}{}
  directionSet["East"] = struct{}{}
  directionSet["West"] = struct{}{}
  directionSet["North"] = struct{}{}
  
  for k,_ := range directionSet {
    fmt.Printf("Direction is %s\n",k)
  }
  
  count:=0
  for i:=0; i<4; i++ {
    matrix = append(matrix, []int{})
    for j:=0; j<3; j++ {
      matrix[i] = append(matrix[i], count)
      count++
    }
  }
  
  fmt.Println(matrix)
  
  for i:=0; i<4; i++ {
    for j:=0; j<3; j++ {
      fmt.Println(matrix[i][j])
    }
  }
  
}
