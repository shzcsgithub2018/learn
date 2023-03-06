package Leaf

import (
	"fmt"
)

type Img struct {
	Name string
}

func (h *Img) Search(keyword string) {
	fmt.Printf("Searching for keyworkd %s in img %s\n", keyword, h.Name)
}
