package main

import (
	"fmt"
	"github.com/brettbeloin/E1LinkCP/utils"
)

func main() {
	w := utils.NewSingleLinkList[int]()
	w.Add(1)
	w.Add(2)
	w.Add(21)
	fmt.Println(w.ToString())
}
