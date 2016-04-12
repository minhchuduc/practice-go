package main
import (
    "fmt"
    "runtime"
    "time"
)


func main() {
    fmt.Print("Go run on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OSX")
    case "linux2":
        fmt.Println("Linux2")
    default:
        fmt.Printf("%s.\n", os)
    }

    fmt.Println("When's Wednesday?")
    today := time.Now().Weekday()
    switch time.Wednesday {  // switch-case process from top to bottom, stop when success
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
    fmt.Printf("Type of weekday is %T \n", time.Wednesday) // so WEIRD

    t := time.Now()
    switch {   // switch without a condition ~ "Long if/else chains" (cleaner way)
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }

    do(21)
    do("trường")
    do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
