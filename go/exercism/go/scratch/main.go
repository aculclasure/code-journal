package main

import (
	"fmt"
	"regexp"
)

func main() {
	firstStr := "The Road _Not_ Taken"
	//secondStr := "metal-oxide"
	re := regexp.MustCompile(`(\w+('\w+)?)*`)

	metalResults := re.FindAllString(firstStr, -1)
	fmt.Printf("%v\n", metalResults)
}