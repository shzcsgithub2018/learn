package Leaf

import (
	"fmt"
)

type Video struct {
	Name string
}

func (h *Video) Search(keyword string) {
	fmt.Printf("Searching for keyworkd %s in video %s\n", keyword, h.Name)
}
