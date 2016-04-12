package main
import "fmt"
//import "go/types"

func f1() {
    scores := make([]int, 0, 5)
    c := cap(scores)
    fmt.Println(c)
    for i := 0; i < 25; i++ {
        scores = append(scores, i)
        // if our capacity has changed,
        // Go had to grow our array to accommodate the new data
        if cap(scores) != c {
            c = cap(scores)
            fmt.Println(c)
        }
    }
}

func removeAtIndex(source []int, index int) []int {
    lastIndex := len(source) - 1
    //swap the last value and the value we want to remove
    source[index], source[lastIndex] = source[lastIndex], source[index]
    return source[:lastIndex]
}
func f2() {
    scores := []int{1, 2, 3, 4, 5}
    scores = removeAtIndex(scores, 2)
    fmt.Println(scores)
}

func main() {
    scores := []int{1,4,293,4,9}
    scores2 := make([]int, 2, 10)  // 2 is the start-up length, 10 is the capacity (max possible length)

    scores3 := make([]int, 0, 10)
    scores3 = scores3[0:6]
    scores3[5] = 9033
    scores3 = append(scores3, 34)

    fmt.Println(scores, scores2, scores3)
    f1()

    scores4 := make([]int, 5)
    scores4 = append(scores4, 666)
    var scores5 []string
    fmt.Println(scores4, scores5)

    names := []string{"sf"}
    fmt.Println(len(names), cap(names))
    /* Array có số trong ngoặc [num], Slice thì ko []
    Ultimately, there are four common ways to initialize a slice:
    names := []string{"leto", "jessica", "paul"}
    checks := make([]bool, 10)
    var names []string
    scores := make([]int, 0, 20)
    */

    f2()
}
