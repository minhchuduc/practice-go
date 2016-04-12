package main

import "fmt"

func produce(c chan int) {
     for i := 0; i < 10; i++ {
          c <- i
     }
     close(c)
}

func consume(c chan int, done chan bool) {
     for i := range c {
          fmt.Println(i)
     }
     close(done)
}

func main() {
     c := make(chan int)
     done := make(chan bool)
     go produce(c)
     go consume(c, done)
     <- done
}
