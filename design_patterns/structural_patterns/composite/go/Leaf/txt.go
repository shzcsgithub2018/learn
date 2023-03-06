package Leaf

import (
	"fmt"
)

type Txt struct {
	Name string
}

func (h *Txt) Search(keyword string) {
	fmt.Printf("Searching for keyworkd %s in txt %s\n", keyword, h.Name)
}
