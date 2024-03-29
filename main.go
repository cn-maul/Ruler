package main

import (
	"Ruler/internal/regex"
	"fmt"
)

func main() {
	var id = "15736967656"
	output := regex.Match(id)
	fmt.Println(id + " 的类型为 " + output)
}
