package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]bool

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[2] = "Columbia"
	Countries[3] = "USA"

	qty := len(Countries)
	fmt.Println("Countries Saved", qty)
}
