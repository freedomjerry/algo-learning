package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type CodeCate struct {
	CodeVersion map[string]*map[string]string
	CurrentDir  string
}

/*type LeetCodeShow interface {
	GetDirsInfo()

}*/
var suffixMap = []string{"txt","cpp", "c", "go", "py", "java"}

//var suffixMap = map[int]string{0: "cpp", 1: "c", 2: "go", 3: "py", 4: "java"}
/*func (codecate *CodeCate) JudgeType(FileName string) int {
	varity := strings.Split(FileName, ".")[1]
	if Type, ok := suffixMap[varity]; ok == true {
		return Type
	}else{
		return -1
	}
}*/
func (codecate *CodeCate) WriteToMD() {
	FILE, err := os.Create("LeetCode.md")
	defer FILE.Close()
	if err != nil {
		fmt.Println("err = ", err)
	}
	FILE.WriteString("## LeetCode Part\n\n")
	FILE.WriteString("### This part is my solution & code for LeetCode\n")
	FILE.WriteString("|<div style='width:100px'> Subject </div>|")
	for _, v := range suffixMap{
		if v == "py" {v = "python"}
		if v == "txt" {v = "tag"}
		FILE.WriteString("<div style='width:50px'>"+ v +"</div>|")
	}
	FILE.WriteString("\n|")
	for i:=0; i <= len(suffixMap); i++ {
		FILE.WriteString(" :----: |")
	}
	for key, value := range codecate.CodeVersion {
		FILE.WriteString("\n| <b>" + key + "</b> |")
		for _,variety := range suffixMap{
			if done, ok := (*value)[variety]; ok == true{
				if done == "√" {
					FILE.WriteString("[√](./" + key + "/" + key + "." + variety + ") |")
				} else {
					FILE.WriteString(" |")
				}
			}else{
				FILE.WriteString(" |")
			}
		}
	}
}
func (codecate *CodeCate) WalkFunc(Path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		if info.Name() != codecate.CurrentDir {
			codecate.CodeVersion[info.Name()] = &map[string]string{"cpp": " ", "c": " ",
				"go": " ", "py": " ", "java": " "}
			codecate.CurrentDir = info.Name()
		}
	} else {
		if info.Name() == "catagenerator.go" {
			return nil
		}
		variety := strings.Split(path.Ext(info.Name()), ".")[1]
		if variety == ".txt" {

		}else {
			cur := codecate.CurrentDir
			if _, ok := (*codecate.CodeVersion[cur])[variety]; ok == true {
				(*codecate.CodeVersion[cur])[variety] = "√"
			} else {
			}
		}
	}
	return nil
}
func main() {
	pwd, _ := os.Getwd()
	root := path.Base(strings.Replace(pwd, "\\", "/", -1))
	codecate := CodeCate{CodeVersion: map[string]*map[string]string{}, CurrentDir: root}
	filepath.Walk(pwd, codecate.WalkFunc)
	codecate.WriteToMD()
}
