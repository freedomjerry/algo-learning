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
var TypeMap = []string{"txt","cpp", "c", "go", "py", "java"}

//var suffixMap = map[int]string{0: "cpp", 1: "c", 2: "go", 3: "py", 4: "java"}
/*func (codecate *CodeCate) JudgeType(FileName string) int {
	varity := strings.Split(FileName, ".")[1]
	if Type, ok := TypeMap[varity]; ok == true {
		return Type
	}else{
		return -1
	}
}*/
func (codecate *CodeCate) WriteToMD() {
	//fmt.Println("here is create func")
	//fmt.Println(codecate.CurrentDir)
	FILE, err := os.Create("LeetCode.md")
	if err != nil {
		fmt.Println("err = ", err)
	}
	FILE.WriteString("## LeetCode Part\n\n")
	FILE.WriteString("### This part is my solution & code for LeetCode\n")
	FILE.WriteString("|<div style='width:100px'> Subject </div>|")
	for _, v := range TypeMap{
		if v == "py" {v = "python"}
		if v == "txt" {v = "tag"}
		FILE.WriteString("<div style='width:50px'>"+ v +"</div>|")
	}
	FILE.WriteString("\n|")
	for i:=0; i <= len(TypeMap); i++ {
		FILE.WriteString(" :----: |")
	}
	for key, value := range codecate.CodeVersion {
		FILE.WriteString("\n| <b>" + key + "</b> |")
		/*for subtype, done := range *value {
			if done == "√" {
				FILE.WriteString("[√](./" + key + "/" + key + "." + subtype + ") |")
			} else {
				FILE.WriteString(" |")
			}
		}*/
		for _,variety := range TypeMap{
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
			//fmt.Println(info.Name())
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
				//fmt.Println(variety)
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
	for key , v := range codecate.CodeVersion{
		fmt.Println(key, *v)
	}
	codecate.WriteToMD()
}
