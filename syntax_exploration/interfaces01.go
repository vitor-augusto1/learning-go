package main

import "fmt"


// Convention `-er` suffix name of the interface
// It only describes the expected behavior (methods)
type Printer interface {
  print()
}

type Game struct {
  name       string
  mutiplayer bool
  price      int
}

type Book struct {
  name  string
  price int
}

func (g *Game) print() {
  fmt.Println(g.name)
  fmt.Println(g.price)
  fmt.Println(g.mutiplayer)
}

func (b *Book) print() {
  fmt.Println(b.name)
  fmt.Println(b.price)
}

type list []Printer

func (l list) print() {
  if len(l) == 0 {
    panic("List has no vlaues.")
  }

  for _, it := range l {
    it.print()
  }
}

func main() {
  var (
    mobydick = Book{name: "Moby dick", price: 20}
    minecraft = Game{name: "Minecraft", mutiplayer: true, price: 73}
    mario = Game{name: "Mario", mutiplayer: false, price: 25}
  )
  var store list


  store = append(store, &mobydick, &minecraft, &mario)
  store.print()
}
