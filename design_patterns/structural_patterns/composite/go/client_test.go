package _go

import (
	"github.com/shzgithub2018/learn/design_patterns/structural_patterns/composite/go/Leaf"
	"github.com/shzgithub2018/learn/design_patterns/structural_patterns/composite/go/composite"
	"testing"
)

func TestClient(t *testing.T) {
	img1 := &Leaf.Img{Name: "img1"}
	txt1 := &Leaf.Txt{Name: "txt1"}
	video1 := &Leaf.Video{Name: "video1"}
	img2 := &Leaf.Img{Name: "img2"}

	folder1 := &composite.Folder{
		Name: "Folder1",
	}
	folder1.Add(img1)
	folder1.Add(txt1)

	folder2 := &composite.Folder{
		Name: "Folder2",
	}
	folder2.Add(video1)
	folder2.Add(folder1)

	folder3 := &composite.Folder{
		Name: "Folder3",
	}
	folder3.Add(img2)
	folder3.Add(folder2)

	folder3.Search("txt1")
	/*Output
	Searching for keyworkd txt1 in folder Folder3
	Searching for keyworkd txt1 in img img2
	Searching for keyworkd txt1 in folder Folder2
	Searching for keyworkd txt1 in video video1
	Searching for keyworkd txt1 in folder Folder1
	Searching for keyworkd txt1 in img img1
	Searching for keyworkd txt1 in txt txt1
	*/
}
