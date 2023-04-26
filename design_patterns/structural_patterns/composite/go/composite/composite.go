package composite

import (
	"fmt"
)

type Folder struct {
	Name     string
	Children []Component
}

func (h *Folder) Search(keyword string) {
	fmt.Printf("Searching for keyworkd %s in folder %s\n", keyword, h.Name)
	for _, composite := range h.Children {
		composite.Search(keyword)
	}
}

//Add 该操作也可以考虑放到 component interface，透明性
func (h *Folder) Add(c Component) {
	h.Children = append(h.Children, c)
}

func (h *Folder) Remove(c Component) {

}
