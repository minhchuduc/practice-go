package main
import "fmt"

func main() {
    a := &float64(5.23435) // why can not take address of float64 value???
    fmt.Println(a)
}
