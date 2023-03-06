package composite

import (
	"fmt"
	_go "github.com/shzgithub2018/learn/design_patterns/structural_patterns/composite/go/interface"
)

type Folder struct {
	Name     string
	Children []_go.Component
}

func (h *Folder) Search(keyword string) {
	fmt.Printf("Searching for keyworkd %s in folder %s\n", keyword, h.Name)
	for _, composite := range h.Children {
		composite.Search(keyword)
	}
}

//Add 该操作也可以考虑放到 component interface，透明性
func (h *Folder) Add(c _go.Component) {
	h.Children = append(h.Children, c)
}

func (h *Folder) Remove(c _go.Component) {

}
