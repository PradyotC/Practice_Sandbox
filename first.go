package main

import "fmt"

type Shape interface {
  Area() int
}

type Rectangle struct {
  length int
  breadth int
}

type Square struct {
  length int
}

func (r *Rectangle) Area() int{
  return r.length * r.breadth
}

func (s *Square) Area() int{
  return s.length * s.length
}

func main(){
  sq1 := Square{}
  rec1 := Rectangle{}
  sq1.length = 5
  rec1.length = 4
  rec1.breadth = 6
  fmt.Println(sq1.Area())
  fmt.Println(rec1.Area())
}
