package main

import "fmt"
import "./utils"

func main() {
	time := utils.NowWithMillisecond()

	fmt.Println(time)
}
