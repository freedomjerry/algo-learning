package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)
type CodeCate struct {
	 CodeVersion map[string]*[5]int
	 CurrentDir string
}
/*type LeetCodeShow interface {
	GetDirsInfo()

}*/
var TypeMap = map[string]int{"cpp":0, "c": 1, "go": 2, "py": 3, "java":4}

func (codecate *CodeCate) JudgeType(FileName string) int {
	varity := strings.Split(FileName, ".")[1]
	if Type, ok := TypeMap[varity]; ok == true {
		return Type
	}else{
		return -1
	}
}
func (codecate *CodeCate) WriteToMD()  {
	//fmt.Println("here is create func")
	//fmt.Println(codecate.CurrentDir)
	FILE, err := os.Create("LeetCode.md")
	if err != nil {
		fmt.Println("err = ", err)
	}
	FILE.WriteString("## LeetCode Part\n\n")
	FILE.WriteString("### This part is my solution & code for LeetCode\n")
	FILE.WriteString("| Subject | &nbsp;&nbsp;C++&nbsp;&nbsp;&nbsp;   | &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;C&nbsp;&nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;GO&nbsp;&nbsp;&nbsp;&nbsp; | Python3 | &nbsp;&nbsp;Java&nbsp;&nbsp; |\n| ---- | ---- | ---- | ---- | ---- | ---- |\n")
	for key, value := range codecate.CodeVersion{
		FILE.WriteString("| <b>"+key+"</b> |")
		for i := 0; i < 5; i++ {
			if value[i] == 1 {
				FILE.WriteString("√ |")
			}else{
				FILE.WriteString(" |")
			}
		}
		FILE.WriteString("\n")
	}
}
func (codecate *CodeCate)WalkFunc(Path string, info os.FileInfo, err error) error{
	if info.IsDir() {
		if info.Name() != codecate.CurrentDir {
			codecate.CodeVersion[info.Name()] = &[5]int{0, 0, 0, 0, 0}
			codecate.CurrentDir = info.Name()
			//fmt.Println(info.Name())
		}
	} else {
		variety := codecate.JudgeType(info.Name())
		if variety == -1{
			//fmt.Println(info.Name())
		}else {
			cur := codecate.CurrentDir
			codecate.CodeVersion[cur][variety] = 1
			//fmt.Println(variety)
		}
	}
	return nil
}
func main(){
	pwd, _ := os.Getwd()
	root := path.Base(strings.Replace(pwd, "\\", "/", -1))
	codecate := CodeCate{CodeVersion: map[string]*[5]int{},CurrentDir: root}
	filepath.Walk(pwd, codecate.WalkFunc)
	codecate.WriteToMD()
}