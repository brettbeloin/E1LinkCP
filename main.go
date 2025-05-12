package main

import (
	"E1Link/utils"
	"fmt"
)

func main() {
	w := singelLink.NewSingleLinkList[int]()
	w.Add(1)
	w.Add(2)
	w.Add(21)
	fmt.Println(w.ToString())
}
