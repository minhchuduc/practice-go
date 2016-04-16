package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 8)
	fmt.Println(worst, scores[4:20])
	copy(worst[2:4], scores[4:])
	fmt.Println(worst)
	//0 1 2 3 4 5 6

	// Slice of slices of slices
	board := [][][]string{
		[][]string{[]string{"_", "_"}, []string{"_", "_"}},
		[][]string{[]string{"_", "_"}, []string{"_", "_"}},
		[][]string{[]string{"_", "_"}, []string{"_", "_"}},
	}
	fmt.Println(board)

	cfgStr := `; Comment line
    [section1]
    var-name=value # comment`
	fmt.Println("cfgStr = ", cfgStr)

	cfg := struct {
		Section1 struct {
			FieldName string `gcfg:"var-name"`
		}
	}{}
	fmt.Printf("Type of cfg is %T \n", cfg)
	fmt.Println("cfg Struct = ", cfg)

	err := gcfg.ReadStringInto(&cfg, cfgStr)
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	fmt.Println(cfg.Section1.FieldName)

	fmt.Println("------------------------------------------")
	cfgStr2 := `; Comment line
    [section]
    multi=value1
    multi=value2`
	cfg2 := struct {
		Section struct {
			Multi []string
		}
	}{}
	err2 := gcfg.ReadStringInto(&cfg2, cfgStr2)
	if err2 != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err2)
	}
	fmt.Println(cfg2.Section.Multi)
	fmt.Printf("Type of cfg2.Section.Multi is %T\n", cfg2.Section.Multi)

	a := []string{"abasd", "jhfh", "ccc"}
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println(len(a), cap(a))
	// initialize a 3-dimension slices with random numbers
	// How?
}
