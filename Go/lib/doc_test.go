package main

//
//import (
//	"baliance.com/gooxml/document"
//	"fmt"
//	"github.com/xuri/excelize/v2"
//	"os"
//	"testing"
//)
//
//func TestWrapDoc(t *testing.T) {
//	doc, err := document.Open("test.docx")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	t.Log(doc)
//
//	file := excelize.NewFile()
//	sheetName := "Sheet1"
//	sheet, err := file.NewSheet(sheetName)
//	if err != nil {
//		t.Error(err)
//	}
//
//	var (
//		colIndex       = 1
//		columnNameList = []string{"question", "answer"}
//	)
//
//	for _, columnName := range columnNameList {
//		cellName, _ := excelize.CoordinatesToCellName(colIndex, 1)
//		file.SetCellValue(sheetName, cellName, columnName)
//		colIndex++
//	}
//
//	doc
//	for _, para := range doc.Paragraphs() {
//		text := ""
//		//run为每个段落相同格式的文字组成的片段
//		for _, run := range para.Runs(). {
//			text += run.Text()
//			//t.Log("text:", run.Text())
//			//res := strings.Split(run.Text(), "?")
//			//t.Log(res[0])
//			//cellName, _ := excelize.CoordinatesToCellName(colIndex, i+2)
//			//file.SetCellValue(sheetName, cellName, res[0])
//			//
//			//cellName2, _ := excelize.CoordinatesToCellName(colIndex+1, i+2)
//			//file.SetCellValue(sheetName, cellName2, res[1])
//		}
//		t.Log(text)
//	}
//	file.SetActiveSheet(sheet)
//	buffer, err := file.WriteToBuffer()
//
//	// 创建文件
//	file1, err := os.Create("output.txt")
//	if err != nil {
//		fmt.Println("创建文件失败：", err)
//		return
//	}
//	defer file1.Close()
//
//	// 将 buffer 中的内容写入文件
//	_, err = buffer.WriteTo(file1)
//	if err != nil {
//		fmt.Println("写入文件失败：", err)
//		return
//	}
//
//	fmt.Println("文件保存成功。")
//
//	return
//}
